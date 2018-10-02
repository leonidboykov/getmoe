package getmoe

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// Board related errors
var (
	ErrURLNotSpecified      = errors.New("board: URL not specified")
	ErrProviderNotSpecified = errors.New("board: provider not specified")
	ErrProviderNotFound     = "board: provider %s not found"
)

// Board holds data for API access
type Board struct {
	Provider   Provider
	httpClient *http.Client
}

// NewBoard creates a new board with provided configuration
func NewBoard(provider Provider) *Board {
	// func New(config getmoe.BoardConfiguration) (*Board, error) {
	// if config.Provider.Name == "" {
	// 	return nil, ErrProviderNotSpecified
	// }
	// p, ok := provider.Providers[config.Provider.Name]
	// if !ok {
	// 	return nil, fmt.Errorf(ErrProviderNotFound, config.Provider.Name)
	// }

	return &Board{
		Provider: provider,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Request gets images by tags
func (b *Board) Request() ([]Post, error) {
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
func (b *Board) RequestAll() ([]Post, error) {
	var pages []Post
	for {
		b.Provider.NextPage()
		page, err := b.Request()
		if err != nil {
			return pages, err
		}
		if len(page) == 0 {
			break
		}
		pages = append(pages, page...)
	}
	return pages, nil
}
