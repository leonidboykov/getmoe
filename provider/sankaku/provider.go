/*
Package sankaku implements a simple library for accessing Sankakucomplex-based
image boards.
*/
package sankaku

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/imdario/mergo"

	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/internal/hash"
	"github.com/leonidboykov/getmoe/internal/query"
)

const (
	defaultPasswordSalt = "choujin-steiner--%s--"
	defaultPostsLimit   = 100
)

const (
	loginKey        = "login"
	passwordHashKey = "password_hash"
	appkeyKey       = "appkey"
	pageKey         = "page"
	tagsKey         = "tags"
)

var defaultProvider = &Provider{
	URL: &url.URL{
		Scheme: "https",
		Path:   "post/index.json",
	},
	Headers: map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36",
		"Referer":    "https://beta.sankakucomplex.com",
		"Origin":     "https://beta.sankakucomplex.com",
		"Accept":     "application/json",
	},
	PasswordSalt: "choujin-steiner--%s--",
	AppkeySalt:   "sankakuapp_%s_Z5NE9YASej",
	PostsLimit:   100,
}

// Provider implements moebooru provider
type Provider struct {
	URL          *url.URL
	Headers      map[string]string
	PasswordSalt string
	AppkeySalt   string
	PostsLimit   int
}

// New creates a new moebooru provider with configuration
func New(config getmoe.ProviderConfiguration) *Provider {
	provider := &Provider{
		URL:          &config.URL.URL,
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
func (p *Provider) Auth(config getmoe.AuthConfiguration) {
	var login, password, hashedPassword = config.Login, config.Password, config.HashedPassword
	q := p.URL.Query()
	if login != "" {
		appKey := hash.Sha1(strings.ToLower(login), p.AppkeySalt)
		q.Set(loginKey, login)
		q.Set(appkeyKey, appKey)
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
	for k, v := range p.Headers {
		req.Header.Set(k, v)
	}
	if err != nil {
		return req, err
	}
	return req, err
}

// Parse parses result from query
func (p *Provider) Parse(data []byte) ([]getmoe.Post, error) {
	var page []Post
	if err := json.Unmarshal(data, &page); err != nil {
		fmt.Println(p.URL.String())
		fmt.Println(string(data))
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
			Author:    page[i].findArtist(),
			Source:    page[i].Source,
			Rating:    page[i].Rating,
			Hash:      page[i].Md5,
			Tags:      page[i].parseTags(),
			Score:     page[i].TotalScore,
		}
	}
	return result, nil
}
