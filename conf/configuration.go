package conf

// Configuration provides global configuration
type Configuration struct {
	Boards struct {
		Auth      AuthConfiguration       `yaml:",inline"`
		Providers []ProviderConfiguration `yaml:"providers"`
	} `yaml:"boards"`
}

// Header holds custom header for HTTP client
type Header map[string]string

// AuthConfiguration provides configuration for authenticating
type AuthConfiguration struct {
	Login          string `yaml:"login"`
	Password       string `yaml:"password"`
	HashedPassword string `yaml:"hashed_password"`
}
