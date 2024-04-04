package binaryops

import "testing"

func TestBinaryGap(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestBinaryGap_Solution1",
			args: args{
				n: 22,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryGap(tt.args.n); got != tt.want {
				t.Errorf("BinaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
