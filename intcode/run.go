package intcode

import "fmt"

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

func Run(program []int, input int, expect int) ([]int, int, error) {
	programLength := len(program)

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

			if code != modeGet && code != modeSet {
				if (program[i]/1_000)%10 == 0 {
					param2 = program[param2]
				}

				if code != modeJumpIfFalse && code != modeJumpIfTrue {
					if (program[i]/10_000)%10 == 0 {
						param3 = program[param3]
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
			program[param1] = input
			i += 2
		case modeGet:
			input = program[param1]
			if expect != -1 && input != expect {
				return program, input, fmt.Errorf("faled validation with code %d", input)
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
			return program, input, nil
		default:
			return program, program[0], fmt.Errorf("invalid opcode %d at %d", code, i)
		}
	}

	return program, input, nil
}
