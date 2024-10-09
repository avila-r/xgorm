package xgorm

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Json map[string]interface{}

func (j Json) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Json) Scan(v interface{}) error {
	b, ok := v.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &j)
}
