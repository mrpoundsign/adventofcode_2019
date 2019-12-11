package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type thingy struct {
	name     string
	children []*thingy
}

func (t *thingy) addChild(thing *thingy) {
	t.children = append(t.children, thing)
}

func (t thingy) orbits() int {
	out := len(t.children)
	for _, tt := range t.children {
		out += tt.orbits()
	}
	return out
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	things := map[string]*thingy{}

	for scanner.Scan() {
		var thing *thingy
		s := strings.Split(scanner.Text(), ")")

		_, ok := things[s[0]]
		if !ok {
			things[s[0]] = &thingy{name: s[0]}
		}

		thing, ok = things[s[1]]
		if !ok {
			thing = &thingy{name: s[1]}
			things[s[1]] = thing
		}

		things[s[0]].addChild(thing)
	}

	orbits := 0
	for _, t := range things {
		orbits += t.orbits()
	}
	log.Println("Total orbits:", orbits)
}
