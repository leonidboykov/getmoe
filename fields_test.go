package getmoe

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestStringOrSlice_UnmarshalYAML(t *testing.T) {
	type ts struct {
		Value stringOrSlice `yaml:"key"`
	}
	tt := []struct {
		name string
		args string
		want stringOrSlice
		err  string
	}{
		{
			name: "string",
			args: "key: value",
			want: []string{"value"},
		},
		{
			name: "slice",
			args: "key: [ value1, value2 ]",
			want: []string{"value1", "value2"},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var field ts
			err := yaml.Unmarshal([]byte(tc.args), &field)
			if err != nil && err.Error() != tc.err {
				t.Fatalf("expected err %v, got %v", tc.err, err)
			}
			if !reflect.DeepEqual(field.Value, tc.want) {
				t.Errorf("Unmarshal(%s) == %s, want %s", tc.args, field.Value, tc.want)
			}
		})
	}
}

func TestIncludeExclude_UnmarshalYAML(t *testing.T) {
	type ts struct {
		Value includeExclude `yaml:"key"`
	}
	tt := []struct {
		name string
		args string
		want includeExclude
		err  string
	}{
		{
			name: "string",
			args: "key: value",
			want: includeExclude{
				Include: []string{"value"},
			},
		},
		{
			name: "split string",
			args: "key: value1 value2",
			want: includeExclude{
				Include: []string{"value1", "value2"},
			},
		},
		{
			name: "slice",
			args: "key: [ value1, value2 ]",
			want: includeExclude{
				Include: []string{"value1", "value2"},
			},
		},
		{
			name: "struct",
			args: `key:
  include: [ value1, value2 ]
  exclude: [ value3, value4 ]`,
			want: includeExclude{
				Include: []string{"value1", "value2"},
				Exclude: []string{"value3", "value4"},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var field ts
			err := yaml.Unmarshal([]byte(tc.args), &field)
			if err != nil && err.Error() != tc.err {
				t.Fatalf("expected err %v, got %v", tc.err, err)
			}
			if !reflect.DeepEqual(field.Value.Include, tc.want.Include) {
				t.Errorf("Unmarshal(%s) == %s, want %s", tc.args, field.Value.Include, tc.want.Include)
			}
			if !reflect.DeepEqual(field.Value.Exclude, tc.want.Exclude) {
				t.Errorf("Unmarshal(%s) == %s, want %s", tc.args, field.Value.Exclude, tc.want.Exclude)
			}
		})
	}
}
