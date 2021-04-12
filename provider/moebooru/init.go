package moebooru

import "github.com/leonidboykov/getmoe"

func init() {
	getmoe.RegisterProvider(providerName, New)

	// Default providers.
	// getmoe.RegisterSettings("yande.re", &getmoe.ProviderConfiguration{
	// 	Name: providerName,
	// 	URL:  "https://yande.re",
	// })
	getmoe.RegisterSettings("konachan.com", &getmoe.ProviderConfiguration{
		Name:         providerName,
		URL:          "https://konachan.com",
		PasswordSalt: "So-I-Heard-You-Like-Mupkids-?--%s--",
	})
}
