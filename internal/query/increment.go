package query

import (
	"net/url"
	"strconv"
	"strings"
)

// Increment increments value based on key provided
func Increment(q *url.Values, key string) error {
	strValue := q.Get(key)
	// handle empty value
	if strValue == "" {
		q.Set(key, "1")
		return nil
	}

	value, err := strconv.Atoi(strValue)
	if err != nil {
		return err
	}

	value++

	q.Set(key, strconv.Itoa(value))

	return nil
}

// Int allows to set integer as query value
func Int(q *url.Values, key string, value int) {
	q.Set(key, strconv.Itoa(value))
}

// Array allows to set array of strings as query value
func Array(q *url.Values, key string, value []string) {
	q.Set(key, strings.Join(value, " "))
}
