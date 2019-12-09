package main

import (
	"os"
	"log"
	"encoding/csv"
	"strconv"
	"fmt"
	"errors"
)

func runWithReplacements(program []int, noun, verb int) (int, error) {
	// Restore gravity
	program[1] = noun
	program[2] = verb

	// log.Println(program)

	programLength := len(program)

	for i := 0; i < programLength; {
		switch(program[i]) {
		case 1:
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
			i += 4
		case 2:
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
			i += 4
		case 99:
			return program[0], nil
		default:
			return program[0], errors.New("invalid opcode")
		}
	}

	return program[0], nil
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

	inputs := make([]int, len(record))
	fmt.Println("Inputs length is", len(inputs))

	for i, value := range record {
		input, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("input is not valid")
		}
		inputs[i] = input
	}

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			program := make([]int, len(inputs))
			copy(program, inputs)
			result, err := runWithReplacements(program, noun, verb)
			if err != nil {
				log.Println(err)
			}
			if result == 19_690_720 {
				fmt.Printf("19690720 found with noun %d verb %d\n", noun, verb)
				fmt.Printf("Result is %d\n", (100*noun) + verb)
				return
			}
		}
	}
	
}