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
	"net/url"

	"github.com/leonidboykov/getmoe/conf"
	"github.com/leonidboykov/getmoe/internal/hash"
)

const (
	defaultPasswordSalt = "choujin-steiner--%s--"
)

// Provider implements moebooru provider
type Provider struct {
	// Login          string
	// Password       string
	// HashedPassword string
	// PasswordSalt   string
}

// func New(config conf.BoardConfiguration) (*Provider, error) {

// }

// Auth builds query based on AuthConfiguration
func (p *Provider) Auth(config conf.AuthConfiguration, u *url.URL) {
	var login, password, hashedPassword = config.Login, config.Password, config.HashedPassword
	var passwordSalt = config.PasswordSalt
	if passwordSalt == "" {
		passwordSalt = defaultPasswordSalt
	}

	q := u.Query()
	if login != "" {
		q.Set("login", login)
	}
	if hashedPassword == "" && password != "" {
		hashedPassword = hash.Sha1(password, passwordSalt)
	}
	if hashedPassword != "" {
		q.Set("password_hash", hashedPassword)
	}

	u.RawQuery = q.Encode()
}
