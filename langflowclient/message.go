package langflowclient

import (
	"encoding/json"
	"time"
)

func (m *Message) UnmarshalJSON(data []byte) error {
	type Alias Message
	aux := &struct {
		Timestamp string `json:"timestamp"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02 15:04:05", aux.Timestamp)
	if err != nil {
		return err
	}

	m.Timestamp = t
	return nil
}
