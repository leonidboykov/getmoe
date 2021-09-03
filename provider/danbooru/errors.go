package danbooru

import "strings"

// ErrResponse wraps an error from Danbooru API response.
type ErrResponse struct {
	Success   bool     `json:"success"`
	Message   string   `json:"message"`
	Backtrace []string `json:"backtrace"`
}

func (e ErrResponse) Error() string {
	return "danbooru: " + strings.ToLower(e.Message)
}
