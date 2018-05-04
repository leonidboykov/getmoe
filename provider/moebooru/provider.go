/*
Package moebooru implements a simple library for accessing Moebooru-based image
boards.

Source code of Moebooru is available at https://github.com/moebooru/moebooru

Default configurations are available for the following websites

  * yande.re
  * konachan.com
*/
package moebooru

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/leonidboykov/getmoe"

	"github.com/imdario/mergo"

	"github.com/leonidboykov/getmoe/conf"
	"github.com/leonidboykov/getmoe/internal/hash"
	"github.com/leonidboykov/getmoe/internal/query"
)

const (
	defaultPasswordSalt = "choujin-steiner--%s--"
	defaultPostsLimit   = 1000
)

const (
	loginKey        = "login"
	passwordHashKey = "password_hash"
	pageKey         = "page"
	tagsKey         = "tags"
)

var defaultProvider = &Provider{
	URL: &url.URL{
		Scheme: "https",
		Path:   "post.json",
	},
	PasswordSalt: "choujin-steiner--%s--",
	PostsLimit:   1000,
}

// Provider implements moebooru provider
type Provider struct {
	URL          *url.URL
	Headers      map[string]string
	PasswordSalt string
	PostsLimit   int
}

// New creates a new moebooru provider with configuration
func New(config conf.ProviderConfiguration) *Provider {
	provider := &Provider{
		URL:          &config.URL.URL,
		Headers:      config.Headers,
		PasswordSalt: config.PasswordSalt,
		PostsLimit:   config.PostsLimit,
	}
	// Apply defaults
	mergo.Merge(provider, defaultProvider)
	// Authenticate if login/password have provided
	provider.Auth(config.Auth)
	return provider
}

// Auth builds query based on AuthConfiguration
func (p *Provider) Auth(config conf.AuthConfiguration) {
	var login, password, hashedPassword = config.Login, config.Password, config.HashedPassword
	q := p.URL.Query()
	if login != "" {
		q.Set(loginKey, login)
	}
	if hashedPassword == "" && password != "" {
		hashedPassword = hash.Sha1(password, p.PasswordSalt)
	}
	if hashedPassword != "" {
		q.Set(passwordHashKey, hashedPassword)
	}
	p.URL.RawQuery = q.Encode()
}

// BuildRequest builds query based on RequestConfiguration
func (p *Provider) BuildRequest(config conf.RequestConfiguration) {
	q := p.URL.Query()
	query.Array(&q, "tags", config.Tags)
	query.Int(&q, "limit", p.PostsLimit)
	p.URL.RawQuery = q.Encode()
}

// NextPage increments page by 1
func (p *Provider) NextPage() {
	q := p.URL.Query()
	query.Increment(&q, "page")
	p.URL.RawQuery = q.Encode()
}

// PageRequest builds page request from URL
func (p *Provider) PageRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", p.URL.String(), nil)
	if err != nil {
		return req, err
	}
	// Set headers if provided
	for k, v := range p.Headers {
		req.Header.Set(k, v)
	}
	return req, err
}

// Parse parses result from query
func (p *Provider) Parse(data []byte) ([]getmoe.Post, error) {
	var page []Post
	if err := json.Unmarshal(data, &page); err != nil {
		return nil, err
	}
	result := make([]getmoe.Post, len(page))
	for i := range page {
		result[i] = getmoe.Post{
			ID:        page[i].ID,
			FileURL:   page[i].FileURL,
			FileSize:  page[i].FileSize,
			Width:     page[i].Width,
			Height:    page[i].Height,
			CreatedAt: page[i].parseTime(),
			Author:    page[i].Author,
			Source:    page[i].Source,
			Rating:    page[i].Rating,
			Hash:      page[i].Md5,
			Tags:      page[i].parseTags(),
			Score:     page[i].Score,
		}
	}
	return result, nil
}
