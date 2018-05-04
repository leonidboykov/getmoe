package board

import (
	"errors"
	"fmt"
	"io/ioutil"
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
	Provider   getmoe.Provider
	httpClient *http.Client
}

// New creates a new Board with provided configuration
func New(config conf.BoardConfiguration) (*Board, error) {
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
			Timeout: 30 * time.Second,
		},
	}

	return board, nil
}

// Request gets images by tags
func (b *Board) Request() ([]getmoe.Post, error) {
	req, err := b.Provider.PageRequest()
	if err != nil {
		return nil, err
	}

	resp, err := b.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	page, err := b.Provider.Parse(body)
	if err != nil {
		return nil, err
	}

	return page, nil
}

// RequestAll checks all pages
func (b *Board) RequestAll() ([]getmoe.Post, error) {
	var pages []getmoe.Post
	for {
		page, err := b.Request()
		if err != nil {
			return pages, err
		}
		if len(page) == 0 {
			break
		}
		pages = append(pages, page...)
		b.Provider.NextPage()
	}
	return pages, nil
}
