// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

import "errors"

// internal core error variable for data logic implement.
var (
	ErrNotImplemented = errors.New("not implemented")
	ErrNoPermission   = errors.New("no permission")
)
