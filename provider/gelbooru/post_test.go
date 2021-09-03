package gelbooru

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestTags_UnmarshalJSON(t *testing.T) {
	tt := []struct {
		name string
		args string
		want tags
		err  string
	}{
		{
			name: "success",
			args: `"tag1 tag2 tag3"`,
			want: []string{"tag1", "tag2", "tag3"},
		},
		{
			name: "empty",
			args: `""`,
			want: []string{},
		},
		{
			name: "not string",
			args: "1",
			err:  "invalid syntax",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var got tags
			err := json.Unmarshal([]byte(tc.args), &got)
			if err != nil && err.Error() != tc.err {
				t.Fatalf("unexpected err %v", err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestUnixtime_UnmarshalJSON(t *testing.T) {
	tt := []struct {
		name string
		args string
		want time.Time
		err  string
	}{
		{
			name: "success",
			args: "1624321618",
			want: time.Date(2021, 06, 22, 00, 26, 58, 0, time.UTC),
		},
		{
			name: "error",
			args: `"not_string"`,
			want: time.Time{},
			err:  `strconv.ParseInt: parsing "\"not_string\"": invalid syntax`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var got unixtime
			err := json.Unmarshal([]byte(tc.args), &got)
			if err != nil && err.Error() != tc.err {
				t.Fatalf("expected err %v, got %v", tc.err, err)
			}
			if !tc.want.Equal(time.Time(got)) {
				t.Fatalf("expected %v, got %v", tc.want, got)
			}
		})
	}
}
