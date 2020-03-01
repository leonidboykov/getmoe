/*
Package sankaku implements a simple library for accessing Sankakucomplex-based
image boards.
*/
package sankaku

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/schema"
	"github.com/imdario/mergo"

	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/internal/hash"
	"github.com/leonidboykov/getmoe/internal/query"
)

const (
	loginKey        = "login"
	passwordHashKey = "password_hash"
	appkeyKey       = "appkey"
	pageKey         = "page"
	tagsKey         = "tags"
)

type queryValues struct {
	Login    string `schema:"login,omitempty"`
	Password string `schema:"password_hash,omitempty"`
	Appkey   string `schema:"appkey,omitempty"`
	Tags     string `schema:"tags,omitempty"`
	Page     int    `schema:"page,omitempty"`
}

var defaultProvider = &Provider{
	URL: &url.URL{
		Scheme: "https",
		// Path:   "post/index.json",
		Path: "posts",
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

// Provider implements sankaku provider.
type Provider struct {
	URL           *url.URL
	Headers       map[string]string
	PasswordSalt  string
	AppkeySalt    string
	PostsLimit    int
	schemaEncoder *schema.Encoder
}

// New creates a new sankaku provider with configuration.
func New(config getmoe.ProviderConfiguration) *Provider {
	var provider *Provider
	provider.New(config)
	return provider
}

// New creates a new sankaku provider with configuration.
func (p *Provider) New(config getmoe.ProviderConfiguration) {
	p.URL = &config.URL.URL
	p.PasswordSalt = config.PasswordSalt
	p.PostsLimit = config.PostsLimit
	p.schemaEncoder = schema.NewEncoder()

	// Apply defaults
	mergo.Merge(p, defaultProvider)
	// Authenticate if login/password have provided
	p.Auth(config.Auth)
}

// Auth builds query based on AuthConfiguration.
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

// BuildRequest builds query based on RequestConfiguration.
func (p *Provider) BuildRequest(config getmoe.RequestConfiguration) {
	q := p.URL.Query()
	q.Set("tags", config.Tags.String())
	query.Int(&q, "limit", p.PostsLimit)
	p.URL.RawQuery = q.Encode()
}

// NextPage increments page by 1.
func (p *Provider) NextPage() {
	q := p.URL.Query()
	query.Increment(&q, "page")
	p.URL.RawQuery = q.Encode()
}

// PageRequest builds page request from URL.
func (p *Provider) PageRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", p.URL.String(), nil)
	for k, v := range p.Headers {
		req.Header.Set(k, v)
	}
	if err != nil {
		return nil, err
	}
	return req, err
}

// Parse parses result from query.
func (p *Provider) Parse(data []byte) ([]getmoe.Post, error) {
	var page []post
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
			CreatedAt: page[i].CreatedAt.Time,
			Author:    page[i].findArtist(),
			Source:    page[i].Source,
			Rating:    page[i].Rating,
			Hash:      page[i].Hash,
			Tags:      page[i].parseTags(),
			Score:     page[i].TotalScore,
			VoteCount: page[i].VoteCount,
			FavCount:  page[i].FavCount,
		}
	}
	return result, nil
}

func init() {
	getmoe.RegisterProvider("sankaku", &Provider{})
}
