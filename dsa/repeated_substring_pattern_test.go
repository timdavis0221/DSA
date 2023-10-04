package dsa

import (
	"testing"
)

func TestRepeatedSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "RepeatedSubstringPattern_Test1",
			args: args{
				s: "abab",
			},
			want: true,
		}, {
			name: "RepeatedSubstringPattern_Test2",
			args: args{
				s: "aba",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepeatedSubstringPattern(tt.args.s); got != tt.want {
				t.Errorf("RepeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepeatedSubstringPatternApproach2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "RepeatedSubstringPatternApproach2_Test1",
			args: args{
				s: "aabbccaabbcc",
			},
			want: true,
		}, {
			name: "RepeatedSubstringPatternApproach2_Test2",
			args: args{
				s: "abaa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepeatedSubstringPatternApproach2(tt.args.s); got != tt.want {
				t.Errorf("RepeatedSubstringPatternApproach2() = %v, want %v", got, tt.want)
			}
		})
	}
}
