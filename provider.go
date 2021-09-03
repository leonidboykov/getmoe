package getmoe

import (
	"fmt"
	"sync"
)

// Provider describes a board provider.
type Provider interface {
	RequestPage(tags Tags, page int) ([]Post, error)
}

// ProviderFactory represents provider constructor.
type ProviderFactory func(ProviderConfiguration) Provider

var (
	providersMu sync.RWMutex
	providers   = make(map[string]ProviderFactory)
)

// RegisterProvider registers provider constructor to factory.
func RegisterProvider(name string, provider ProviderFactory) {
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

// NewProvider construcs a new provider with congifuration.
func NewProvider(name string, config ProviderConfiguration) (Provider, error) {
	providersMu.RLock()
	defer providersMu.RUnlock()
	provider, ok := providers[name]
	if !ok {
		return nil, fmt.Errorf("getmoe: unknown provider '%s'", name)
	}
	return provider(config), nil
}
