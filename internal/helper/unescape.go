package helper

import (
	"net/url"
	"path"
	"strings"
)

// FileURLUnescape extracts file name from URL,
// unescapes URL and replaces spaces with dashes
func FileURLUnescape(s string) (string, error) {
	fileName := path.Base(s)

	fileName, err := url.QueryUnescape(fileName)
	if err != nil {
		return "", nil
	}

	// Replace spaces with dashes
	fileName = strings.Replace(fileName, " ", "-", -1)

	return fileName, nil
}
