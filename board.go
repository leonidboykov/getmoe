package getmoe

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Board holds data for API access
type Board struct {
	Provider   Provider
	httpClient *http.Client
}

// NewBoard creates a new board with provided configuration
// func NewBoard(provider Provider) *Board {
// 	return &Board{
// 		Provider: provider,
// 		httpClient: &http.Client{
// 			Timeout: 10 * time.Second,
// 		},
// 	}
// }

// NewBoard creates a new board with provided configuration
func NewBoard(providerName string, config BoardConfiguration) (*Board, error) {
	providersMu.RLock()
	provider, ok := providers[providerName]
	providersMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("getmoe: unknown provider %s", providerName)
	}

	provider.New(config.Provider)
	return &Board{
		Provider: provider,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil
}

// NewBoardWithProvider creates a new board with provided configuration
func NewBoardWithProvider(provider Provider) *Board {
	return &Board{
		Provider: provider,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
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
