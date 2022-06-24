package util

import "testing"

func TestGetIPLoc(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				ip: "",
			},
			want: "",
		},
		{
			name: "test2",
			args: args{
				ip: "103.197.70.244",
			},
			want: "香港",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIPLoc(tt.args.ip); got != tt.want {
				t.Errorf("GetIPLoc() = %v, want %v", got, tt.want)
			}
		})
	}
}
