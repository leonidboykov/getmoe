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
	"net/url"

	"github.com/leonidboykov/getmoe"
)

var (
	// YandeReConfig preconfigured config for yande.re site
	YandeReConfig = getmoe.Board{
		URL: url.URL{
			Scheme: "https",
			Host:   "yande.re",
			Path:   "post.json", // TODO: Make this arg more configurable
		},
		PasswordSalt: "choujin-steiner--%s--",
		Limit:        100,
	}
	// KonachanConfig preconfigured config for konachan.com site
	KonachanConfig = getmoe.Board{
		URL: url.URL{
			Scheme: "https",
			Host:   "konachan.com",
			Path:   "post.json", // TODO: Make this arg more configurable
		},
		PasswordSalt: "So-I-Heard-You-Like-Mupkids-?--%s--",
		Limit:        100,
	}
)

// GetConfig by name
func GetConfig(url string) getmoe.Board {
	switch url {
	case "yande.re":
		return YandeReConfig
	case "konachan.com":
		return KonachanConfig
	default:
		return getmoe.Board{}
	}
}
