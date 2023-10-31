// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/RoaringBitmap/roaring"
	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/alimy/tryst/types"
)

type roaringBitmap struct {
	Map *roaring.Bitmap
}

type roaring64Bitmap struct {
	Map *roaring64.Bitmap
}

func (m *roaringBitmap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return []byte{}, nil
	}
	return m.Map.MarshalBinary()
}

func (m *roaringBitmap) UnmarshalBinary(data []byte) (res *roaringBitmap, err error) {
	res = &roaringBitmap{
		Map: roaring.New(),
	}
	err = res.Map.UnmarshalBinary(data)
	return
}

func (m *roaring64Bitmap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return []byte{}, nil
	}
	return m.Map.MarshalBinary()
}

func (m *roaring64Bitmap) UnmarshalBinary(data []byte) (res *roaring64Bitmap, err error) {
	res = &roaring64Bitmap{
		Map: roaring64.New(),
	}
	err = res.Map.UnmarshalBinary(data)
	return
}

// Bitmap alias type of types.Binary[*roaringBitmap]
type Bitmap = types.Binary[*roaringBitmap]

// NullBitmap alias type of types.NullBinary[*roaringBitmap]
type NullBitmap = types.NullBinary[*roaringBitmap]

// Bitmap64 alias type of types.Binary[*roaring64Bitmap]
type Bitmap64 = types.Binary[*roaring64Bitmap]

// NullBitmap64 alias type of types.NullBinary[*roaring64Bitmap]
type NullBitmap64 = types.NullBinary[*roaring64Bitmap]

// NewBitmap create a Bitmap instance
func NewBitmap() (res *Bitmap) {
	return &types.Binary[*roaringBitmap]{
		Data: &roaringBitmap{
			Map: roaring.New(),
		},
	}
}

// NewNullBitmap create a NullBitmap instance
func NewNullBitmap() *NullBitmap {
	return &types.NullBinary[*roaringBitmap]{
		Data: &roaringBitmap{
			Map: roaring.New(),
		},
	}
}

// NewBitmap64 create a Bitmap64 instance
func NewBitmap64() *Bitmap64 {
	return &types.Binary[*roaring64Bitmap]{
		Data: &roaring64Bitmap{
			Map: roaring64.New(),
		},
	}
}

// NewNullBitmap64 create a NullBitmap64 instance
func NewNullBitmap64() *NullBitmap64 {
	return &types.NullBinary[*roaring64Bitmap]{
		Data: &roaring64Bitmap{
			Map: roaring64.New(),
		},
	}
}

// MustBitmap create a Bitmap instance
func MustBitmap(data ...[]byte) (res *Bitmap) {
	res = &types.Binary[*roaringBitmap]{
		Data: &roaringBitmap{
			Map: roaring.New(),
		},
	}
	if len(data) > 0 && len(data[0]) > 0 {
		if err := res.Data.Map.UnmarshalBinary(data[0]); err != nil {
			panic(err)
		}
	}
	return
}

// MustNullBitmap create a NullBitmap instance
func MustNullBitmap(data ...[]byte) (res *NullBitmap) {
	res = &types.NullBinary[*roaringBitmap]{
		Data: &roaringBitmap{
			Map: roaring.New(),
		},
	}
	if len(data) > 0 && len(data[0]) > 0 {
		if err := res.Data.Map.UnmarshalBinary(data[0]); err != nil {
			panic(err)
		}
	}
	return
}

// MustBitmap64 create a Bitmap64 instance
func MustBitmap64(data ...[]byte) (res *Bitmap64) {
	res = &types.Binary[*roaring64Bitmap]{
		Data: &roaring64Bitmap{
			Map: roaring64.New(),
		},
	}
	if len(data) > 0 && len(data[0]) > 0 {
		if err := res.Data.Map.UnmarshalBinary(data[0]); err != nil {
			panic(err)
		}
	}
	return
}

// MustNullBitmap64 create a NullBitmap64 instance
func MustNullBitmap64(data ...[]byte) (res *NullBitmap64) {
	res = &types.NullBinary[*roaring64Bitmap]{
		Data: &roaring64Bitmap{
			Map: roaring64.New(),
		},
	}
	if len(data) > 0 && len(data[0]) > 0 {
		if err := res.Data.Map.UnmarshalBinary(data[0]); err != nil {
			panic(err)
		}
	}
	return
}
