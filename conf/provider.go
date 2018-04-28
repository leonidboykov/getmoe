package conf

import (
	"encoding/json"
	"net/url"
)

// ProviderConfiguration holds provider related configuration
type ProviderConfiguration struct {
	Name     string            `json:"name" yaml:"name"`
	Provider string            `json:"provider" yaml:"provider"`
	Host     url.URL           `json:"host" yaml:"host"`
	Auth     AuthConfiguration `json:"auth" yaml:",inline"`
	Headers  []Header          `json:"headers" yaml:"headers"`
}

// UnmarshalJSON implements custom JSON unmarshaler for parsing URL to string
func (c *ProviderConfiguration) UnmarshalJSON(data []byte) error {
	type Alias ProviderConfiguration
	aux := &struct {
		Host string `json:"host"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	u, err := url.Parse(aux.Host)
	if err != nil {
		return err
	}
	c.Host = *u
	return nil
}
