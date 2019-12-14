package intcode

import "fmt"

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
	cmdEnd = 99
)

type runner struct {
	prog   []int
	extMem map[int]int
	rw     ioReadWriter
}

func (r *runner) set(offset, value int) {
	if offset < len(r.prog) {
		r.prog[offset] = value
		return
	}

	r.extMem[offset] = value
}

func (r *runner) get(offset int) int {
	if offset < len(r.prog) {
		return r.prog[offset]
	}

	value, ok := r.extMem[offset]
	if !ok {
		return 0
	}
	return value
}

func (r *runner) run() error {
	defer r.rw.Fail()
	progLen := len(r.prog)

	for i := 0; i < progLen; {
		code := cmd(r.get(i) % 100)

		// Default to parameter mode
		param1 := i + 1
		param2 := i + 2
		param3 := i + 3

		if code != cmdEnd {
			if (r.get(i)/100)%10 == 0 {
				param1 = r.get(param1)
			}

			if code != cmdGet && code != cmdSet {
				if (r.get(i)/1_000)%10 == 0 {
					param2 = r.get(param2)
				}

				if code != cmdJumpIfFalse && code != cmdJumpIfTrue {
					if (r.get(i)/10_000)%10 == 0 {
						param3 = r.get(param3)
					}
				}
			}
		}

		switch code {
		case cmdAdd:
			r.set(param3, r.get(param1)+r.get(param2))
			i += 4
		case cmdMultiply:
			r.set(param3, r.get(param1)*r.get(param2))
			i += 4
		case cmdSet:
			ip, err := r.rw.ReadValue()
			if err != nil {
				return fmt.Errorf("read failure %w", err)
			}
			r.set(param1, ip)
			i += 2
		case cmdGet:
			err := r.rw.WriteValue(r.get(param1))
			if err != nil {
				return fmt.Errorf("error reading input %w", err)
			}
			i += 2
		case cmdJumpIfTrue:
			if r.get(param1) != 0 {
				i = r.get(param2)
				continue
			}
			i += 3
		case cmdJumpIfFalse:
			if r.get(param1) == 0 {
				i = r.get(param2)
				continue
			}
			i += 3
		case cmdLessThan:
			if r.get(param1) < r.get(param2) {
				r.set(param3, 1)
			} else {
				r.set(param3, 0)
			}
			i += 4
		case cmdEquals:
			if r.get(param1) == r.get(param2) {
				r.set(param3, 1)
			} else {
				r.set(param3, 0)
			}
			i += 4
		case cmdEnd:
			r.rw.Exit()
			return nil
		default:
			return fmt.Errorf("invalid opcode %d at %d", code, i)
		}
	}

	return nil
}

func Run(program []int, input int) ([]int, int, error) {
	vh := &ValueHolder{value: input}
	prog, err := RunWithAmp(program, vh)
	return prog, vh.value, err
}

func RunWithAmp(program []int, rw ioReadWriter) ([]int, error) {
	r := runner{prog: program, rw: rw}
	err := r.run()
	return r.prog, err
}
