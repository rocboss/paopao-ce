// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

// Empty empty alias type
type Empty = struct{}

// Fn empty argument func alias type
type Fn = func()

// Boxes Box/Unbox interface
type Boxes[T any] interface {
	Box(t T)
	Unbox() T
}
