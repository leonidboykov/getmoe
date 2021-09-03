package getmoe_test

import (
	"net/url"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/leonidboykov/getmoe"
)

type TestURLString struct {
	Host getmoe.URLString `yaml:"host"`
}

func TestUnmarshalYAML(t *testing.T) {
	tt := []struct {
		name string
		args string
		want getmoe.URLString `yaml:"host"`
		err  string
	}{
		{
			name: "success",
			args: "host: https://example.com",
			want: getmoe.URLString{URL: url.URL{Scheme: "https", Host: "example.com"}},
			err:  "",
		},
		{
			name: "empty",
			args: "",
			want: getmoe.URLString{URL: url.URL{}},
			err:  "",
		},
		{
			name: "url parse error",
			args: "host: ':'",
			want: getmoe.URLString{URL: url.URL{}},
			err:  `parse ":": missing protocol scheme`,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var field TestURLString
			err := yaml.Unmarshal([]byte(tc.args), &field)
			if err != nil && err.Error() != tc.err {
				t.Fatalf("expected err %v, got %v", tc.err, err)
			}
			if field.Host.String() != tc.want.String() {
				t.Errorf("Unmarshal(%s) == %s, want %s", tc.args, field.Host.String(), tc.want.String())
			}
		})
	}
}
