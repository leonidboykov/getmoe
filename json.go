package getmoe

import (
	"encoding/json"
	"time"
)

// JSONTime provides unmarshaller for json timestamp.
//
// Example:
//  "created_at": {
//      "json_class": "Time",
//      "s": 1552231391,
//      "n": 0
//  }
type JSONTime struct {
	time.Time
}

// UnmarshalJSON marshals JSONTime to time.Time.
func (t *JSONTime) UnmarshalJSON(data []byte) error {
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
