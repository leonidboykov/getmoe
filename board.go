package getmoe

import (
	"net/url"
)

// BoardController ...
type BoardController interface {
	Request() []Post
	BuildAuth(...string) url.Values
}
