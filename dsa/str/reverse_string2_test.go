package str

import "testing"

func TestReverseString2(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ReverseString2_Test1",
			args: args{
				s: "hello",
				k: 2,
			},
			want: "ehllo",
		}, {
			name: "ReverseString2_Test2",
			args: args{
				s: "abcdefg",
				k: 2,
			},
			want: "bacdfeg",
		}, {
			name: "ReverseString2_Test3",
			args: args{
				s: "abcd",
				k: 2,
			},
			want: "bacd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString2(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("ReverseString2() = %v, want %v", got, tt.want)
			}
		})
	}
}
