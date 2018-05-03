package provider

import (
	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/conf"
	"github.com/leonidboykov/getmoe/provider/moebooru"
)

type factory func(conf.ProviderConfiguration) getmoe.Provider

func newMoebooru(config conf.ProviderConfiguration) getmoe.Provider {
	return moebooru.New(config)
}

// Providers is a list of avalilable providers
var Providers = map[string]factory{
	"moebooru": newMoebooru,
}
