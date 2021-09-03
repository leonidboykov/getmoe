package getmoe

import (
	"fmt"
	"sync"

	"github.com/imdario/mergo"
)

// presets contains default configurations for providers.
var (
	presetsMu sync.RWMutex
	presets   = make(map[string]*ProviderConfiguration)
)

// RegisterPresets registers pre-defined settings for boards.
func RegisterPresets(name string, config *ProviderConfiguration) {
	presetsMu.Lock()
	defer presetsMu.Unlock()
	if config == nil {
		panic("getmoe: unable to register a nil board presets")
	}
	if _, dup := presets[name]; dup {
		panic("getmoe: presets '" + name + "' is already registered")
	}
	presets[name] = config
}

// applyPresets applies presets to provider.
func applyPresets(name string, config *ProviderConfiguration) error {
	if defaultPresets, ok := presets[name]; ok {
		mergo.Merge(config, defaultPresets)
		return nil
	}
	return fmt.Errorf("unknown presets '%s'", name)
}
