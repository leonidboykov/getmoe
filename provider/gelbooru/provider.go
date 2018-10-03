/*
Package gelbooru implements a simple library for accessing Gelbooru-based image
boards.
*/
package gelbooru

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/imdario/mergo"

	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/internal/query"
)

const (
	defaultPasswordSalt = "choujin-steiner--%s--"
	defaultPostsLimit   = 1000
)

const (
	loginKey        = "login"
	passwordHashKey = "password_hash"
	pageKey         = "pid"
	tagsKey         = "tags"
)

var defaultProvider = &Provider{
	URL: &url.URL{
		Scheme:   "https",
		Path:     "index.php",
		RawQuery: "page=dapi&s=post&q=index&json=1",
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
	// No auth required
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
			Width:     page[i].Width,
			Height:    page[i].Height,
			CreatedAt: page[i].parseTime(),
			Rating:    page[i].Rating,
			Hash:      page[i].Hash,
			Tags:      page[i].parseTags(),
			Score:     page[i].Score,
		}
	}
	return result, nil
}
