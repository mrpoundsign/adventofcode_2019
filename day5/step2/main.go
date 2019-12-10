package main

import (
	"encoding/csv"
	// "errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Mode int

const (
	ModeAdd Mode = iota + 1
	ModeMultiply
	ModeSet
	ModeGet
	ModeJumpIfTrue
	ModeJumpIfFalse
	ModeLessThan
	ModeEquals
	ModeEnd = 99
)

func intCodeRun(program []int, input int, expect int) (int, error) {
	programLength := len(program)

	for i := 0; i < programLength; {
		code := Mode(program[i] % 100)

		// Default to parameter mode
		param1 := i + 1
		param2 := i + 2
		param3 := i + 3

		switch code {
		case ModeAdd:
			fallthrough
		case ModeMultiply:
			fallthrough
		case ModeEquals:
			fallthrough
		case ModeLessThan:
			if (program[i]/10_000)%10 == 0 {
				param3 = program[param3]
			}
			fallthrough
		case ModeJumpIfFalse:
			fallthrough
		case ModeJumpIfTrue:
			if (program[i]/1_000)%10 == 0 {
				param2 = program[param2]
			}
			fallthrough
		case ModeGet:
			fallthrough
		case ModeSet:
			if (program[i]/100)%10 == 0 {
				param1 = program[param1]
			}
		}

		switch code {
		case ModeAdd:
			program[param3] = program[param1] + program[param2]
			i += 4
		case ModeMultiply:
			program[param3] = program[param1] * program[param2]
			i += 4
		case ModeSet:
			program[param1] = input
			i += 2
		case ModeGet:
			input = program[param1]
			if expect != -1 && input != expect {
				return input, fmt.Errorf("faled validation with code %d", input)
			}
			i += 2
		case ModeJumpIfTrue:
			if program[param1] != 0 {
				i = program[param2]
				continue
			}
			i += 3
		case ModeJumpIfFalse:
			if program[param1] == 0 {
				i = program[param2]
				continue
			}
			i += 3
		case ModeLessThan:
			if program[param1] < program[param2] {
				program[param3] = 1
			} else {
				program[param3] = 0
			}
			i += 4
		case ModeEquals:
			if program[param1] == program[param2] {
				program[param3] = 1
			} else {
				program[param3] = 0
			}
			i += 4
		case ModeEnd:
			return input, nil
		default:
			return program[0], fmt.Errorf("invalid opcode %d at %d", code, i)
		}
	}

	return input, nil
}

func main() {
	csvFile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Could not open CSV file", err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)

	record, err := r.Read()
	if err != nil {
		log.Fatalln("could not read line", err)
	}

	program := make([]int, len(record))
	fmt.Println("Program length is", len(program))

	for i, value := range record {
		input, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("input is not valid")
		}
		program[i] = input
	}

	result, err := intCodeRun(program, 5, 0)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(result)

}
