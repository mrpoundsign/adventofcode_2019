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
	parent   *thingy
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

func (t thingy) distanceTo(name string, from []string) (int, bool) {
	if t.name == name {
		return 0, true
	}

Search:
	for i := range t.children {
		for _, n := range from {
			if n == t.name {
				continue Search
			}
		}

		out, found := t.children[i].distanceTo(name, append(from, t.name))
		if found {
			return out + 1, found
		}
	}

	if t.parent == nil {
		return 0, false
	}

	for _, n := range from {
		if n == t.parent.name {
			return 0, false
		}
	}
	out, found := t.parent.distanceTo(name, append(from, t.name))
	return out + 1, found
}

func main() {
	filename := "input.txt"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	file, err := os.Open(filename)
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
		thing.parent = things[s[0]]
	}

	orbits := 0
	for _, t := range things {
		orbits += t.orbits()
	}
	log.Println("Total orbits:", orbits)
	distance, found := things["YOU"].distanceTo("SAN", []string{})
	if !found {
		log.Fatalf("could not find route form YOU to SAN")
	}
	log.Println("From YOU to SAN:", distance-2)
}
