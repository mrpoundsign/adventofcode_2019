package intcode

import "fmt"

import "errors"

type ioReadWriter interface {
	ReadValue() (int, error)
	WriteValue(int) error // error if we should halt
	Fail()
	Exit()
}

type cmd int

const (
	cmdAdd cmd = iota + 1
	cmdMultiply
	cmdSet
	cmdGet
	cmdJumpIfTrue
	cmdJumpIfFalse
	cmdLessThan
	cmdEquals
	cmdSetRelativeBase
	cmdEnd = 99
)

type mode int

const (
	modePosition mode = iota
	modeImmiedate
	modeRelative
)

type runner struct {
	prog    []int
	extMem  map[int]int
	rw      ioReadWriter
	rbase   int
	pointer int
}

func (r *runner) getOffset(addr int, m mode) int {
	switch m {
	case modeImmiedate:
		return addr
	case modeRelative:
		return addr + r.rbase
	}

	return r.get(addr)
}

func (r *runner) set(addr, value int) {
	if addr < len(r.prog) {
		r.prog[addr] = value
		return
	}

	r.extMem[addr-len(r.prog)] = value
}

func (r *runner) get(addr int) int {
	if addr < len(r.prog) {
		return r.prog[addr]
	}

	value, ok := r.extMem[addr-len(r.prog)]
	if !ok {
		return 0
	}
	return value
}

func (r *runner) getCmdCode(addr int) (cmd, error) {
	if addr < len(r.prog) {
		return cmd(r.prog[addr] % 100), nil
	}

	return 0, errors.New("attempted to execute in extended memory")
}

func (r *runner) setRbase(addr int) {
	r.rbase = addr
}

func (r *runner) run() error {
	defer r.rw.Fail()

	progLen := len(r.prog)

	for r.pointer < progLen {
		code, err := r.getCmdCode(r.pointer)
		if err != nil {
			return fmt.Errorf("run error retreiving code, %w", err)
		}

		// Default to parameter mode
		param1 := r.pointer + 1
		param2 := r.pointer + 2
		param3 := r.pointer + 3

		pointerAdd := 1
		at := r.get(r.pointer)

		if code != cmdEnd {
			param1 = r.getOffset(r.pointer+1, mode((at/100)%10))
			pointerAdd++

			if code != cmdGet && code != cmdSet && code != cmdSetRelativeBase {
				param2 = r.getOffset(r.pointer+2, mode((at/1_000)%10))
				pointerAdd++

				if code != cmdJumpIfFalse && code != cmdJumpIfTrue {
					param3 = r.getOffset(r.pointer+3, mode((at/10_000)%10))
					pointerAdd++
				}
			}
		}

		switch code {
		case cmdAdd:
			r.set(param3, r.get(param1)+r.get(param2))
		case cmdMultiply:
			r.set(param3, r.get(param1)*r.get(param2))
		case cmdSet:
			ip, err := r.rw.ReadValue()
			if err != nil {
				return fmt.Errorf("read failure %w", err)
			}
			r.set(param1, ip)
		case cmdGet:
			err := r.rw.WriteValue(r.get(param1))
			if err != nil {
				return fmt.Errorf("error reading input %w", err)
			}
		case cmdJumpIfTrue:
			if r.get(param1) != 0 {
				r.pointer = r.get(param2)
				continue
			}
		case cmdJumpIfFalse:
			if r.get(param1) == 0 {
				r.pointer = r.get(param2)
				continue
			}
		case cmdLessThan:
			if r.get(param1) < r.get(param2) {
				r.set(param3, 1)
			} else {
				r.set(param3, 0)
			}
		case cmdEquals:
			if r.get(param1) == r.get(param2) {
				r.set(param3, 1)
			} else {
				r.set(param3, 0)
			}
		case cmdSetRelativeBase:
			r.setRbase(param1)
		case cmdEnd:
			r.rw.Exit()
			return nil
		default:
			return fmt.Errorf("invalid opcode %d at %d", code, r.pointer)
		}

		r.pointer += pointerAdd
	}

	return nil
}

func Run(program []int, input int) ([]int, int, error) {
	vh := &ValueHolder{value: input}
	prog, err := RunWithAmp(program, vh)
	return prog, vh.value, err
}

func RunWithAmp(program []int, rw ioReadWriter) ([]int, error) {
	r := runner{prog: program, rw: rw, extMem: make(map[int]int)}
	err := r.run()
	return r.prog, err
}
