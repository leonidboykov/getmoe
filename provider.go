package getmoe

import (
	"net/url"

	"github.com/leonidboykov/getmoe/conf"
)

// Provider describes Board provider
type Provider interface {
	Auth(conf.AuthConfiguration, *url.URL)
	// Parse()
}
