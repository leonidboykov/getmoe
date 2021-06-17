package getmoe

import (
	"strings"
	"time"
)

// Boorus rating system has safe, questionable and explicit tags.
const (
	RatingSafe         string = "s"
	RatingQuestionable string = "q"
	RatingExplicit     string = "e"
)

const (
	orPrefix = "~"
	noPrefix = "-"

	lessPrefix           = "<"
	lessOrEqualPrefix    = "<="
	greaterPrefix        = ">"
	greaterOrEqualPrefix = ">="
)

const (
	ratingKey = "rating:"
	dateKey   = "date:"
)

const timeFormat = "2006-01-02"

type Tags struct {
	Keywords includeExclude            `yaml:"tags"`
	Meta     map[string]includeExclude `yaml:",inline"`
}

// NewTags allocates a new list. You may provide as many tags as you want.
//  t := getmoe.NewTags("tag1", "tag2")
func NewTags(tags ...string) *Tags {
	return new(Tags).And(tags...)
}

// And appends tags to tag list.
//  t.And("tag3", "tag4")
func (t *Tags) And(tags ...string) *Tags {
	t.Keywords.Include = append(t.Keywords.Include, tags...)
	return t
}

// Or appends tags with 'or' prefix to tag list.
//  t.Or("tag5", "tag6")
func (t *Tags) Or(tags ...string) *Tags {
	for i := range tags {
		t.Keywords.Include = append(t.Keywords.Include, orPrefix+tags[i])
	}
	return t
}

// No appends tags with 'no' prefix to tag list.
//  t.No("tag7", "tag8")
func (t *Tags) No(tags ...string) *Tags {
	t.Keywords.Exclude = append(t.Keywords.Exclude, tags...)
	return t
}

// WithRating appends rating tags.
func (t *Tags) WithRating(r ...string) *Tags {
	tmp := t.Meta[ratingKey]
	tmp.Include = append(tmp.Include, r...)
	t.Meta[ratingKey] = tmp
	return t
}

// WithoutRating appends rating tags with 'no' prefix.
func (t *Tags) WithoutRating(r ...string) *Tags {
	tmp := t.Meta[ratingKey]
	tmp.Exclude = append(tmp.Exclude, r...)
	t.Meta[ratingKey] = tmp
	return t
}

// AtDate appends date tag.
func (t *Tags) AtDate(date time.Time) *Tags {
	tmp := t.Meta[dateKey]
	tmp.Include = append(tmp.Include, date.Format(timeFormat))
	t.Meta[ratingKey] = tmp
	return t
}

// BeforeDate appends date tag with before prefix.
func (t *Tags) BeforeDate(date time.Time) *Tags {
	tmp := t.Meta[dateKey]
	tmp.Include = append(tmp.Include, lessOrEqualPrefix+date.Format(timeFormat))
	t.Meta[ratingKey] = tmp
	return t
}

// AfterDate appends date tag with after prefix.
func (t *Tags) AfterDate(date time.Time) *Tags {
	tmp := t.Meta[dateKey]
	tmp.Include = append(tmp.Include, greaterOrEqualPrefix+date.Format(timeFormat))
	t.Meta[ratingKey] = tmp
	return t
}

func (t *Tags) String() string {
	var tags []string
	tags = append(tags, t.Keywords.StringWithPrefix(""))
	for key, value := range t.Meta {
		tags = append(tags, value.StringWithPrefix(key))
	}
	return strings.Join(tags, " ")
}
