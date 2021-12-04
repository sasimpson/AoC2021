package main

import "testing"

func Test_getBit(t *testing.T) {
	type args struct {
		data int64
		pos  int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "data 1, pos 0",
			args: args{
				data: 1,
				pos:  0,
			},
			want: 1,
		}, {
			name: "data 1, pos 1",
			args: args{
				data: 1,
				pos:  1,
			},
			want: 0,
		}, {
			name: "data 2, pos 1",
			args: args{
				data: 2,
				pos:  1,
			},
			want: 1,
		}, {
			name: "data 2, pos 0",
			args: args{
				data: 2,
				pos:  0,
			},
			want: 0,
		}, {
			name: "data 2, pos 2",
			args: args{
				data: 2,
				pos:  0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBit(tt.args.data, tt.args.pos); got != tt.want {
				t.Errorf("getBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mostCommonBit(t *testing.T) {
	type args struct {
		data []int64
		pos  int
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "msb test 1",
			args: args{
				data: []int64{1, 1, 0, 0, 2, 3, 5},
				pos:  0,
			},
			want: 1,
		}, {
			name: "msb test 2",
			args: args{
				data: []int64{1, 0, 0, 0, 2, 3, 0},
				pos:  0,
			},
			want: 0,
		}, {
			name: "msb test 3",
			args: args{
				data: []int64{2, 0, 3, 0, 2, 3, 0},
				pos:  1,
			},
			want: 1,
		}, {
			name: "msb test 4",
			args: args{
				data: []int64{2, 0, 3, 0, 2, 0},
				pos:  1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCommonBit(tt.args.data, tt.args.pos); got != tt.want {
				t.Errorf("mostCommonBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leastCommonBit(t *testing.T) {
	type args struct {
		data []int64
		pos  int
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "lsb test 1",
			args: args{
				data: []int64{1, 1, 0, 0, 2, 3, 5},
				pos:  0,
			},
			want: 0,
		}, {
			name: "lsb test 2",
			args: args{
				data: []int64{1, 0, 0, 0, 2, 3, 0},
				pos:  0,
			},
			want: 1,
		}, {
			name: "lsb test 3",
			args: args{
				data: []int64{2, 0, 3, 0, 2, 3, 0},
				pos:  1,
			},
			want: 0,
		}, {
			name: "lsb test 4",
			args: args{
				data: []int64{2, 0, 3, 0, 2, 0},
				pos:  1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastCommonBit(tt.args.data, tt.args.pos); got != tt.want {
				t.Errorf("leastCommonBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
