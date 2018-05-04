package conf_test

import (
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/leonidboykov/getmoe/conf"
)

var testProviders = []byte(`
providers:
  login: globalUser
  password: globalPassword
  yande.re:
    login: username
    password: password
    provider: moebooru
    host: 'https://yande.re/'
  sankaku:
    login: username2
    password: password2
    provider: sankaku
`)

func TestProviders(t *testing.T) {
	var config conf.GlobalConfiguration
	if err := yaml.Unmarshal(testProviders, &config); err != nil {
		t.Error(err)
	}
}
