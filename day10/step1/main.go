package main

import (
	"bufio"
	// "fmt"
	"io"
	"log"
	"math"
	"os"
)

type point struct {
	X int
	Y int
}

func (p point) sameDirection(dest, check point) bool {
	if dest.X <= p.X && check.X > p.X {
		return false
	}

	if dest.Y <= p.Y && check.Y > p.Y {
		return false
	}

	if dest.X >= p.X && check.X < p.X {
		return false
	}

	if dest.Y >= p.Y && check.Y < p.Y {
		return false
	}

	return true
}

func (p point) ticksTo(p2 point) (float64, float64) {
	diffX := float64(p.X - p2.X)
	diffY := float64(p.Y - p2.Y)

	if diffX == diffY {
		return 1, 1
	}

	diff := 0.0

	if diffY > diffX {
		diff = diffX / diffY
		return diff, 1
	}

	diff = diffY / diffX
	return 1, diff
}

func checkBlocked(src, dest, check point) bool {
	if check == src || check == dest {
		return false
	}

	dist := math.Abs(float64(src.X-dest.X)) + math.Abs(float64(src.Y-dest.Y))
	dist2 := math.Abs(float64(src.X-check.X)) + math.Abs(float64(src.Y-check.Y))

	if dist <= dist2 {
		return false
	}

	if !src.sameDirection(dest, check) {
		return false
	}

	dirX, dirY := src.ticksTo(dest)
	dir2X, dir2Y := src.ticksTo(check)

	if dirX == dir2X && dirY == dir2Y {
		return true
	}

	return false
}

func main() {

	log.Println(os.Args)
	fileName := "input.txt"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Panicln("could not read file", fileName, err)
	}

	x := 0
	y := 0

	// map of asteroids with visible count
	asteroids := map[point]int{}

	r := bufio.NewReader(file)

	for {
		l, p, err := r.ReadLine()

		if err == io.EOF {
			log.Println("EOF reached")
			break
		}

		if err != nil {
			log.Panicln("error reading file", err)
		}

		for _, b := range l {
			if b == '#' {
				asteroids[point{X: x, Y: y}] = 0
			}

			x++
		}

		if !p {
			x = 0
			y++
		}
	}

	for src := range asteroids {
	DestCheck:
		for dest := range asteroids {
			if src == dest {
				continue
			}

			for check := range asteroids {
				if checkBlocked(src, dest, check) {
					continue DestCheck
				}
			}
			asteroids[src]++
		}
	}

	most := 0
	station := point{}

	for p, i := range asteroids {
		if i > most {
			most = i
			station = p
		}
	}

	log.Println(most, station)
}
