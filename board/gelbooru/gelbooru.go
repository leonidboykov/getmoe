/*
Package gelbooru implements a simple library for accessing Gelbooru-based image
boards.

Source code of Moebooru is available at https://github.com/moebooru/moebooru

Default configurations are available for the following websites

	gelbooru

*/
package gelbooru

import (
	"net/url"

	"github.com/leonidboykov/getmoe"
)

var (
	// GelbooruConfig preconfigured config for Sankaku Channel site
	GelbooruConfig = getmoe.Board{
		URL: url.URL{
			Scheme:   "https",
			Host:     "gelbooru.com",
			Path:     "index.php",
			RawQuery: "page=dapi&s=post&q=index&json=1",
		},
		Limit:    1000,
		HasPages: false,
	}
)
