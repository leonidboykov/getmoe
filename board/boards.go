package board

import (
	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/board/danbooru"
	"github.com/leonidboykov/getmoe/board/gelbooru"
	"github.com/leonidboykov/getmoe/board/moebooru"
	"github.com/leonidboykov/getmoe/board/sankaku"
)

// AvailableBoards ...
var AvailableBoards = map[string]getmoe.Board{
	"yande.re":                moebooru.YandeReConfig,
	"konachan.com":            moebooru.KonachanConfig,
	"gelbooru.com":            gelbooru.GelbooruConfig,
	"danbooru.donmai.us":      danbooru.DanbooruDonmaiUsConfig,
	"chan.sankakucomplex.com": sankaku.ChanSankakuConfig,
	"idol.sankakucomplex.com": sankaku.IdolSankakuConfig,
}
