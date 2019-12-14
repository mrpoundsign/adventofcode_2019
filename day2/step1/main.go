package main

import (
	"encoding/csv"
	"fmt"
	"github.com/mrpoundsign/adventofcode_2019/intcode"
	"log"
	"os"
	"strconv"
)

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

	inputs := make([]int64, len(record))
	fmt.Println("Inputs length is", len(inputs))

	for i, value := range record {
		input, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			log.Fatalf("input is not valid")
		}
		inputs[i] = input
	}

	// Restore gravity
	inputs[1] = 12
	inputs[2] = 2

	program, _, err := intcode.Run(inputs, 0)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(program[0])
}
