package provider

import (
	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/provider/danbooru"
	"github.com/leonidboykov/getmoe/provider/gelbooru"
	"github.com/leonidboykov/getmoe/provider/moebooru"
	"github.com/leonidboykov/getmoe/provider/sankaku"
)

type factory func(getmoe.ProviderConfiguration) getmoe.Provider

func newDanbooru(config getmoe.ProviderConfiguration) getmoe.Provider {
	return danbooru.New(config)
}

func newGelbooru(config getmoe.ProviderConfiguration) getmoe.Provider {
	return gelbooru.New(config)
}

func newMoebooru(config getmoe.ProviderConfiguration) getmoe.Provider {
	return moebooru.New(config)
}

func newSankaku(config getmoe.ProviderConfiguration) getmoe.Provider {
	return sankaku.New(config)
}

// Providers is a list of avalilable providers
var Providers = map[string]factory{
	"danbooru": newDanbooru,
	"gelbooru": newGelbooru,
	"moebooru": newMoebooru,
	"sankaku":  newSankaku,
}
