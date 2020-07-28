package boop

import "testing"

func TestSayMeep(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "String Output Success",
			want: "MEEP",
		},
		// {
		// 	name: "String Ouput Failure",
		// 	want: "MOOP",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SayMeep(); got != tt.want {
				t.Errorf("SayMeep() = %v, want %v", got, tt.want)
			}
		})
	}
}
