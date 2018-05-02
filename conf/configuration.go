package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// GlobalConfiguration provides global configuration
type GlobalConfiguration struct {
	Boards   map[string]BoardConfiguration `yaml:"boards"`
	Download DownloadConfiguration         `yaml:"download"`
	// Boards    struct {
	// 	Auth AuthConfiguration `yaml:",inline"`
	// } `yaml:"boards"`
}

// AuthConfiguration provides configuration for authenticating
type AuthConfiguration struct {
	Login          string `yaml:"login"`
	Password       string `yaml:"password"`
	HashedPassword string `yaml:"hashed_password"`
	PasswordSalt   string `yaml:"password_salt"`
}

// BoardConfiguration holds provider related configuration
type BoardConfiguration struct {
	Name     string            `yaml:"name"`
	Provider string            `yaml:"provider"`
	URL      string            `yaml:"url"`
	Auth     AuthConfiguration `yaml:",inline"`
	Headers  []Header          `yaml:"headers"`
	// Host     URLString         `yaml:"host"`
}

// Header holds custom header for HTTP client
type Header map[string]string

// DownloadConfiguration holds download related configuration
type DownloadConfiguration struct {
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
