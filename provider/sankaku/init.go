package sankaku

import "github.com/leonidboykov/getmoe"

func init() {
	getmoe.RegisterProvider(providerName, New)

	// Default providers.
	getmoe.RegisterSettings("chan.sankakucomplex.com", &getmoe.ProviderConfiguration{
		Name: providerName,
		URL:  "https://capi-v2.sankakucomplex.com",
	})
	getmoe.RegisterSettings("idol.sankakucomplex.com", &getmoe.ProviderConfiguration{
		Name: providerName,
		URL:  "https://iapi.sankakucomplex.com",
	})
}
