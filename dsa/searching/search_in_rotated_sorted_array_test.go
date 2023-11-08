package searching

import "testing"

func TestSearchInRotatedSortedArray(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "SearchInRotatedSortedArray_Test1",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		}, {
			name: "SearchInRotatedSortedArray_Test2",
			args: args{
				nums:   []int{5, 1, 3},
				target: 5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInRotatedSortedArray(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInRotatedSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
