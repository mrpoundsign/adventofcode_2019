package main

import (
	"encoding/csv"
	"fmt"
	"github.com/mrpoundsign/adventofcode_2019/intcode"
	"image"
	"image/color"
	"image/png"
	"log"
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

	v := intcode.NewRover(1)
	_, err = intcode.RunWithIO(program, v)
	if err != nil {
		log.Println(err)
	}

	log.Println(v.PaintCount())

	paints := v.Paints()
	minX := 0
	minY := 0
	maxX := 0
	maxY := 0
	for p := range paints {
		if p.X < minX {
			minX = p.X
		}

		if p.X > maxX {
			maxX = p.X
		}

		if p.Y < minY {
			minY = p.Y
		}

		if p.Y > maxY {
			maxY = p.Y
		}
	}

	upLeft := image.Point{X: -100, Y: -100}
	// lowRight := image.Point{X: maxX - minX, Y: maxY - minY}
	lowRight := image.Point{X: 100, Y: 100}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	black := color.RGBA{0, 0, 0, 0xff}
	white := color.RGBA{255, 255, 255, 0xff}

	for p := range paints {
		switch paints[p] {
		case 0:
			img.Set(p.X, p.Y, black)
		case 1:
			img.Set(p.X, p.Y, white)
		}
	}

	f, _ := os.Create("step2.png")
	png.Encode(f, img)
	log.Println("wrote step2.png")

}
