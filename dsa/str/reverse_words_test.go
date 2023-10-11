package str

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ReverseWords_Test1",
			args: args{
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWords(tt.args.s); got != tt.want {
				t.Errorf("ReverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseWords2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ReverseWords2_Test1",
			args: args{
				s: "     the y blue is sky",
			},
			want: "sky is blue y the",
		}, {
			name: "ReverseWords2_Test2",
			args: args{
				s: "  hello world  ",
			},
			want: "world hello",
		}, {
			name: "ReverseWords2_Test3",
			args: args{
				s: "a good   example",
			},
			want: "example good a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWords2(tt.args.s); got != tt.want {
				t.Errorf("ReverseWords2() = %v, want %v", got, tt.want)
			}
		})
	}
}
