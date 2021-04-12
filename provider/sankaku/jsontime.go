package sankaku

import (
	"encoding/json"
	"time"
)

type jsonTime struct {
	time.Time
}

func (t *jsonTime) UnmarshalJSON(data []byte) error {
	var jsonTime struct {
		S int64 `json:"s,omitempty"`
		N int64 `json:"n,omitempty"`
	}

	if err := json.Unmarshal(data, &jsonTime); err != nil {
		return err
	}

	t.Time = time.Unix(jsonTime.S, jsonTime.N)
	return nil
}
