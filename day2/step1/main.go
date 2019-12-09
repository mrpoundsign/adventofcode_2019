package main

import (
	"os"
	"log"
	"encoding/csv"
	"strconv"
	"fmt"
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

	inputs := make([]int, len(record))
	fmt.Println("Inputs length is", len(inputs))

	for i, value := range record {
		input, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("input is not valid")
		}
		inputs[i] = input
	}

	// Restore gravity
	inputs[1] = 12
	inputs[2] = 2

	inputsLength := len(inputs)

	for i := 0; i < inputsLength; i = i + 4 {
		if inputsLength < i+3 {
			log.Fatalln("Program cannot continue. 0 value is now", inputs[0])
		}
		switch(inputs[i]) {
		case 1:
			inputs[inputs[i+3]] = inputs[inputs[i+1]] + inputs[inputs[i+2]]
		case 2:
			inputs[inputs[i+3]] = inputs[inputs[i+1]] * inputs[inputs[i+2]]
		}
	}

	fmt.Println(inputs[0])
}