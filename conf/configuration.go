package conf

// Configuration provides global YAML-based configuration
type Configuration struct {
	Providers ProvidersConfiguration `json:"providers" yaml:"providers"`
}

// ProvidersConfiguration provides configration for all providers
type ProvidersConfiguration struct {
	Auth      AuthConfiguration       `json:"auth" yaml:",inline"`
	Providers []ProviderConfiguration `json:"providers" yaml:"providers"`
}

// ProviderConfiguration holds provider related configuration
type ProviderConfiguration struct {
	Name    string            `json:"name" yaml:"name"`
	Auth    AuthConfiguration `json:"auth" yaml:",inline"`
	Host    string            `json:"host" yaml:"host"`
	Headers []Header          `json:"headers" yaml:"headers"`
}

// Header holds custom header for HTTP client
type Header map[string]string

// AuthConfiguration provides configuration for authenticating
type AuthConfiguration struct {
	Login          string `json:"login" yaml:"login"`
	Password       string `json:"password" yaml:"password"`
	HashedPassword string `json:"hashed_password" yaml:"hashed_password"`
}
