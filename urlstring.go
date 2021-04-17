package getmoe

import (
	"net/url"

	"gopkg.in/yaml.v3"
)

// URLString provides a helper to parse string as url.URL.
type URLString struct {
	url.URL
}

// UnmarshalYAML implements unmarshaller interface for YAML
func (f *URLString) UnmarshalYAML(value *yaml.Node) error {
	var s string
	if err := value.Decode(&s); err != nil {
		return err
	}
	u, err := url.Parse(s)
	if err != nil {
		return err
	}
	f.URL = *u
	return nil
}
