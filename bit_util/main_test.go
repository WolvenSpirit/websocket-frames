/**
*	For education purpose only, by Mihai Dragusin
* 	Do NOT copy
 */

package bit_util

import (
	"testing"
)

func Test_getPositiveMaskForBitRange(t *testing.T) {
	type args struct {
		bitLeft  int
		bitRight int
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "00000111",
			args: args{
				bitRight: 2,
				bitLeft:  3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPositiveMaskForBitRange(tt.args.bitLeft, tt.args.bitRight); got != tt.want {
				t.Errorf("getPositiveMaskForBitRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPositiveMaskForBitPosition(t *testing.T) {
	type args struct {
		bit int
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "2",
			args: args{bit: 2},
			want: 2,
		},
		{
			name: "3",
			args: args{bit: 3},
			want: 4,
		},
		{
			name: "3",
			args: args{bit: 4},
			want: 8,
		},
		{
			name: "127",
			args: args{bit: 8},
			want: 127,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPositiveMaskForBitPosition(tt.args.bit); got != tt.want {
				t.Errorf("getPositiveMaskForBitPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
