package boop_test

import (
	"testing"

	"github.com/aplombomb/jdp/pkg/boop"
)

func TestString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "String Output Success",
			want: "HELLO",
		},
		// {
		// 	name: "String Output Failure",
		// 	want: "BONJOUR",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := boop.SayHello(); got != tt.want {
				t.Errorf("Test() = %v, want %v", got, tt.want)
			}
		})
	}
}
