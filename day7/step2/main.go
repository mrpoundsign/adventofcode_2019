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

	sequences := make([][]int, 5)

	for i := 0; i < 5; i++ {
		sequences[i] = make([]int, 2)
		record, err := r.Read()
		if err != nil {
			log.Fatalln("could not read line", err)
		}
		for j := range record {
			if j > 1 {
				break
			}
			input, err := strconv.Atoi(record[j])
			if err != nil {
				log.Fatalf("input is not valid")
			}
			sequences[i][j] = input + 5
		}
	}

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

	var wg sync.WaitGroup
	var mux sync.Mutex

	big := 0
	bigSequence := 0

	for a := sequences[0][0]; a <= sequences[0][1]; a++ {
		for b := sequences[1][0]; b <= sequences[1][1]; b++ {
			for c := sequences[2][0]; c <= sequences[2][1]; c++ {
				for d := sequences[3][0]; d <= sequences[3][1]; d++ {
					for e := sequences[4][0]; e <= sequences[4][1]; e++ {

						// test for uniqueness of values
						junk := make(map[int]bool)
						junk[a] = true
						junk[b] = true
						junk[c] = true
						junk[d] = true
						junk[e] = true

						if len(junk) < 5 {
							continue
						}

						wg.Add(1)

						go func(s []int) {
							defer wg.Done()
							amps := make([]*intcode.Amplifier, 5)
							var wg2 sync.WaitGroup

							amps[0] = intcode.NewAmplifier([]int{s[0], 0})
							amps[1] = intcode.NewChainAmplifier(s[1], amps[0])
							amps[2] = intcode.NewChainAmplifier(s[2], amps[1])
							amps[3] = intcode.NewChainAmplifier(s[3], amps[2])
							amps[4] = intcode.NewChainAmplifier(s[4], amps[3])
							amps[0].SetPrevAmp(*amps[4])

							for i := 0; i < 5; i++ {
								wg2.Add(1)
								go func(amp *intcode.Amplifier) {
									defer wg2.Done()
									code := make([]int, len(program))
									copy(code, program)
									_, err := intcode.RunWithAmp(code, amp)
									if err != nil {
										log.Println(amp.Phase(), "error", err)
									}
								}(amps[i])
							}

							wg2.Wait()

							mux.Lock()
							defer mux.Unlock()
							value, _ := amps[4].ReadValue()

							if value > big {
								big = value
								seq := s[0]*10_000 + s[1]*1_000 + s[2]*100 + s[3]*10 + s[4]
								bigSequence = seq
							}
						}([]int{a, b, c, d, e})
					}
				}
			}
		}
	}
	wg.Wait()

	log.Println("Max thrust", big, "sequence", bigSequence)

}
