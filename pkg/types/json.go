// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

import (
	"database/sql"
	"database/sql/driver"
	stdjson "encoding/json"
	"errors"
	"fmt"

	"github.com/rocboss/paopao-ce/pkg/json"
)

var (
	_ stdjson.Marshaler   = (*JsonType[any])(nil)
	_ stdjson.Unmarshaler = (*JsonType[any])(nil)
	_ driver.Valuer       = (*JsonType[any])(nil)
	_ sql.Scanner         = (*JsonType[any])(nil)
)

type JsonType[T any] struct {
	Data T
}

func (j *JsonType[T]) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte(`null`), nil
	}
	return json.Marshal(j.Data)
}

func (j *JsonType[T]) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("JSONText.UnmarshalJSON: on nil pointer")
	}
	return json.Unmarshal(data, &j.Data)

}

func (j *JsonType[T]) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return j.MarshalJSON()
}

func (j *JsonType[T]) Scan(value any) error {
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
		return fmt.Errorf("JSONText.Scan: expected []byte or string, got %T (%q)", value, value)
	}
	return j.UnmarshalJSON(b)
}
