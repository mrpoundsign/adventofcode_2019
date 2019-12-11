package intcode

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	progLessThan8 := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	type args struct {
		program []int
		input   int
		expect  int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		want1   int
		wantErr bool
	}{
		{
			name: "d5s1t1 10 returns 10",
			args: args{
				program: []int{3, 0, 4, 0, 99},
				input:   10,
				expect:  10,
			},
			want:    []int{10, 0, 4, 0, 99},
			want1:   10,
			wantErr: false,
		},
		{
			name: "d5s2t1 test 0",
			args: args{
				program: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
				input:   10,
				expect:  1,
			},
			want:    []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 10, 1, 1, 9},
			want1:   1,
			wantErr: false,
		},
		{
			name: "d5s2t1 test 1",
			args: args{
				program: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
				input:   0,
				expect:  0,
			},
			want:    []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 0, 0, 1, 9},
			want1:   0,
			wantErr: false,
		},
		{
			name: "d5s2t2 less than 8",
			args: args{
				program: progLessThan8,
				input:   0,
				expect:  999,
			},
			want:    progLessThan8,
			want1:   999,
			wantErr: false,
		},
		{
			name: "d5s2t2 greater 8",
			args: args{
				program: progLessThan8,
				input:   9,
				expect:  1001,
			},
			want:    progLessThan8,
			want1:   1001,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Run(tt.args.program, tt.args.input, tt.args.expect)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Run() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
