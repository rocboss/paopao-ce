package util

import "testing"

func TestEncodeMD5(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				value: "123456",
			},
			want: "e10adc3949ba59abbe56e057f20f883e",
		},
		{
			name: "test2",
			args: args{
				value: "",
			},
			want: "d41d8cd98f00b204e9800998ecf8427e", // really odd, why?
		},
		{
			name: "test3",
			args: args{
				value: "paopaocestr",
			},
			want: "8a5033dda1a8919224c66e68d846a289",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeMD5(tt.args.value); got != tt.want {
				t.Errorf("EncodeMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}
