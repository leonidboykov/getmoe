package getmoe

import (
	"net/http"

	"github.com/leonidboykov/getmoe/conf"
)

// Provider describes Board provider
type Provider interface {
	Auth(conf.AuthConfiguration)
	BuildRequest(conf.RequestConfiguration)
	NextPage()
	PageRequest() (*http.Request, error)
	Parse([]byte) ([]Post, error)
}
