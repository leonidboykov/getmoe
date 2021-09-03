package getmoe

import (
	"os"

	"gopkg.in/yaml.v3"
)

// RootConfiguration provides root configuration.
type RootConfiguration struct {
	Boards   map[string]BoardConfiguration `yaml:"boards"`
	Download DownloadConfiguration         `yaml:"download"`
}

// BoardConfiguration holds board related configuration.
type BoardConfiguration struct {
	Settings string                `yaml:"settings"`
	Provider ProviderConfiguration `yaml:",inline"`
}

// ProviderConfiguration holds provider configuration.
type ProviderConfiguration struct {
	Name         string      `yaml:"provider"`
	URL          string      `yaml:"url"`
	PasswordSalt string      `yaml:"password_salt"`
	PostsLimit   int         `yaml:"posts_limits"`
	Credentials  Credentials `yaml:"credentials"`
}

// Credentials holds information for authenticating.
type Credentials struct {
	Login          string `yaml:"login"`
	UserID         int    `yaml:"user_id"`
	Password       string `yaml:"password"`
	HashedPassword string `yaml:"hashed_password"`
	APIKey         string `yaml:"apikey"`
}

// DownloadConfiguration holds download related configuration.
type DownloadConfiguration struct {
	Searches []SearchConfiguration `yaml:"searches"`
	Filters  stringOrSlice         `yaml:"filters"`
	SavePath string                `yaml:"save_to"`
}

// SearchConfiguration holds confguration about searches.
type SearchConfiguration struct {
	Tags   Tags          `yaml:"search"`
	Boards stringOrSlice `yaml:"boards"`
}

// ReadConfiguraton reads a root configuration from file.
func ReadConfiguraton(fname string) (*RootConfiguration, error) {
	var config RootConfiguration
	f, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(f, &config); err != nil {
		return nil, err
	}

	return &config, err
}
