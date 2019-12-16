package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"sort"
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

func (p point) angleTo(p2 point) float64 {
	deltaX := float64(p2.X) - float64(p.X)
	deltaY := float64(p2.Y) - float64(p.Y)

	deg := (math.Atan2(-deltaY, deltaX) * 180) / math.Pi

	// inRads = (2 * math.Pi) - inRads

	// Pretty sure this can be simplified with a modulus, but can't see it
	if deg <= 90 && deg >= 0 {
		deg = math.Abs(deg - 90)
	} else if deg < 0 {
		deg = math.Abs(deg) + 90
	} else {
		deg = 450 - deg
	}

	return deg
	// return 0
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

	log.Println(most, station, len(asteroids))

	// delete(asteroids, station)

	var targeted map[float64]point
	var n200 point

	removed := 0

Destroy:
	for {
		log.Println("scan", removed)
		targeted = map[float64]point{}

		for dest := range asteroids {
			for check := range asteroids {
				if !checkBlocked(station, dest, check) {
					targeted[station.angleTo(dest)] = dest
					continue
				}
			}
		}

		keys := make([]float64, 0, len(targeted))
		for k := range targeted {
			keys = append(keys, k)
		}
		sort.Float64s(keys)

		for _, k := range keys {
			// log.Println(k)
			delete(asteroids, targeted[k])
			if removed >= 199 {
				n200 = targeted[k]
				break Destroy
			}

			removed++
		}

	}

	log.Println(len(asteroids), n200, removed)
}
