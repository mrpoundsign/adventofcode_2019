package main

import (
	"reflect"
	"testing"
)

func Test_point_setDir(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		i int
	}

	up := fields{Y: 1}
	down := fields{Y: -1}
	right := fields{X: 1}
	left := fields{X: -1}
	turnL := args{0}
	turnR := args{1}

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantPoint fields
	}{
		// left turns
		{
			name:      "up turn left",
			fields:    up,
			args:      turnL,
			wantPoint: left,
		},
		{
			name:      "down turn left",
			fields:    down,
			args:      turnL,
			wantPoint: right,
		},
		{
			name:      "left turn left",
			fields:    left,
			args:      turnL,
			wantPoint: down,
		},
		{
			name:      "right turn left",
			fields:    right,
			args:      turnL,
			wantPoint: up,
		},

		// right turns
		{
			name:      "up turn right",
			fields:    up,
			args:      turnR,
			wantPoint: right,
		},
		{
			name:      "down turn right",
			fields:    down,
			args:      turnR,
			wantPoint: left,
		},
		{
			name:      "left turn right",
			fields:    left,
			args:      turnR,
			wantPoint: up,
		},
		{
			name:      "right turn right",
			fields:    right,
			args:      turnR,
			wantPoint: down,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			p.setDir(tt.args.i)

			if !reflect.DeepEqual(p.X, tt.wantPoint.X) {
				t.Errorf("Run() p.X = %v, want %v", p.X, tt.wantPoint.X)
			}
			if !reflect.DeepEqual(p.Y, tt.wantPoint.Y) {
				t.Errorf("Run() p.Y = %v, want %v", p.Y, tt.wantPoint.Y)
			}

		})
	}
}

func Test_paint(t *testing.T) {
	type args struct {
		path []int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simples",
			args: args{path: []int64{0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paint(tt.args.path); got != tt.want {
				t.Errorf("paint() = %v, want %v", got, tt.want)
			}
		})
	}
}
