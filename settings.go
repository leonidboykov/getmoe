package getmoe

import (
	"fmt"
	"sync"

	"github.com/imdario/mergo"
)

// settings contains default configurations for providers.
var (
	settingsMu sync.RWMutex
	settings   = make(map[string]*ProviderConfiguration)
)

// RegisterSettings registers pre-defined settings for boards.
func RegisterSettings(name string, config *ProviderConfiguration) {
	settingsMu.Lock()
	defer settingsMu.Unlock()
	if config == nil {
		panic("getmoe: unable to register a nil board settings")
	}
	if _, dup := settings[name]; dup {
		panic("getmoe: settings '" + name + "' is already registered")
	}
	settings[name] = config
}

// applySettings applies settings to provider.
func applySettings(name string, config *ProviderConfiguration) error {
	if defaultSettings, ok := settings[name]; ok {
		mergo.Merge(config, defaultSettings)
		return nil
	}
	return fmt.Errorf("getmoe: unknown settings '%s'", name)
}
