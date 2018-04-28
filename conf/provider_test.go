package conf_test

import (
	"encoding/json"
	"net/url"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/leonidboykov/getmoe/conf"
)

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		sample   []byte
		expected url.URL
	}{
		{sample: []byte("{\"host\": \"https://example.com\"}"), expected: url.URL{Scheme: "https", Host: "example.com"}},
	}

	for _, test := range tests {
		var result conf.ProviderConfiguration
		if err := json.Unmarshal(test.sample, &result); err != nil {
			t.Error(err)
		}
		if result.Host.String() != test.expected.String() {
			t.Errorf("For %s wanted %s got %s", string(test.sample), test.expected.String(), result.Host.String())
		}
	}
}

func TestUnmarshalYAML(t *testing.T) {
	tests := []struct {
		sample   []byte
		expected url.URL
	}{
		{sample: []byte("host: https://example.com"), expected: url.URL{Scheme: "https", Host: "example.com"}},
	}

	for _, test := range tests {
		var result conf.ProviderConfiguration
		if err := yaml.Unmarshal(test.sample, &result); err != nil {
			t.Error(err)
		}
		if result.Host.String() != test.expected.String() {
			t.Errorf("For %s wanted %s got %s", string(test.sample), test.expected.String(), result.Host.String())
		}
	}
}
