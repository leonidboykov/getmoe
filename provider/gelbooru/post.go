package gelbooru

import (
	"strconv"
	"strings"
	"time"
)

// post contains native Gelbooru data.
type post struct {
	ID           int      `json:"id"`
	Width        int      `json:"width"`
	Height       int      `json:"height"`
	FileURL      string   `json:"file_url"`
	PreviewURL   string   `json:"preview_url"`
	SampleURL    string   `json:"sample_url"`
	Directory    int      `json:"directory"`
	Hash         string   `json:"hash"`
	Image        string   `json:"image"`
	Change       unixtime `json:"change"`
	Owner        string   `json:"owner"`
	ParentID     int      `json:"parent_id"`
	Rating       string   `json:"rating"`
	Sample       int      `json:"sample"`
	SampleHeight int      `json:"sample_height"`
	SampleWidth  int      `json:"sample_width"`
	Score        int      `json:"score"`
	Tags         tags     `json:"tags"`
}

type tags []string

// UnmarshalJSON implements JSON Unmarshaler.
func (t *tags) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\\", "") // hack for escaped slash in tag.
	str, err := strconv.Unquote(str)
	if err != nil {
		return err
	}
	*t = strings.Fields(str)
	return nil
}

type unixtime time.Time

// UnmarshalJSON implements JSON Unmarshaler.
func (u *unixtime) UnmarshalJSON(data []byte) error {
	q, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(u) = time.Unix(q, 0)
	return nil
}
