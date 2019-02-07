package helper

import "testing"

func TestFileURLUnescape(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		{"https://files.yande.re/image/c8594617d9261c3e9ce47a32f548c448/yande.re%20411170%20kobayakawa_sae%20mizumoto_yukari%20naked%20nipples%20photoshop%20pussy%20sakuma_mayu%20the_idolm%40ster%20the_idolm%40ster_cinderella_girls%20uncensored.png", "yande.re-411170-kobayakawa_sae-mizumoto_yukari-naked-nipples-photoshop-pussy-sakuma_mayu-the_idolm@ster-the_idolm@ster_cinderella_girls-uncensored.png"},
		{"hello%�world", ""},
	}

	for _, pair := range tests {
		v, err := FileURLUnescape(pair.value)
		if err != nil {
			t.Error(
				"For", pair.value,
				"got error:", err,
			)
		}
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
