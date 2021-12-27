package main

import "testing"

func Test_crab_fuelToPos(t *testing.T) {
	type fields struct {
		position int
	}
	type args struct {
		pos int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "test 1",
			fields: fields{
				position: 1,
			},
			args: args{
				pos: 5,
			},
			want: 10,
		}, {
			name: "test 2",
			fields: fields{
				position: 2,
			},
			args: args{
				pos: 5,
			},
			want: 6,
		}, {
			name: "test 3",
			fields: fields{
				position: 16,
			},
			args: args{
				pos: 5,
			},
			want: 66,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crab{
				position: tt.fields.position,
			}
			if got := c.fuelToPos(tt.args.pos); got != tt.want {
				t.Errorf("fuelToPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
