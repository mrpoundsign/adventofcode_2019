package intcode

import "fmt"

// import "log"

// import "log"

type ioReadWriter interface {
	ReadInput() (int, error)
	WriteOutput(int) // error if we should halt
	Log(...interface{})
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
			rw.Log("add", i, param1, param2, param3, program[param3])
			i += 4
		case modeMultiply:
			program[param3] = program[param1] * program[param2]
			rw.Log("multiply", i, param1, param2, param3, program[param3])
			i += 4
		case modeSet:
			ip, err := rw.ReadInput()
			if err != nil {
				return program, fmt.Errorf("error reading input %w", err)
			}
			program[param1] = ip
			rw.Log("write to", i, param1, ip)
			i += 2
		case modeGet:
			rw.Log("read from", i, param1, program[param1])
			// ip, err := rw.ReadInput()
			// if err != nil {
			// 	return program, fmt.Errorf("faled validation with code %d", input)
			// }
			rw.WriteOutput(program[param1])
			i += 2
		case modeJumpIfTrue:
			rw.Log("jump if true", i, param1, param2)
			if program[param1] != 0 {
				i = program[param2]
				continue
			}
			i += 3
		case modeJumpIfFalse:
			rw.Log("jump if false", i, param1, param2)
			if program[param1] == 0 {
				i = program[param2]
				continue
			}
			i += 3
		case modeLessThan:
			rw.Log("test less than", i, param1, param2, param3)
			if program[param1] < program[param2] {
				program[param3] = 1
			} else {
				program[param3] = 0
			}
			i += 4
		case modeEquals:
			rw.Log("test equals", i, param1, param2, param3)
			if program[param1] == program[param2] {
				program[param3] = 1
			} else {
				program[param3] = 0
			}
			i += 4
		case modeEnd:
			// v, _ := rw.ReadInput()
			return program, nil
		default:
			return program, fmt.Errorf("invalid opcode %d at %d", code, i)
		}
	}

	return program, nil
}
