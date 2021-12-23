package api

import "testing"

func TestEquals(t *testing.T) {
	type args struct {
		v1 []string
		v2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok",
			args: args{
				v1: []string{"1", "2", "3", "4"},
				v2: []string{"1", "2", "3", "4"},
			},
			want: true,
		},
		{
			name: "ko",
			args: args{
				v1: []string{"1", "2", "3", "4"},
				v2: []string{"1", "2", "3", "5"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equals(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
