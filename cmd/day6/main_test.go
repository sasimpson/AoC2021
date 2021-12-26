package main

import (
	"reflect"
	"testing"
)

func Test_school_incrementDay(t *testing.T) {
	type fields struct {
		day         int
		fish        []fish
		generations [][]generation
	}
	tests := []struct {
		name   string
		fields fields
		want   []generation
	}{
		{
			name: "normal iteration",
			fields: fields{
				day: 0,
				generations: [][]generation{
					{{0, 0}, {1, 1}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}},
				},
			},
			want: []generation{{0, 1}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}},
		}, {
			name: "spawn iteration",
			fields: fields{
				day: 0,
				generations: [][]generation{
					{{0, 1}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}},
				},
			},
			want: []generation{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 1}, {7, 0}, {8, 1}},
		}, {
			name: "spawn iteration 2",
			fields: fields{
				day: 0,
				generations: [][]generation{
					{{0, 1}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 1}, {8, 0}},
				},
			},
			want: []generation{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 2}, {7, 0}, {8, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &school{
				day:         tt.fields.day,
				fish:        tt.fields.fish,
				generations: tt.fields.generations,
			}
			if got := s.incrementDay(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("incrementDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
