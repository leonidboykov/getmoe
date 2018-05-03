package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// GlobalConfiguration provides global configuration
type GlobalConfiguration struct {
	Boards   map[string]BoardConfiguration `yaml:"boards"`
	Download DownloadConfiguration         `yaml:"download"`
}

// AuthConfiguration provides configuration for authenticating
type AuthConfiguration struct {
	Login          string `yaml:"login"`
	Password       string `yaml:"password"`
	HashedPassword string `yaml:"hashed_password"`
}

// BoardConfiguration holds board related configuration
type BoardConfiguration struct {
	// Name string `yaml:"name"`
	// URL  string `yaml:"url"`
	// Auth     AuthConfiguration     `yaml:",inline"`

	Provider ProviderConfiguration `yaml:",inline"`
	Headers  map[string]string     `yaml:"headers"`
}

// ProviderConfiguration holds provider related configuration
type ProviderConfiguration struct {
	Name         string            `yaml:"provider"`
	URL          URLString         `yaml:"url"`
	Auth         AuthConfiguration `yaml:",inline"`
	PasswordSalt string            `yaml:"password_salt"`
	AppkeySalt   string            `yaml:"appkey_salt"`
	PostsLimit   int               `yaml:"posts_limit"`
}

// DownloadConfiguration holds download related configuration
type DownloadConfiguration struct {
	Request RequestConfiguration `yaml:",inline"`
}

// RequestConfiguration holds request related configuration
type RequestConfiguration struct {
	Tags []string `yaml:"tags"`
}

// Load loads global configuration
func Load(filename string) (*GlobalConfiguration, error) {
	config := new(GlobalConfiguration)
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(f, &config); err != nil {
		return nil, err
	}

	return config, nil
}
