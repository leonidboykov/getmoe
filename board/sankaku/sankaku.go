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
		UserAgent:    "SCChannelApp/2.8 (Android; black)",
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
		UserAgent:    "SCChannelApp/2.8 (Android; idol)",
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
			Score:     page[i].TotalScore,
		}
	}

	return result, nil
}
