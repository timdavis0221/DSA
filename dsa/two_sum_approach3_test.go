package dsa

import (
	"reflect"
	"testing"
)

func TestTwoSumApproach3(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "TestTwoSumApproach3",
			args: args{
				nums:   []int{2, 3, 5, 7},
				target: 9,
			},
			want: []int{0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSumApproach3(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumApproach3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTwoSumApproach3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSumApproach3([]int{2, 3, 5, 7}, 9)
	}
}
