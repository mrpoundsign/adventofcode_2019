package intcode

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	progLessThan8 := []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	type args struct {
		program []int64
		value   int64
	}
	tests := []struct {
		name      string
		args      args
		want      []int64
		wantValue int64
		wantErr   bool
	}{
		{
			name: "d5s1t1 10 returns 10",
			args: args{
				program: []int64{3, 0, 4, 0, 99},
				value:   10,
			},
			want:      []int64{10, 0, 4, 0, 99},
			wantValue: 10,
			wantErr:   false,
		},
		{
			name: "d5s2t1 test 0",
			args: args{
				program: []int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
				value:   10,
			},
			want:      []int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 10, 1, 1, 9},
			wantValue: 1,
			wantErr:   false,
		},
		{
			name: "d5s2t1 test 1",
			args: args{
				program: []int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
				value:   0,
			},
			want:      []int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 0, 0, 1, 9},
			wantValue: 0,
			wantErr:   false,
		},
		{
			name: "d5s2t2 less than 8",
			args: args{
				program: progLessThan8,
				value:   0,
			},
			want:      progLessThan8,
			wantValue: 999,
			wantErr:   false,
		},
		{
			name: "d5s2t2 greater 8",
			args: args{
				program: progLessThan8,
				value:   9,
			},
			want:      progLessThan8,
			wantValue: 1001,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotValue, err := Run(tt.args.program, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
			if gotValue != tt.wantValue {
				t.Errorf("Run() rw.ReadValue() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func Test_runner_getOffset(t *testing.T) {
	io := &ValueHolder{value: 3}
	prog := []int64{1002, 11, 9, 4}

	type fields struct {
		prog    []int64
		extMem  map[int64]int64
		rw      ioReadWriter
		rbase   int
		pointer int
	}

	f := fields{
		prog:    prog,
		rbase:   1,
		rw:      io,
		pointer: 0,
	}

	type args struct {
		offset int
		m      mode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "position mode",
			fields: f,
			args:   args{offset: 1, m: modePosition},
			want:   11,
		},
		{
			name:   "immiedate mode",
			fields: f,
			args:   args{offset: 2, m: modeImmiedate},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &runner{
				prog:   tt.fields.prog,
				extMem: tt.fields.extMem,
				rw:     tt.fields.rw,
				rbase:  tt.fields.rbase,
				// pointer: tt.fields.pointer,
			}
			if got := r.getOffset(tt.args.offset, tt.args.m); got != tt.want {
				t.Errorf("runner.getOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}
