package array

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Product_Except_Self_Test1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		}, {
			name: "Product_Except_Self_Test2",
			args: args{
				nums: []int{-1, 1, 0, -3, 3},
			},
			want: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductExceptSelfApproach2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Product_Except_Self_Test3",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		}, {
			name: "Product_Except_Self_Test4",
			args: args{
				nums: []int{-1, 1, 0, -3, 3},
			},
			want: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductExceptSelfApproach2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductExceptSelfApproach2() = %v, want %v", got, tt.want)
			}
		})
	}
}
