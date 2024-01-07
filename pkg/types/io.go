// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

import (
	"io"
)

type readerWrap struct {
	r io.Reader
}

func (rw *readerWrap) Read(p []byte) (n int, err error) {
	return rw.r.Read(p)
}

// PureReader wrap a pure io.Reader object
func PureReader(r io.Reader) io.Reader {
	return &readerWrap{
		r: r,
	}
}
