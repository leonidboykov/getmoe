package getmoe

// Board holds data for API access.
type Board struct {
	name     string
	provider Provider
}

// NewBoard creates a new Board.
func NewBoard(name string, config BoardConfiguration) (*Board, error) {
	if err := applyPresets(config.Settings, &config.Provider); err != nil {
		return nil, err
	}

	provider, err := NewProvider(config.Provider.Name, config.Provider)
	if err != nil {
		return nil, err
	}

	return &Board{
		name:     name,
		provider: provider,
	}, nil
}

func (b *Board) RequestAll(tags Tags) ([]Post, error) {
	var pages []Post
	currentPage := 0
	for {
		page, err := b.provider.RequestPage(tags, currentPage)
		if err != nil {
			return pages, err
		}
		if len(page) == 0 {
			break
		}
		pages = append(pages, page...)
		currentPage++
	}
	return pages, nil
}
