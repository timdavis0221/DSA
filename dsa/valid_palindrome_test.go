package dsa

import "testing"

func TestValidPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid_Palindrome_Test1",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		}, {
			name: "Valid_Palindrome_Test2",
			args: args{
				s: "race a car",
			},
			want: false,
		}, {
			name: "Valid_Palindrome_Test3",
			args: args{
				s: " ",
			},
			want: true,
		}, {
			name: "Valid_Palindrome_Test4",
			args: args{
				s: "!043XjqjX043!",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidPalindrome(tt.args.s); got != tt.want {
				t.Errorf("ValidPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
