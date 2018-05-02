package getmoe

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/leonidboykov/getmoe/conf"
	"github.com/leonidboykov/getmoe/internal/hash"
)

// Board holds data for API access
type Board struct {
	URL          url.URL
	Provider     *Provider
	PasswordSalt string
	Limit        int
	UserAgent    string
	AppkeySalt   string
	PageTag      string
	Parse        func(data []byte) ([]Post, error)
	Query
}

// Query ...
type Query struct {
	Tags []string
	Page int
}

// NewBoard creates a new board
func NewBoard(config conf.BoardConfiguration) (*Board, error) {
	u, err := url.Parse(config.URL)
	if err != nil {
		return nil, err
	}

	// p, ok := provider.Providers[config.Provider]
	// if !ok {
	// 	return nil, fmt.Errorf("board: unable to use provider %s", config.Provider)
	// }

	board := &Board{
		URL: *u,
		// Provider: *p,
	}

	return board, nil
}

// BuildAuth creates query for auth
func (b *Board) BuildAuth(login, password string) {
	q := b.URL.Query()
	q.Set("login", login)
	q.Set("password_hash", hash.Sha1(password, b.PasswordSalt))

	// if AppkeySalt is not empty (for Sankaku Channel)
	if b.AppkeySalt != "" {
		q.Set("appkey", hash.Sha1(strings.ToLower(login), b.AppkeySalt))
	}

	b.URL.RawQuery = q.Encode()
}

// BuildRequest ...Set
func (b *Board) BuildRequest() url.URL {
	u := b.URL

	q := u.Query()

	t := strings.Join(b.Query.Tags, " ")
	q.Set("tags", t)
	q.Set("limit", strconv.Itoa(b.Limit))

	q.Set(b.PageTag, strconv.Itoa(b.Query.Page))

	u.RawQuery = q.Encode()
	return u
}

// Request gets images by tags
func (b *Board) Request() ([]Post, error) {
	// Remove Board reference from BuildTags
	url := b.BuildRequest()

	// There is no point to create new http client every request
	client := &http.Client{}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return []Post{}, nil
	}

	if b.UserAgent != "" {
		req.Header.Set("User-Agent", b.UserAgent)
	}

	resp, err := client.Do(req)
	if err != nil {
		return []Post{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Post{}, err
	}

	page, err := b.Parse(body)
	if err != nil {
		return []Post{}, err
	}

	return page, nil
}

// RequestAll checks all pages
func (b *Board) RequestAll() ([]Post, error) {
	var pages []Post

	for {
		page, err := b.Request()
		if err != nil {
			return pages, err
		}

		if len(page) == 0 {
			break
		}

		pages = append(pages, page...)

		b.Query.Page++
	}
	return pages, nil
}
