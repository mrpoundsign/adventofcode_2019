package intcode

import "fmt"

type ioReadWriter interface {
	ReadValue() int
	WriteValue(int) error // error if we should halt
	Fail()
	End()
}

type mode int

const (
	modeAdd mode = iota + 1
	modeMultiply
	modeSet
	modeGet
	modeJumpIfTrue
	modeJumpIfFalse
	modeLessThan
	modeEquals
	modeEnd = 99
)

func Run(program []int, rw ioReadWriter) ([]int, error) {
	programLength := len(program)

	// input := 0

	for i := 0; i < programLength; {
		code := mode(program[i] % 100)

		// Default to parameter mode
		param1 := i + 1
		param2 := i + 2
		param3 := i + 3

		if code != modeEnd {
			if (program[i]/100)%10 == 0 {
				param1 = program[param1]
			}

			if param1 > programLength {
				rw.Fail()
				return program, fmt.Errorf("attempted to access index out of range at %d (%d)", i, param1)
			}

			if code != modeGet && code != modeSet {
				if (program[i]/1_000)%10 == 0 {
					param2 = program[param2]
				}

				if param2 > programLength {
					rw.Fail()
					return program, fmt.Errorf("attempted to access index out of range at %d (%d %d)", i, param1, param2)
				}

				if code != modeJumpIfFalse && code != modeJumpIfTrue {
					if (program[i]/10_000)%10 == 0 {
						param3 = program[param3]
					}

					if param3 > programLength {
						rw.Fail()
						return program, fmt.Errorf("attempted to access index out of range at %d (%d %d %d)", i, param1, param2, param3)
					}
				}
			}
		}

		switch code {
		case modeAdd:
			program[param3] = program[param1] + program[param2]
			i += 4
		case modeMultiply:
			program[param3] = program[param1] * program[param2]
			i += 4
		case modeSet:
			ip := rw.ReadValue()
			program[param1] = ip
			i += 2
		case modeGet:
			err := rw.WriteValue(program[param1])
			if err != nil {
				rw.Fail()
				return program, fmt.Errorf("error reading input %w", err)
			}
			i += 2
		case modeJumpIfTrue:
			if program[param1] != 0 {
				i = program[param2]
				continue
			}
			i += 3
		case modeJumpIfFalse:
			if program[param1] == 0 {
				i = program[param2]
				continue
			}
			i += 3
		case modeLessThan:
			if program[param1] < program[param2] {
				program[param3] = 1
			} else {
				program[param3] = 0
			}
			i += 4
		case modeEquals:
			if program[param1] == program[param2] {
				program[param3] = 1
			} else {
				program[param3] = 0
			}
			i += 4
		case modeEnd:
			rw.End()
			return program, nil
		default:
			rw.End()
			return program, fmt.Errorf("invalid opcode %d at %d", code, i)
		}
	}

	return program, nil
}
