package gelbooru

import "github.com/leonidboykov/getmoe"

func init() {
	getmoe.RegisterProvider(providerName, New)

	// Default providers.
	getmoe.RegisterSettings("gelbooru.com", &getmoe.ProviderConfiguration{
		Name: providerName,
		URL:  "https://gelbooru.com/",
	})
}
