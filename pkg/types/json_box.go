// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

import (
	"database/sql"
	"database/sql/driver"
	stdjson "encoding/json"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/rocboss/paopao-ce/pkg/json"
)

var (
	_ stdjson.Marshaler   = (*JsonBox[any])(nil)
	_ stdjson.Unmarshaler = (*JsonBox[any])(nil)
	_ driver.Valuer       = (*JsonBox[any])(nil)
	_ sql.Scanner         = (*JsonBox[any])(nil)
	_ Boxes[any]          = (*JsonBox[any])(nil)
)

// JsonBox Json box for process database/sql json data
type JsonBox[T any] struct {
	data T
}

func (j *JsonBox[T]) Box(t T) {
	j.data = t
}

func (j *JsonBox[T]) Unbox() T {
	return j.data
}

func (j *JsonBox[T]) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte(`null`), nil
	}
	return json.Marshal(j.data)
}

func (j *JsonBox[T]) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("JsonBox.UnmarshalJSON: on nil pointer")
	}
	return json.Unmarshal(data, &j.data)

}

func (j *JsonBox[T]) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return j.MarshalJSON()
}

func (j *JsonBox[T]) Scan(value any) error {
	if value == nil {
		return nil
	}
	var b []byte
	switch v := value.(type) {
	case []byte:
		b = v
	case string:
		b = []byte(v)
	default:
		return fmt.Errorf("JsonBox.Scan: expected []byte or string, got %T (%q)", value, value)
	}
	return j.UnmarshalJSON(b)
}

// NewJsonBox create a new JsonBox instance
func NewJsonBox[T any](t ...T) *JsonBox[T] {
	if len(t) > 0 {
		return &JsonBox[T]{data: t[0]}
	}
	return &JsonBox[T]{}
}
