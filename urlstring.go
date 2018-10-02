package getmoe

import "net/url"

// URLString provides a helper to parse string as url.URL
type URLString struct {
	url.URL
}

// UnmarshalYAML implements unmarshaller interface for YAML
func (f *URLString) UnmarshalYAML(unmashal func(interface{}) error) error {
	var s string
	if err := unmashal(&s); err != nil {
		return err
	}
	u, err := url.Parse(s)
	if err != nil {
		return err
	}
	f.URL = *u
	return nil
}
