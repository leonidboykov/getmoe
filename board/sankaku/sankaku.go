/*
Package sankaku implements a simple library for accessing Sankakucomplex-based
image boards.

Default configurations are available for the following websites

	chan.sankakucomplex.com
	idol.sankakucomplex.com

*/
package sankaku

import (
	"encoding/json"
	"net/url"

	"github.com/leonidboykov/getmoe"
)

const (
	autosuggestHost = "https://ias.sankakucomplex.com" // /tag/autosuggest?tag={tag}
	// https://cas.sankakucomplex.com ?? alternative
)

var (
	// ChanSankakuConfig preconfigured config for Sankaku Channel site
	ChanSankakuConfig = getmoe.Board{
		URL: url.URL{
			Scheme: "https",
			Host:   "capi.sankakucomplex.com",
			Path:   "post/index.json",
		},
		PasswordSalt: "choujin-steiner--%s--",
		Limit:        100,
		UserAgent:    "SCChannelApp/2.7 (Android; black)",
		AppkeySalt:   "sankakuapp_%s_Z5NE9YASej",
		PageTag:      "page",
		Parse:        parse,
	}
	// IdolSankakuConfig preconfigured config for Sankaku Idol site
	IdolSankakuConfig = getmoe.Board{
		URL: url.URL{
			Scheme: "https",
			Host:   "iapi.sankakucomplex.com",
			Path:   "post/index.json",
		},
		PasswordSalt: "choujin-steiner--%s--",
		Limit:        100,
		UserAgent:    "SCChannelApp/2.7 (Android; idol)",
		AppkeySalt:   "sankakuapp_%s_Z5NE9YASej",
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
			Score:     p.TotalScore,
		}
	}

	return result, nil
}
