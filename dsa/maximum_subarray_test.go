package dsa

import (
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Maximum_Subarray_Test1",
			args: args{
				nums: []int{-2, 2, 5, -11, 6},
			},
			want: 7,
		}, {
			name: "Maximum_Subarray_Test2",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		}, {
			name: "Maximum_Subarray_Test3",
			args: args{
				nums: []int{5, 4, -1, 7, 8},
			},
			want: 23,
		}, {
			name: "Maximum_Subarray_Test4",
			args: args{
				nums: []int{8, -19, 5, -4, 20},
			},
			want: 21,
		}, {
			name: "Maximum_Subarray_Test5",
			args: args{
				nums: []int{3, -2, -3, -3, 1, 3, 0},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumSubarray(tt.args.nums); got != tt.want {
				t.Errorf("MaximumSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaximumSubarrayApproach2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Maximum_Subarray_Approach2_Test1",
			args: args{
				nums: []int{-2, 2, 5, -11, 6},
			},
			want: 7,
		}, {
			name: "Maximum_Subarray_Approach2_Test2",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		}, {
			name: "Maximum_Subarray_Approach2_Test3",
			args: args{
				nums: []int{5, 4, -1, 7, 8},
			},
			want: 23,
		}, {
			name: "Maximum_Subarray_Approach2_Test4",
			args: args{
				nums: []int{8, -19, 5, -4, 20},
			},
			want: 21,
		}, {
			name: "Maximum_Subarray_Approach2_Test5",
			args: args{
				nums: []int{3, -2, -3, -3, 1, 3, 0},
			},
			want: 4,
		}, {
			name: "Maximum_Subarray_Approach2_Test6",
			args: args{
				nums: []int{-2, -1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumSubarrayApproach2(tt.args.nums); got != tt.want {
				t.Errorf("MaximumSubarrayApproach2() = %v, want %v", got, tt.want)
			}
		})
	}
}
