package conf

// Configuration provides global configuration
type Configuration struct {
	Boards struct {
		Auth      AuthConfiguration       `json:"auth" yaml:",inline"`
		Providers []ProviderConfiguration `json:"providers" yaml:"providers"`
	} `json:"boards" yaml:"boards"`
}

// Header holds custom header for HTTP client
type Header map[string]string

// AuthConfiguration provides configuration for authenticating
type AuthConfiguration struct {
	Login          string `json:"login" yaml:"login"`
	Password       string `json:"password" yaml:"password"`
	HashedPassword string `json:"hashed_password" yaml:"hashed_password"`
}
