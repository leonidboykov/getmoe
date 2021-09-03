package sankaku

type ErrResponse struct {
	Success bool   `json:"success"`
	Message string `json:"error"`
}

func (e ErrResponse) Error() string {
	return e.Message
}
