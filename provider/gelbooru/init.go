package gelbooru

import "github.com/leonidboykov/getmoe"

func init() {
	getmoe.RegisterProvider(providerName, New)

	// Default providers.
	getmoe.RegisterPresets("gelbooru.com", &getmoe.ProviderConfiguration{
		Name: providerName,
		URL:  "https://gelbooru.com/",
	})
	getmoe.RegisterPresets("rule34.xxx", &getmoe.ProviderConfiguration{
		Name: providerName,
		URL:  "https://rule34.xxx",
	})
}
