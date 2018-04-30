package conf_test

import (
	"net/url"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/leonidboykov/getmoe/conf"
)

type TestURLString struct {
	Host conf.URLString `yaml:"host"`
}

func TestUnmarshalYAML(t *testing.T) {
	tests := []struct {
		in   string
		want conf.URLString `yaml:"host"`
	}{
		{"host: https://example.com", conf.URLString{URL: url.URL{Scheme: "https", Host: "example.com"}}},
	}
	for _, test := range tests {
		var field TestURLString
		if err := yaml.Unmarshal([]byte(test.in), &field); err != nil {
			t.Error(err)
		}
		if field.Host.String() != test.want.String() {
			t.Errorf("Unmarshal(%s) == %s, want %s", test.in, field.Host.String(), test.want.String())
		}
	}
}
