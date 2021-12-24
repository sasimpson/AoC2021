package main

import (
	"reflect"
	"testing"
)

func Test_line_Iterator(t *testing.T) {
	type fields struct {
		start point
		end   point
	}
	tests := []struct {
		name   string
		fields fields
		wantX  []int
		wantY  []int
	}{
		{
			name: "normal horizontal",
			fields: fields{
				start: point{x: 0, y: 0},
				end:   point{x: 5, y: 0},
			},
			wantX: []int{0, 1, 2, 3, 4, 5},
		}, {
			name: "normal vertical",
			fields: fields{
				start: point{x: 0, y: 0},
				end:   point{x: 0, y: 5},
			},
			wantY: []int{0, 1, 2, 3, 4, 5},
		}, {
			name: "normal diagonal",
			fields: fields{
				start: point{x: 0, y: 0},
				end:   point{x: 5, y: 5},
			},
			wantX: []int{0, 1, 2, 3, 4, 5},
			wantY: []int{0, 1, 2, 3, 4, 5},
		}, {
			name: "reverse horizontal",
			fields: fields{
				start: point{x: 5, y: 0},
				end:   point{x: 0, y: 0},
			},
			wantX: []int{5, 4, 3, 2, 1, 0},
		}, {
			name: "reverse vertical",
			fields: fields{
				start: point{x: 0, y: 5},
				end:   point{x: 0, y: 0},
			},
			wantY: []int{5, 4, 3, 2, 1, 0},
		}, {
			name: "reverse diagonal",
			fields: fields{
				start: point{x: 5, y: 5},
				end:   point{x: 0, y: 0},
			},
			wantX: []int{5, 4, 3, 2, 1, 0},
			wantY: []int{5, 4, 3, 2, 1, 0},
		}, {
			name: "odd diagonal",
			fields: fields{
				start: point{x: 0, y: 0},
				end:   point{x: 5, y: 3},
			},
			wantX: []int{0, 1, 2, 3, 4, 5},
			wantY: []int{0, 1, 2, 3},
		}, {
			name: "odd diagonal",
			fields: fields{
				start: point{x: 8, y: 0},
				end:   point{x: 0, y: 8},
			},
			wantX: []int{8, 7, 6, 5, 4, 3, 2, 1, 0},
			wantY: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := line{
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			gotX, gotY := l.Iterator()
			if !reflect.DeepEqual(gotX, tt.wantX) {
				t.Errorf("Iterator() gotX = %v, want %v", gotX, tt.wantX)
			}
			if !reflect.DeepEqual(gotY, tt.wantY) {
				t.Errorf("Iterator() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func Test_line_IsDiagonal(t *testing.T) {
	type fields struct {
		start point
		end   point
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "is diagonal",
			fields: fields{
				point{0, 0},
				point{2, 2},
			},
			want: true,
		}, {
			name: "is horizontal",
			fields: fields{
				point{0, 0},
				point{2, 0},
			},
			want: false,
		}, {
			name: "is vertical",
			fields: fields{
				point{0, 0},
				point{0, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := line{
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			if got := l.IsDiagonal(); got != tt.want {
				t.Errorf("IsDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}
