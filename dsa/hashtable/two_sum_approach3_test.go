package hashtable

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
		{
			name: "Two_Sum_Approach3_Test1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		}, {
			name: "Two_Sum_Approach3_Test2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
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
