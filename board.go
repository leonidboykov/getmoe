package getmoe

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Board holds data for API access
type Board struct {
	URL          url.URL
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

// BuildAuth creates query for auth
func (c *Board) BuildAuth(login, password string) {
	q := c.URL.Query()
	q.Set("login", login)
	q.Set("password_hash", Sha1(password, c.PasswordSalt))

	// if AppkeySalt is not empty (for Sankaku Channel)
	if c.AppkeySalt != "" {
		q.Set("appkey", Sha1(strings.ToLower(login), c.AppkeySalt))
	}

	c.URL.RawQuery = q.Encode()
}

// BuildRequest ...Set
func (c *Board) BuildRequest() url.URL {
	u := c.URL

	q := u.Query()

	t := strings.Join(c.Query.Tags, " ")
	q.Set("tags", t)
	q.Set("limit", strconv.Itoa(c.Limit))

	q.Set(c.PageTag, strconv.Itoa(c.Query.Page))

	u.RawQuery = q.Encode()
	return u
}

// Request gets images by tags
func (c *Board) Request() ([]Post, error) {
	// Remove Board reference from BuildTags
	url := c.BuildRequest()

	// There is no point to create new http client every request
	client := &http.Client{}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return []Post{}, nil
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
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

	page, err := c.Parse(body)
	if err != nil {
		return []Post{}, err
	}

	return page, nil
}

// RequestAll checks all pages
func (c *Board) RequestAll() ([]Post, error) {
	var pages []Post

	for {
		page, err := c.Request()
		if err != nil {
			return pages, err
		}

		if len(page) == 0 {
			break
		}

		pages = append(pages, page...)

		c.Query.Page++
	}
	return pages, nil
}

// Sha1 builds Sha1 hash with proper salt
func Sha1(value, salt string) string {
	value = fmt.Sprintf(salt, value)
	hash := sha1.New()
	hash.Write([]byte(value))
	sha1Hash := hex.EncodeToString(hash.Sum(nil))
	return sha1Hash
}
