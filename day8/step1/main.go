package main

import (
	"io"
	"log"
	"os"
	// "strconv"
)

func main() {
	filename := "input.txt"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	input := make([]byte, 25*6)

	minZero := len(input)
	value := 0

	for {
		_, err := inputFile.Read(input)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Panicln(err)
		}

		zeroCount := 0
		curOne := 0
		curTwo := 0
		for _, v := range input {
			switch v {
			case '0':
				zeroCount++
			case '1':
				curOne++
			case '2':
				curTwo++
			}
		}

		if zeroCount > minZero {
			continue
		}

		minZero = zeroCount
		value = curOne * curTwo
	}

	log.Println("minimum zeros", minZero, "1s times 2s", value)
}
