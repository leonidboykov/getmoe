/*
Package danbooru implements a simple library for accessing Danbooru-based image
boards.

Source code of Moebooru is available at https://github.com/r888888888/danbooru

Default configurations are available for the following websites

	danbooru.donmai.us

*/
package danbooru

import (
	"net/url"

	"github.com/leonidboykov/getmoe"
)

var (
	// DanbooruDonmaiUsConfig preconfigured config for Sankaku Channel site
	DanbooruDonmaiUsConfig = getmoe.Board{
		URL: url.URL{
			Scheme: "https",
			Host:   "danbooru.donmai.us",
			Path:   "posts.json",
		},
		Limit:    200,
		HasPages: true,
	}
)
