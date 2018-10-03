package getmoe_test

import (
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/leonidboykov/getmoe"
)

var testProviders = []byte(`
boards:
  yande.re:
    login: username
    password: password
    provider: moebooru
    host: 'https://yande.re/'
  sankaku:
    login: username2
    password: password2
    provider: sankaku
download:
  tags: [ "tag1", "tag2" ]
`)

func TestProviders(t *testing.T) {
	var config getmoe.GlobalConfiguration
	if err := yaml.Unmarshal(testProviders, &config); err != nil {
		t.Error(err)
	}
}
