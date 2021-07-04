package sankaku

import (
	"encoding/json"
	"testing"
	"time"
)

func TestJSONTime_UnmarshalJSON(t *testing.T) {
	tt := []struct {
		name string
		args string
		want time.Time
		err  string
	}{
		{
			name: "success",
			args: `{"s": 1612256246, "n": 0}`,
			want: time.Date(2021, 2, 2, 8, 57, 26, 0, time.UTC),
			err:  "",
		},
		{
			name: "unmarshal error",
			args: `{"s": "1612256246", "n": "0"}`,
			want: time.Time{},
			err:  "json: cannot unmarshal string into Go struct field .s of type int64",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var jt jsonTime
			err := json.Unmarshal([]byte(tc.args), &jt)
			if err != nil && err.Error() != tc.err {
				t.Fatalf("expected err %v, got %v", tc.err, err)
			}
			if !jt.Time.Equal(tc.want) {
				t.Fatalf("expected %v, got %v", tc.want, jt)
			}
		})
	}
}
