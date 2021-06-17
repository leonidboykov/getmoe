package danbooru

import "github.com/leonidboykov/getmoe"

func init() {
	getmoe.RegisterProvider(providerName, New)

	// Default providers.
	getmoe.RegisterPresets("danbooru.donmai.us", &getmoe.ProviderConfiguration{
		Name: providerName,
		URL:  "https://danbooru.donmai.us/",
	})
}
