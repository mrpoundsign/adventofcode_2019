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

	program := make([]int64, len(record))
	fmt.Println("Program length is", len(program))

	for i, value := range record {
		input, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			log.Fatalf("input is not valid")
		}
		program[i] = input
	}

	_, result, err := intcode.Run(program, 5)
	if err != nil {
		log.Println(err)
	}

	log.Println(result)

}
