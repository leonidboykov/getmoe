/*
Package gelbooru implements a simple library for accessing Gelbooru-based image
boards.

Source code of Moebooru is available at https://github.com/moebooru/moebooru

Default configurations are available for the following websites

	gelbooru

*/
package gelbooru

import (
	"encoding/json"
	"net/url"

	"github.com/leonidboykov/getmoe"
)

var (
	// GelbooruConfig preconfigured config for Gelbooru site
	GelbooruConfig = getmoe.Board{
		URL: url.URL{
			Scheme:   "https",
			Host:     "gelbooru.com",
			Path:     "index.php",
			RawQuery: "page=dapi&s=post&q=index&json=1",
		},
		Limit:   1000,
		PageTag: "pid",
		Parse:   parse,
	}
)

func parse(data []byte) ([]getmoe.Post, error) {
	var page []Post
	if err := json.Unmarshal(data, &page); err != nil {
		return nil, err
	}

	result := make([]getmoe.Post, len(page))

	for i, p := range page {
		result[i] = getmoe.Post{
			ID:        p.ID,
			FileURL:   p.FileURL,
			Width:     p.Width,
			Height:    p.Height,
			CreatedAt: p.parseTime(),
			Rating:    p.Rating,
			Hash:      p.Hash,
			Tags:      p.parseTags(),
			Score:     p.Score,
		}
	}

	return result, nil
}
