/*
Package sankaku implements a simple library for accessing Sankakucomplex-based
image boards.

Default configurations are available for the following websites

	chan.sankakucomplex.com
	idol.sankakucomplex.com

*/
package sankaku

import (
	"github.com/leonidboykov/getmoe"
)

var supported = [...]string{"chan.sankakucomplex.com", "idol.sankakucomplex.com"}

const (
	autosuggestHost = "https://ias.sankakucomplex.com" // /tag/autosuggest?tag={tag}
	// https://cas.sankakucomplex.com ?? alternative
)

var (
	// ChanSankakuConfig preconfigured config for Sankaku Channel site
	ChanSankakuConfig = getmoe.Config{
		BaseURL:      "https://capi.sankakucomplex.com",
		PasswordSalt: "choujin-steiner--%s--",
		Limit:        100,
		UserAgent:    "SCChannelApp/2.7 (Android; black)",
		AppkeySalt:   "sankakuapp_%s_Z5NE9YASej",
	}
	// IdolSankakuConfig preconfigured config for Sankaku Idol site
	IdolSankakuConfig = getmoe.Config{
		BaseURL:      "https://iapi.sankakucomplex.com",
		PasswordSalt: "choujin-steiner--%s--",
		Limit:        100,
		UserAgent:    "SCChannelApp/2.7 (Android; idol)",
		AppkeySalt:   "sankakuapp_%s_Z5NE9YASej",
	}
)
