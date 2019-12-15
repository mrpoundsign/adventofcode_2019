package main

import (
	"encoding/csv"
	"fmt"
	"github.com/mrpoundsign/adventofcode_2019/intcode"
	"log"
	// "math"
	"os"
	"strconv"
)

func main() {
	filename := "input.csv"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
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

	v := intcode.NewArcade(0)
	_, err = intcode.RunWithIO(program, v)
	if err != nil {
		log.Println(err)
	}

	log.Println(v.BlockCount())

}
