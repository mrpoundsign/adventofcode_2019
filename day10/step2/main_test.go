package main

import "testing"

func Test_point_angleTo(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		p2 point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name:   "5,5 to 5,-16",
			fields: fields{X: 5, Y: 5},
			args:   args{point{X: 5, Y: -16}},
			want:   0,
		},
		{
			name:   "5,5 to 8,2",
			fields: fields{X: 5, Y: 5},
			args:   args{point{X: 8, Y: 2}},
			want:   45,
		},
		{
			name:   "5,5 to 16,5",
			fields: fields{X: 5, Y: 5},
			args:   args{point{X: 16, Y: 5}},
			want:   90,
		},
		{
			name:   "5,5 to 16,16",
			fields: fields{X: 5, Y: 5},
			args:   args{point{X: 16, Y: 16}},
			want:   135,
		},
		{
			name:   "5,5 to 5,16",
			fields: fields{X: 5, Y: 5},
			args:   args{point{X: 5, Y: 16}},
			want:   180,
		},
		{
			name:   "5,5 to -16,5",
			fields: fields{X: 5, Y: 5},
			args:   args{point{X: -16, Y: 5}},
			want:   270,
		},
		{
			name:   "5,5 to 0,0",
			fields: fields{X: 5, Y: 5},
			args:   args{point{X: 0, Y: 0}},
			want:   315,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.angleTo(tt.args.p2); got != tt.want {
				t.Errorf("point.angleTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
