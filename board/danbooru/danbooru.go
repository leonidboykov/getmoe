/*
Package danbooru implements a simple library for accessing Danbooru-based image
boards.

Source code of Moebooru is available at https://github.com/r888888888/danbooru

Default configurations are available for the following websites

	danbooru.donmai.us

*/
package danbooru

import (
	"encoding/json"
	"net/url"

	"github.com/leonidboykov/getmoe"
)

var (
	// DanbooruDonmaiUsConfig preconfigured config for DanbooruDonmaiUs site
	DanbooruDonmaiUsConfig = getmoe.Board{
		URL: url.URL{
			Scheme: "https",
			Host:   "danbooru.donmai.us",
			Path:   "posts.json",
		},
		Limit:   200,
		PageTag: "page",
		Parse:   parse,
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
			Width:     page[i].ImageWidth,
			Height:    page[i].ImageHeight,
			CreatedAt: page[i].CreatedAt,
			Author:    page[i].TagStringArtist,
			Source:    page[i].Source,
			Rating:    page[i].Rating,
			Hash:      page[i].Md5,
			Tags:      page[i].parseTags(),
			Score:     page[i].Score,
		}
	}

	return result, nil
}
