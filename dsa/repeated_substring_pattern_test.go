package dsa

import "testing"

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
