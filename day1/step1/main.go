package main

import (
	"io"
	"os"
	"log"
	"encoding/csv"
	"strconv"
	"math"
	"fmt"
)

func main() {
	csvFile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Could not open CSV file", err)
	}
	defer csvFile.Close()

	requiredFuel := 0
	r := csv.NewReader(csvFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("could not read line", err)
		}

		if len(record) != 1 {
			log.Fatalln("unexpected line in input", record)
		}

		mass, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalln("could not convert input to integer", record[0])
		}

		requiredFuel = requiredFuel + int(math.Trunc(float64(mass) / 3) - 2)
	}

	fmt.Printf("Required fuel is %d\n", requiredFuel)
}