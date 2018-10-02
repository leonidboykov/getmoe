package getmoe

import (
	"net/http"
)

// Provider describes Board provider
type Provider interface {
	Auth(AuthConfiguration)
	BuildRequest(RequestConfiguration)
	NextPage()
	PageRequest() (*http.Request, error)
	Parse([]byte) ([]Post, error)
}
