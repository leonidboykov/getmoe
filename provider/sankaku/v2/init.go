package sankaku

import "github.com/leonidboykov/getmoe"

func init() {
	getmoe.RegisterProvider(providerName, New)

	// Default providers.
	getmoe.RegisterPresets("chan.sankakucomplex.com", &getmoe.ProviderConfiguration{
		Name: providerName,
		URL:  "https://capi-v2.sankakucomplex.com",
	})
}
