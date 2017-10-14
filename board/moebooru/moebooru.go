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

	for i := range page {
		result[i] = getmoe.Post{
			ID:        page[i].ID,
			FileURL:   page[i].FileURL,
			FileSize:  page[i].FileSize,
			Width:     page[i].Width,
			Height:    page[i].Height,
			CreatedAt: page[i].parseTime(),
			Author:    page[i].Author,
			Source:    page[i].Source,
			Rating:    page[i].Rating,
			Hash:      page[i].Md5,
			Tags:      page[i].parseTags(),
			Score:     page[i].Score,
		}
	}

	return result, nil
}
