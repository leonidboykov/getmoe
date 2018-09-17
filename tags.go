package getmoe

import (
	"strings"
	"time"
)

// Tags interface provides a generic interface for the
// type Tags interface {
// 	With(t string) *Tag
// 	AfterDate(t time.Time) *Tag
// 	BeforeDate(t time.Time) *Tag
// }

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

// Rating defines boorus rating system
type Rating string

// Boorus rating system has safe, questionable and explicit tags
const (
	RatingSafe         Rating = "s"
	RatingQuestionable        = "q"
	RatingExplicit            = "e"
)

// Tags ...
type Tags []string

// NewTags ...
func NewTags(tags ...string) *Tags {
	return new(Tags).And(tags...)
}

// And ...
func (t *Tags) And(tags ...string) *Tags {
	*t = append(*t, tags...)
	return t
}

// Or ...
func (t *Tags) Or(tags ...string) *Tags {
	for i := range tags {
		*t = append(*t, orPrefix+tags[i])
	}
	return t
}

// No ...
func (t *Tags) No(tags ...string) *Tags {
	for i := range tags {
		*t = append(*t, noPrefix+tags[i])
	}
	return t
}

// HasRating appends rating tag
func (t *Tags) HasRating(r Rating) *Tags {
	*t = append(*t, ratingKey+string(r))
	return t
}

// HasNotRating appends rating tag with a minus prefix
func (t *Tags) HasNotRating(r Rating) *Tags {
	*t = append(*t, noPrefix+ratingKey+string(r))
	return t
}

// AtDate ...
func (t *Tags) AtDate(date time.Time) *Tags {
	*t = append(*t, dateKey+date.Format(timeFormat))
	return t
}

// BeforeDate ...
func (t *Tags) BeforeDate(date time.Time) *Tags {
	*t = append(*t, dateKey+lessOrEqualPrefix+date.Format(timeFormat))
	return t
}

// AfterDate ...
func (t *Tags) AfterDate(date time.Time) *Tags {
	*t = append(*t, dateKey+greaterOrEqualPrefix+date.Format(timeFormat))
	return t
}

func (t Tags) String() string {
	return strings.Join(t, " ")
}
