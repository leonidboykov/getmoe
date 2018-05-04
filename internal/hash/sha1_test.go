package hash_test

import (
	"testing"

	"github.com/leonidboykov/getmoe/internal/hash"
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
		v := hash.Sha1(test.value, test.salt)
		if v != test.result {
			t.Error(
				"For", test.value, "and", test.salt,
				"expected", test.result,
				"got", v,
			)
		}
	}
}
