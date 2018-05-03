package board

import (
	"net/url"

	"github.com/leonidboykov/getmoe/provider/moebooru"
)

// AvailableBoards ...
// var AvailableBoards = map[string]Board{
// 	"yande.re":                moebooru.YandeReConfig,
// 	"konachan.com":            moebooru.KonachanConfig,
// 	"gelbooru.com":            gelbooru.GelbooruConfig,
// 	"danbooru.donmai.us":      danbooru.DanbooruDonmaiUsConfig,
// 	"chan.sankakucomplex.com": sankaku.ChanSankakuConfig,
// 	"idol.sankakucomplex.com": sankaku.IdolSankakuConfig,
// }

// AvailableBoards ...
var AvailableBoards = map[string]Board{
	"yande.re": Board{
		Provider: &moebooru.Provider{
			URL: &url.URL{
				Scheme: "https",
				Host:   "yande.re",
			},
		},
	},
	"konachan.com": Board{
		Provider: &moebooru.Provider{
			URL: &url.URL{
				Scheme: "https",
				Host:   "konachan.com",
			},
			PasswordSalt: "So-I-Heard-You-Like-Mupkids-?--%s--",
		},
	},
}
