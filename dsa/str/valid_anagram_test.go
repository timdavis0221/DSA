package str

import (
	"testing"
)

func TestValidAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ValidAnagram_Test1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		}, {
			name: "ValidAnagram_Test2",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("ValidAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidAnagramApproach2(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ValidAnagramApproach2_Test1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		}, {
			name: "ValidAnagramApproach2_Test2",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidAnagramApproach2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("ValidAnagramApproach2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidAnagramApproach3(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ValidAnagramApproach3_Test1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		}, {
			name: "ValidAnagramApproach3_Test2",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidAnagramApproach3(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("ValidAnagramApproach3() = %v, want %v", got, tt.want)
			}
		})
	}
}
