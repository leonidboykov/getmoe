package board

import (
	"errors"
	"fmt"
	"net/url"

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
	URL      *url.URL
	Provider getmoe.Provider
}

// New creates a new Board with provided configuration
func New(config conf.BoardConfiguration) (*Board, error) {
	if config.URL == "" {
		return nil, ErrURLNotSpecified
	}
	u, err := url.Parse(config.URL)
	if err != nil {
		return nil, err
	}

	if config.Provider == "" {
		return nil, ErrProviderNotSpecified
	}

	p, ok := provider.Providers[config.Provider]
	if !ok {
		return nil, fmt.Errorf(ErrProviderNotFound, config.Provider)
	}

	board := &Board{
		URL:      u,
		Provider: p,
	}

	return board, nil
}
