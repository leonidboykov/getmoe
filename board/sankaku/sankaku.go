/*
Package sankaku implements a simple library for accessing Sankakucomplex-based
image boards.

Default configurations are available for the following websites

	chan.sankakucomplex.com
	idol.sankakucomplex.com

*/
package sankaku

import (
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
			Path:   "post/index.json", // TODO: Make this arg more configurable
		},
		PasswordSalt: "choujin-steiner--%s--",
		Limit:        100,
		UserAgent:    "SCChannelApp/2.7 (Android; black)",
		AppkeySalt:   "sankakuapp_%s_Z5NE9YASej",
	}
	// IdolSankakuConfig preconfigured config for Sankaku Idol site
	IdolSankakuConfig = getmoe.Board{
		URL: url.URL{
			Scheme: "https",
			Host:   "iapi.sankakucomplex.com",
			Path:   "post/index.json", // TODO: Make this arg more configurable
		},
		PasswordSalt: "choujin-steiner--%s--",
		Limit:        100,
		UserAgent:    "SCChannelApp/2.7 (Android; idol)",
		AppkeySalt:   "sankakuapp_%s_Z5NE9YASej",
	}
)

// GetConfig by name
func GetConfig(url string) getmoe.Board {
	switch url {
	case "chan.sankakucomplex.com":
		return ChanSankakuConfig
	case "idol.sankakucomplex.com":
		return IdolSankakuConfig
	default:
		return getmoe.Board{}
	}
}
