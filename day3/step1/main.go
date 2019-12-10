package main

import (
	"os"
	"fmt"
	"log"
	"encoding/csv"
	"io"
	"strconv"
	"errors"
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

func pointsFromInstruction(start point, instruction string) ([]point, error) {
	if len(instruction) < 2 {
		return []point{}, errors.New("invalid instruction")
	}

	code := instruction[0]
	count, err := strconv.Atoi(instruction[1:])
	if err != nil {
		return []point{}, errors.New("could not convert instruction")
	}

	out := make([]point, count)

	for i := 1; i <= count; i++ {
		switch code {
		case 'U':
			out[i - 1] = point{X: start.X, Y: start.Y + i}
		case 'D':
			out[i- 1] = point{X: start.X, Y: start.Y - i}
		case 'R':
			out[i- 1] = point{X: start.X + i, Y: start.Y}
		case 'L':
			out[i- 1] = point{X: start.X - i, Y: start.Y}
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
	wires := [][]point{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("could not read line", err)
		}

		curPos := point{0, 0}
		points := []point{}

		for _, instruction := range record {
			p, err := pointsFromInstruction(curPos, instruction)
			if err != nil {
				log.Fatalln("Error processing instruction", err)
			}
			curPos = p[len(p)-1]
			log.Println("pos from instruction", instruction, "is", curPos)
			points = append(points, p...)
		}

		wires = append(wires, points)
	}

	if len(wires) != 2 {
		log.Fatalf("expected 2 wires, got", len(wires))
	}

	crosses := []point{}
	for _, w := range wires[0] {
		for _, w2 := range wires[1] {
			if w == w2 {
				crosses = append(crosses, w)
			}
		}
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