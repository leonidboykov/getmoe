package board

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/conf"
	"github.com/leonidboykov/getmoe/provider"
)

// Board related errors
var (
	ErrURLNotSpecified      = errors.New("board: URL not specified")
	ErrProviderNotSpecified = errors.New("board: provider not specified")
	ErrProviderNotFound     = "board: provider %s not found"
)

// Board holds data for API access
type Board struct {
	// URL *url.URL
	Provider   getmoe.Provider
	httpClient *http.Client
}

// New creates a new Board with provided configuration
func New(config conf.BoardConfiguration) (*Board, error) {
	// if config.Provider.URL == "" {
	// 	return nil, ErrURLNotSpecified
	// }
	// u, err := url.Parse(config.Provider.URL)
	// if err != nil {
	// 	return nil, err
	// }

	if config.Provider.Name == "" {
		return nil, ErrProviderNotSpecified
	}

	p, ok := provider.Providers[config.Provider.Name]
	if !ok {
		return nil, fmt.Errorf(ErrProviderNotFound, config.Provider)
	}

	board := &Board{
		// URL:      u,
		Provider: p(config.Provider),
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	return board, nil
}
