package provider

import (
	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/conf"
	"github.com/leonidboykov/getmoe/provider/danbooru"
	"github.com/leonidboykov/getmoe/provider/gelbooru"
	"github.com/leonidboykov/getmoe/provider/moebooru"
	"github.com/leonidboykov/getmoe/provider/sankaku"
)

type factory func(conf.ProviderConfiguration) getmoe.Provider

func newDanbooru(config conf.ProviderConfiguration) getmoe.Provider {
	return danbooru.New(config)
}

func newGelbooru(config conf.ProviderConfiguration) getmoe.Provider {
	return gelbooru.New(config)
}

func newMoebooru(config conf.ProviderConfiguration) getmoe.Provider {
	return moebooru.New(config)
}

func newSankaku(config conf.ProviderConfiguration) getmoe.Provider {
	return sankaku.New(config)
}

// Providers is a list of avalilable providers
var Providers = map[string]factory{
	"danbooru": newDanbooru,
	"gelbooru": newGelbooru,
	"moebooru": newMoebooru,
	"sankaku":  newSankaku,
}
