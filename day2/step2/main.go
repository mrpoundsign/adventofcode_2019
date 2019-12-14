package main

import (
	"encoding/csv"
	"fmt"
	"github.com/mrpoundsign/adventofcode_2019/intcode"
	"log"
	"os"
	"strconv"
)

func runWithReplacements(program []int, noun, verb int) (int, error) {
	// Restore gravity
	program[1] = noun
	program[2] = verb

	program, _, err := intcode.Run(program, 0)
	return program[0], err
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
				fmt.Printf("Result is %d\n", (100*noun)+verb)
				return
			}
		}
	}

}
