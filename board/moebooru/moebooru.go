/*
Package moebooru implements a simple library for accessing Moebooru-based image
boards.

Source code of Moebooru is available at https://github.com/moebooru/moebooru

Default configurations are available for the following websites

	yande.re
	konachan.com
*/
package moebooru

import (
	"encoding/json"
	"net/url"

	"github.com/leonidboykov/getmoe"
)

var (
	// YandeReConfig preconfigured config for yande.re site
	YandeReConfig = getmoe.Board{
		URL: url.URL{
			Scheme: "https",
			Host:   "yande.re",
			Path:   "post.json",
		},
		PasswordSalt: "choujin-steiner--%s--",
		Limit:        1000,
		PageTag:      "page",
		Parse:        parse,
	}
	// KonachanConfig preconfigured config for konachan.com site
	KonachanConfig = getmoe.Board{
		URL: url.URL{
			Scheme: "https",
			Host:   "konachan.com",
			Path:   "post.json",
		},
		PasswordSalt: "So-I-Heard-You-Like-Mupkids-?--%s--",
		Limit:        1000,
		PageTag:      "page",
		Parse:        parse,
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
			FileSize:  p.FileSize,
			Width:     p.Width,
			Height:    p.Height,
			CreatedAt: p.parseTime(),
			Author:    p.Author,
			Source:    p.Source,
			Rating:    p.Rating,
			Hash:      p.Md5,
			Tags:      p.parseTags(),
			Score:     p.Score,
		}
	}

	return result, nil
}
