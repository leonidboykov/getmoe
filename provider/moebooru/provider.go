/*
Package moebooru implements a simple library for accessing Moebooru-based image
boards.

Source code of Moebooru is available at https://github.com/moebooru/moebooru
*/
package moebooru

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/imdario/mergo"

	"github.com/leonidboykov/getmoe"
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
	PasswordSalt string
	PostsLimit   int
}

// New creates a new moebooru provider with configuration
func New(config getmoe.ProviderConfiguration) *Provider {
	var provider *Provider
	provider.New(config)
	return provider
}

// New creates a new moebooru provider with configuration
func (p *Provider) New(config getmoe.ProviderConfiguration) {
	p.URL = &config.URL.URL
	p.PasswordSalt = config.PasswordSalt
	p.PostsLimit = config.PostsLimit
	// Apply defaults
	mergo.Merge(p, defaultProvider)
	// Authenticate if login/password have provided
	p.Auth(config.Auth)
}

// Auth builds query based on AuthConfiguration
func (p *Provider) Auth(config getmoe.AuthConfiguration) {
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
func (p *Provider) BuildRequest(config getmoe.RequestConfiguration) {
	q := p.URL.Query()
	q.Set("tags", config.Tags.String())
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

func init() {
	getmoe.RegisterProvider("moebooru", &Provider{})
}
