package board

import (
	"net/http"
	"net/url"
	"time"

	"github.com/leonidboykov/getmoe/conf"
	"github.com/leonidboykov/getmoe/provider/danbooru"
	"github.com/leonidboykov/getmoe/provider/gelbooru"
	"github.com/leonidboykov/getmoe/provider/moebooru"
	"github.com/leonidboykov/getmoe/provider/sankaku"
)

// AvailableBoards ...
var AvailableBoards = map[string]Board{
	"yande.re": Board{
		Provider: moebooru.New(conf.ProviderConfiguration{
			URL: conf.URLString{url.URL{Host: "yande.re"}},
		}),
		httpClient: &http.Client{Timeout: 30 * time.Second},
	},
	"konachan.com": Board{
		Provider: moebooru.New(conf.ProviderConfiguration{
			URL:          conf.URLString{url.URL{Host: "konachan.com"}},
			PasswordSalt: "So-I-Heard-You-Like-Mupkids-?--%s--",
		}),
		httpClient: &http.Client{Timeout: 30 * time.Second},
	},
	"gelbooru.com": Board{
		Provider: gelbooru.New(conf.ProviderConfiguration{
			URL: conf.URLString{url.URL{Host: "gelbooru.com"}},
		}),
		httpClient: &http.Client{Timeout: 30 * time.Second},
	},
	"danbooru.donmai.us": Board{
		Provider: danbooru.New(conf.ProviderConfiguration{
			URL: conf.URLString{url.URL{Host: "danbooru.donmai.us"}},
		}),
		httpClient: &http.Client{Timeout: 30 * time.Second},
	},
	"chan.sankakucomplex.com": Board{
		Provider: sankaku.New(conf.ProviderConfiguration{
			URL: conf.URLString{url.URL{Host: "capi-beta.sankakucomplex.com"}},
		}),
		httpClient: &http.Client{Timeout: 30 * time.Second},
	},
	"idol.sankakucomplex.com": Board{
		Provider: sankaku.New(conf.ProviderConfiguration{
			URL: conf.URLString{url.URL{Host: "iapi.sankakucomplex.com"}},
		}),
		httpClient: &http.Client{Timeout: 30 * time.Second},
	},
}
