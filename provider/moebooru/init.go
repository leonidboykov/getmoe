package moebooru

import "github.com/leonidboykov/getmoe"

func init() {
	getmoe.RegisterProvider(providerName, New)

	// Default providers.
	getmoe.RegisterPresets("yande.re", &getmoe.ProviderConfiguration{
		Name: providerName,
		URL:  "https://yande.re",
	})
	getmoe.RegisterPresets("konachan.com", &getmoe.ProviderConfiguration{
		Name:         providerName,
		URL:          "https://konachan.com",
		PasswordSalt: "So-I-Heard-You-Like-Mupkids-?--%s--",
	})
}
