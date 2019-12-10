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
	ModeEnd = 99
)

func intCodeRun(program []int) (int, error) {

	programLength := len(program)
	input := 1

	for i := 0; i < programLength; {
		code := Mode(program[i] % 100)

		// Default to parameter mode
		param1 := i + 1
		param2 := i + 2
		param3 := i + 3

		// Optionally use position mode
		if (program[i]/100)%10 == 0 {
			param1 = program[param1]
		}

		if code < 3 {
			if (program[i]/1_000)%10 == 0 {
				param2 = program[param2]
			}

			if (program[i]/10_000)%10 == 0 {
				param3 = program[param3]
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
			if input != 0 {
				return input, fmt.Errorf("faled validation with code %d", input)
			}
			i += 2
		case ModeEnd:
			return input, nil
		default:
			return program[0], fmt.Errorf("invalid opcode %d", code)
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

	result, err := intCodeRun(program)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(result)

}
