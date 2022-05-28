package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}

func (s StrTo) Int64() (int64, error) {
	v, err := strconv.ParseInt(s.String(), 10, 64)
	return v, err
}

func (s StrTo) MustInt64() int64 {
	v, _ := s.Int64()
	return v
}

func (s StrTo) Float64() (float64, error) {
	return strconv.ParseFloat(s.String(), 64)
}

func (s StrTo) MustFloat64() float64 {
	v, _ := strconv.ParseFloat(s.String(), 64)
	return v
}
