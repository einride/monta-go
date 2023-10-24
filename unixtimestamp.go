package monta

import (
	"encoding/json"
	"time"
)

type UnixTimestamp struct {
	time.Time
}

// Unmarshal a unix timestamp (in milliseconds) from a JSON object to a human readable timestamp.
func (ut *UnixTimestamp) UnmarshalJSON(data []byte) error {
	var timestamp int64
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return err
	}
	ut.Time = time.Unix(timestamp/1000, 0)
	return nil
}
