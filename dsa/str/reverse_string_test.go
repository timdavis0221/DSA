package str

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ReverseString_Test1",
			args: args{
				s: inputStringToBytes([]string{"h", "e", "l", "l", "o"}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseString(tt.args.s)
		})
	}
}

func inputStringToBytes(s []string) []byte {
	var byteSlice []byte
	for _, e := range s {
		b := []byte(e)
		byteSlice = append(byteSlice, b...)
	}
	return byteSlice
}

func TestReverseStringApproach2(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ReverseString_Approach2_Test1",
			args: args{
				s: inputStringToBytes([]string{"y", "e", "s"}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseStringApproach2(tt.args.s)
		})
	}
}

func TestReverseStringApproach3(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s []byte
	}{
		{
			name: "ReverseString_Approach3_Test1",
			s:    inputStringToBytes([]string{"h", "e", "l", "l", "o"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseStringApproach3(tt.s)
		})
	}
}
