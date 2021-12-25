package api

import "testing"

func TestEquals(t *testing.T) {
	type args struct {
		v1 []Card
		v2 []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				v1: []Card{{
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}},
				v2: []Card{{
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				v1: []Card{{
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}},
				v2: []Card{{
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Hog Rider",
				}},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				v1: []Card{{
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}},
				v2: []Card{{
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 1,
					Name:  "Valkirie",
				}, {
					Level: 2,
					Name:  "Valkirie",
				}},
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
