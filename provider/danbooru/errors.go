package danbooru

import "strings"

type ErrResponse struct {
	Success   bool     `json:"success"`
	Message   string   `json:"message"`
	Backtrace []string `json:"backtrace"`
}

func (e *ErrResponse) Error() string {
	return "gelbooru: " + strings.ToLower(e.Message)
}
