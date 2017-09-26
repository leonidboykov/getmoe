package getmoe

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Config holds data for API access
type Config struct {
	URL          url.URL
	BaseURL      string
	PasswordSalt string
	Limit        int
	UserAgent    string
	AppkeySalt   string
	Query        struct {
		Tags string
		Page int
	}
}

// BuildAuth creates query for auth
func (c *Config) BuildAuth(login, password string) {
	u := url.Values{}
	u.Add("login", login)
	u.Add("password_hash", Sha1(password, c.PasswordSalt))

	// if AppkeySalt is not empty (for Sankaku Channel)
	if c.AppkeySalt != "" {
		u.Add("appkey", Sha1(login, c.AppkeySalt))
	}

	c.URL.RawQuery = u.Encode()
}

// BuildRequest ...
func (c *Config) BuildRequest(tags []string) url.URL {
	tempURL := c.URL

	u := tempURL.Query()

	t := strings.Join(tags, " ")
	u.Add("tags", t)
	tempURL.RawQuery = u.Encode()
	return tempURL
}

// Request gets images by tags
func (c *Config) Request(tags []string, form url.Values) ([]Post, error) {
	// Remove Config reference from BuildTags
	url := c.BuildRequest(tags)

	resp, err := http.Get(url.String())
	if err != nil {
		return []Post{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Post{}, err
	}

	// TODO: Provide API related reader
	var page []Post
	if err = json.Unmarshal(body, &page); err != nil {
		return []Post{}, err
	}

	return page, nil
}

// RequestAll checks all pages
func (c *Config) RequestAll(tags []string) ([]Post, error) {
	var pages []Post
	localQuery := query

	for {
		page, err := Request(tags)
		if err != nil {
			return pages, err
		}

		localQuery.Page++
		if len(page) == 0 {
			break
		}

		pages = append(pages, page...)
		println(localQuery.Page)
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
