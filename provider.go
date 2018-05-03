package getmoe

import (
	"github.com/leonidboykov/getmoe/conf"
)

// Provider describes Board provider
type Provider interface {
	Auth(conf.AuthConfiguration)
	// Parse()
}
