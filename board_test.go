package getmoe

import (
	"testing"
)

func TestSha1(t *testing.T) {
	var tests = []struct {
		value  string
		salt   string
		result string
	}{
		{"123456789", "choujin-steiner--%s--", "a082648f7bcc1e40b5d562fa9b808689a36ca6be"},
	}

	for _, test := range tests {
		v := Sha1(test.value, test.salt)
		if v != test.result {
			t.Error(
				"For", test.value, "and", test.salt,
				"expected", test.result,
				"got", v,
			)
		}
	}
}

func TestBuildAuth(t *testing.T) {
	var login, password string = "user", "123456789"
	var tests = []struct {
		config Board
		result string
	}{
		{Board{PasswordSalt: "choujin-steiner--%s--"}, "?login=user&password_hash=a082648f7bcc1e40b5d562fa9b808689a36ca6be"},
		{Board{PasswordSalt: "choujin-steiner--%s--", AppkeySalt: "sankakuapp_%s_Z5NE9YASej"}, "?appkey=bf7420a71090010192df8751d8f5504cde002be1&login=user&password_hash=a082648f7bcc1e40b5d562fa9b808689a36ca6be"},
	}

	for _, test := range tests {
		test.config.BuildAuth(login, password)
		e := test.config.URL.String()
		if e != test.result {
			t.Error(
				"For", test.config.PasswordSalt, "and", test.config.AppkeySalt,
				"expected", test.result,
				"got", e,
			)
		}
	}
}
