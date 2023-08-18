package dsa

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "TestMinSubArrayLen_Solution1_Example1",
			args: args{
				nums:   []int{2, 3, 1, 2, 4, 3},
				target: 7,
			},
			want: 2,
		}, {
			name: "TestMinSubArrayLen_Solution1_Example2",
			args: args{
				nums:   []int{1, 4, 4},
				target: 4,
			},
			want: 1,
		}, {
			name: "TestMinSubArrayLen_Solution1_Example3",
			args: args{
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
				target: 11,
			},
			want: 0,
		}, {
			name: "TestMinSubArrayLen_Solution1_Example4",
			args: args{
				nums:   []int{5,1,3,5,10,7,4,9,2,8},
				target: 15,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinSubArrayLen(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("MinSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
