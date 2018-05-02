package provider

import (
	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/provider/moebooru"
)

// Providers is a list of avalilable providers
var Providers = map[string]getmoe.Provider{
	"moebooru": &moebooru.Provider{},
}
