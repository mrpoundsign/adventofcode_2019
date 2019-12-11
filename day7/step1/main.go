package main

import (
	"encoding/csv"
	"fmt"
	"github.com/mrpoundsign/adventofcode_2019/intcode"
	"log"
	"os"
	"strconv"
	"sync"
)

type amplifier struct {
	phase     int
	in        chan int
	out       chan int
	value     int
	sentCount int
}

func (a amplifier) ReadInput() (int, error) {
	value, ok := <-a.in
	if !ok {
		return a.value, nil
	}
	a.value = value
	a.Log("io read", a.value)
	return a.value, nil
}

func (a *amplifier) WriteOutput(i int) {
	a.value = i
	a.Log("io write", a.value)
	if a.sentCount == 2 {
		close(a.out)
		a.sentCount++
		return
	}
	a.Log("io send", a.value)
	a.sentCount++
	a.out <- i
}

func (a amplifier) Log(v ...interface{}) {
	log.Println("phase", a.phase, v)
}

func newAmplifierPhase(phase int) *amplifier {
	ch := make(chan int, 2)

	return newAmplifierWithChan(phase, ch, make(chan int, 2))
}

func newAmplifierWithChan(phase int, in, out chan int) *amplifier {
	in <- phase
	return &amplifier{phase: phase, in: in, out: out}
}

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

	program := make([]int, len(record))
	fmt.Println("Program length is", len(program))

	for i, value := range record {
		input, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("input is not valid")
		}
		program[i] = input
	}

	// sequence := []int{4, 3, 2, 1, 0}
	// sequence := []int{4, 3}
	sequence := []int{0, 1, 2, 3, 4}

	var wg sync.WaitGroup

	// first := make(chan int)
	last := make(chan int)

	var prevOut chan int
	// var prevIn chan int
	// var last chan int

	for i := 0; i < len(sequence); i++ {
		var amp *amplifier

		code := make([]int, len(program))
		copy(code, program)

		switch i {
		case 0:
			// log.Println("Found first")
			amp = newAmplifierPhase(sequence[i])
			go func() {
				amp.in <- 0
				close(amp.in)
			}()
		case len(sequence) - 1:
			// log.Println("Found last at", i, sequence[i])
			amp = newAmplifierWithChan(sequence[i], prevOut, last)
		default:
			// log.Println("amp", i)
			amp = newAmplifierWithChan(sequence[i], prevOut, make(chan int, 2))
		}

		// prevIn = amp.in
		prevOut = amp.out

		wg.Add(1)
		go func(code []int, amp *amplifier) {
			log.Println("starting amp", amp.phase)
			_, err := intcode.Run(code, amp)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("end amp", amp.phase)
			wg.Done()
		}(code, amp)
	}

	wg.Add(1)
	go func() {
		log.Println("result", <-last)
		wg.Done()
	}()

	wg.Wait()

}
