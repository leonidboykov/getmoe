package getmoe

import (
	"strings"

	"gopkg.in/yaml.v3"
)

// stringOrSlice allows to use single string syntax for slices with one element:
//  key: value
// and
//  key: [ value1, value2 ]
type stringOrSlice []string

// UnmarshalYAML implements unmarshaller interface for YAML.
func (f *stringOrSlice) UnmarshalYAML(value *yaml.Node) error {
	var slice []string
	if err := value.Decode(&slice); err != nil {
		var str string
		if err := value.Decode(&str); err != nil {
			return err
		}
		*f = strings.Split(str, " ")
		return nil
	}
	*f = slice
	return nil
}

type includeExclude struct {
	Include stringOrSlice `yaml:"include"`
	Exclude stringOrSlice `yaml:"exclude"`
}

func (f *includeExclude) UnmarshalYAML(value *yaml.Node) error {
	ie := struct {
		Include stringOrSlice `yaml:"include"`
		Exclude stringOrSlice `yaml:"exclude"`
	}{}
	if err := value.Decode(&ie); err != nil {
		var ss stringOrSlice
		if err := value.Decode(&ss); err != nil {
			return err
		}
		f.Include = ss
		return nil
	}
	f.Include = ie.Include
	f.Exclude = ie.Exclude
	return nil
}

func (f *includeExclude) StringWithPrefix(prefix string) string {
	var tags []string
	if prefix != "" {
		prefix += ":"
	}
	for i := range f.Include {
		tags = append(tags, prefix+f.Include[i])
	}
	for i := range f.Exclude {
		tags = append(tags, "-"+prefix+f.Exclude[i])
	}
	return strings.Join(tags, " ")
}
