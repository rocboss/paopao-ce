package convert

import (
	"testing"
)

func TestStrTo_String(t *testing.T) {
	type fields struct {
		Str string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "test",
			fields: fields{Str: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StrTo(tt.fields.Str)
			if got := s.String(); got != tt.want {
				t.Errorf("StrTo.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrTo_Int(t *testing.T) {
	var s StrTo = "123"
	i, err := s.Int()
	if err != nil {
		t.Errorf("s.Int() error: %v", err)
	}
	if i != 123 {
		t.Errorf("s.Int() should be 123")
	}
}

func TestStrTo_MustInt(t *testing.T) {
	var s StrTo = "123"
	i := s.MustInt()
	if i != 123 {
		t.Errorf("s.MustInt() should be 123")
	}
}

func TestStrTo_UInt32(t *testing.T) {
	var s StrTo = "123"
	i, err := s.UInt32()
	if err != nil {
		t.Errorf("s.UInt32() error: %v", err)
	}
	if i != 123 {
		t.Errorf("s.UInt32() should be 123")
	}
}

func TestStrTo_UInt32_2(t *testing.T) {
	var s StrTo = "-123"
	i, err := s.UInt32()
	if err != nil {
		t.Errorf("s.UInt32() error: %v", err)
	}
	if i == 123 {
		t.Errorf("s.UInt32() should not be 123")
	}
}

func TestStrTo_MustUInt32(t *testing.T) {
	var s StrTo = "123"
	i := s.MustUInt32()
	if i != 123 {
		t.Errorf("s.MustUInt32() should be 123")
	}
}

func TestStrTo_Int64(t *testing.T) {
	var s StrTo = "123"
	i, err := s.Int64()
	if err != nil {
		t.Errorf("s.Int64() error: %v", err)
	}
	if i != 123 {
		t.Errorf("s.Int64() should be 123")
	}
}

func TestStrTo_MustInt64(t *testing.T) {
	var s StrTo = "123"
	i := s.MustInt64()
	if i != 123 {
		t.Errorf("s.MustInt64() should be 123")
	}
}

func TestStrTo_Float64(t *testing.T) {
	var s StrTo = "123.456"
	f, err := s.Float64()
	if err != nil {
		t.Errorf("s.Float64() error: %v", err)
	}
	if f != 123.456 {
		t.Errorf("s.Float64() should be 123.456")
	}
}

func TestStrTo_MustFloat64(t *testing.T) {
	var s StrTo = "123.456"
	f := s.MustFloat64()
	if f != 123.456 {
		t.Errorf("s.MustFloat64() should be 123.456")
	}
}
