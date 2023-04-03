// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// PgxArray returns an object usable by pg drivers for passing a []T slice
// into a database as type T[].
func PgxArray[T any](elements []T) pgtype.Array[T] {
	return pgtype.Array[T]{
		Elements: elements,
		Dims:     []pgtype.ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
	}
}
