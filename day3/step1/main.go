package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type point struct {
	X int
	Y int
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pointsFromInstruction(instructions []string) (map[point]int, error) {
	out := map[point]int{}
	curStep := 0
	p := point{}

	for _, ins := range instructions {
		if len(ins) < 2 {
			return out, errors.New("invalid instruction")
		}

		code := ins[0]
		count, err := strconv.Atoi(ins[1:])
		if err != nil {
			return out, errors.New("could not convert instruction")
		}

		for i := 0; i < count; i++ {
			curStep ++

			switch code {
			case 'U':
				p = point{X: p.X, Y: p.Y + 1}
			case 'D':
				p = point{X: p.X, Y: p.Y - 1}
			case 'R':
				p = point{X: p.X + 1, Y: p.Y}
			case 'L':
				p = point{X: p.X - 1, Y: p.Y}
			default:
				return out, errors.New("invalid instruction")
			}
			_, ok := out[p]; if !ok {
				out[p] = curStep
			}
		}
	}

	return out, nil
}

func main() {
	csvFile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Could not open CSV file", err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	wires := []map[point]int{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("could not read line", err)
		}

		wire, err := pointsFromInstruction(record)
		if err != nil {
			log.Fatalln("error processing record", err)
		}
		wires = append(wires, wire)
	}

	if len(wires) != 2 {
		log.Fatalln("expected 2 wires, got", len(wires))
	}

	crosses := []point{}

	for w1 := range wires[0] {
		go func(w1 point) {
			_, ok := wires[1][w1]; if ok {
				crosses = append(crosses, w1)
			}
		}(w1)
	}

	log.Println(crosses)

	dist := 0

	for _, pos := range crosses {
		posDist := Abs(pos.X) + Abs(pos.Y)
		if dist == 0 || posDist < dist {
			dist = posDist
		}
	}

	fmt.Println("Closest cross is", dist)
}
