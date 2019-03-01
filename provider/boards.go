package provider

import (
	"net/url"

	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/provider/danbooru"
	"github.com/leonidboykov/getmoe/provider/gelbooru"
	"github.com/leonidboykov/getmoe/provider/moebooru"
	"github.com/leonidboykov/getmoe/provider/sankaku"
)

// AvailableBoards is a list of predefined boards.
var AvailableBoards = map[string]*getmoe.Board{
	"yande.re": getmoe.NewBoardWithProvider(moebooru.New(getmoe.ProviderConfiguration{
		URL: getmoe.URLString{URL: url.URL{Host: "yande.re"}},
	})),
	"konachan.com": getmoe.NewBoardWithProvider(moebooru.New(getmoe.ProviderConfiguration{
		URL:          getmoe.URLString{URL: url.URL{Host: "konachan.com"}},
		PasswordSalt: "So-I-Heard-You-Like-Mupkids-?--%s--",
	})),
	"gelbooru.com": getmoe.NewBoardWithProvider(gelbooru.New(getmoe.ProviderConfiguration{
		URL: getmoe.URLString{URL: url.URL{Host: "gelbooru.com"}},
	})),
	"danbooru.donmai.us": getmoe.NewBoardWithProvider(danbooru.New(getmoe.ProviderConfiguration{
		URL: getmoe.URLString{URL: url.URL{Host: "danbooru.donmai.us"}},
	})),
	"chan.sankakucomplex.com": getmoe.NewBoardWithProvider(sankaku.New(getmoe.ProviderConfiguration{
		URL: getmoe.URLString{URL: url.URL{Host: "capi-v2.sankakucomplex.com"}},
	})),
	"idol.sankakucomplex.com": getmoe.NewBoardWithProvider(sankaku.New(getmoe.ProviderConfiguration{
		URL: getmoe.URLString{URL: url.URL{
			Host: "iapi.sankakucomplex.com",
			Path: "post/index.json",
		}},
	})),
}
