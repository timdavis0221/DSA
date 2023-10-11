package str

import "testing"

func TestStrStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "StrStr_Test1",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		}, {
			name: "StrStr_Test2",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		}, {
			name: "StrStr_Test3",
			args: args{
				haystack: "a",
				needle:   "a",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("StrStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
