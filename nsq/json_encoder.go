package nsq

import (
	"encoding/json"
)

// JSONEncoder is an Encoder interface implementation for JSON.
type JSONEncoder struct{}

// Encode encodes data to JSON.
func (e *JSONEncoder) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// NewJSONEncoder creates new Encoder.
func NewJSONEncoder() Encoder { return &JSONEncoder{} }
