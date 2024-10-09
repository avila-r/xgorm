package xgorm

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Array []interface{}

func (j Array) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Array) Scan(v interface{}) error {
	b, ok := v.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &j)
}

type ArrayOf[T any] []T

func (j ArrayOf[T]) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *ArrayOf[T]) Scan(v interface{}) error {
	b, ok := v.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &j)
}
