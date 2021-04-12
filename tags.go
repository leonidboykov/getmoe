package getmoe

import (
	"strings"
	"time"
)

// Rating defines boorus rating system
type Rating string

// Boorus rating system has safe, questionable and explicit tags.
const (
	RatingSafe         Rating = "s"
	RatingQuestionable Rating = "q"
	RatingExplicit     Rating = "e"
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

// Tags provides fluent-style builder for boorus tags.
type Tags []string

// NewTags allocates a new list. You may provide as many tags as you want.
//  t := getmoe.NewTags("tag1", "tag2")
func NewTags(tags ...string) *Tags {
	return new(Tags).And(tags...)
}

// And appends tags to tag list.
//  t.And("tag3", "tag4")
func (t *Tags) And(tags ...string) *Tags {
	*t = append(*t, tags...)
	return t
}

// Or appends tags with 'or' prefix to tag list.
// t.Or("tag5", "tag6")
func (t *Tags) Or(tags ...string) *Tags {
	for i := range tags {
		*t = append(*t, orPrefix+tags[i])
	}
	return t
}

// No appends tags with 'no' prefix to tag list.
func (t *Tags) No(tags ...string) *Tags {
	for i := range tags {
		*t = append(*t, noPrefix+tags[i])
	}
	return t
}

// WithRating appends rating tags.
func (t *Tags) WithRating(r ...Rating) *Tags {
	for i := range r {
		*t = append(*t, ratingKey+string(r[i]))
	}
	return t
}

// WithoutRating appends rating tags with 'no' prefix.
func (t *Tags) WithoutRating(r ...Rating) *Tags {
	for i := range r {
		*t = append(*t, noPrefix+ratingKey+string(r[i]))
	}
	return t
}

// AtDate appends date tag.
func (t *Tags) AtDate(date time.Time) *Tags {
	*t = append(*t, dateKey+date.Format(timeFormat))
	return t
}

// BeforeDate appends date tag with before prefix.
func (t *Tags) BeforeDate(date time.Time) *Tags {
	*t = append(*t, dateKey+lessOrEqualPrefix+date.Format(timeFormat))
	return t
}

// AfterDate appends date tag with after prefix.
func (t *Tags) AfterDate(date time.Time) *Tags {
	*t = append(*t, dateKey+greaterOrEqualPrefix+date.Format(timeFormat))
	return t
}

func (t Tags) String() string {
	return strings.Join(t, " ")
}
