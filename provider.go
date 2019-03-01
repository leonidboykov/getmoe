package getmoe

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	providersMu sync.RWMutex
	providers   = make(map[string]Provider)
)

// Provider describes Board provider
type Provider interface {
	New(ProviderConfiguration)
	Auth(AuthConfiguration)
	BuildRequest(RequestConfiguration)
	NextPage()
	PageRequest() (*http.Request, error)
	Parse([]byte) ([]Post, error)
}

// RegisterProvider registers booru provider
func RegisterProvider(name string, provider Provider) {
	providersMu.Lock()
	defer providersMu.Unlock()
	if provider == nil {
		panic("getmoe: unable to register a nil provider")
	}
	if _, dup := providers[name]; dup {
		panic("getmoe: provider '" + name + "' is already registered")
	}
	providers[name] = provider
}

// NewProvider creates a new provider
func NewProvider(providerName string, conf ProviderConfiguration) (*Provider, error) {
	providersMu.RLock()
	provider, ok := providers[providerName]
	providersMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("getmoe: unknown provider %s", providerName)
	}
	provider.New(conf)
	return &provider, nil
}
