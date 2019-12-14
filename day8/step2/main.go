package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
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

	width := 25
	height := 6

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	black := color.RGBA{0, 0, 0, 0xff}
	white := color.RGBA{255, 255, 255, 0xff}
	transparent := color.RGBA{0, 0, 0, 0x00}

	for {
		newImg := image.NewRGBA(image.Rectangle{upLeft, lowRight})
		_, err := inputFile.Read(input)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Panicln(err)
		}

		for i, v := range input {
			x := i % 25
			y := (i / 25) % 6

			switch v {
			case '0':
				newImg.Set(x, y, black)
			case '1':
				newImg.Set(x, y, white)
			case '2':
				newImg.Set(x, y, transparent)
			}
		}

		draw.Draw(newImg, img.Bounds(), img, image.Point{X: 0, Y: 0}, draw.Over)
		img = newImg
	}

	f, _ := os.Create("step2.png")
	png.Encode(f, img)
	log.Println("wrote step2.png")
}
