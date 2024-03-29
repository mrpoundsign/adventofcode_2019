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

	sequences := make([][]int64, 5)

	for i := 0; i < 5; i++ {
		sequences[i] = make([]int64, 2)
		record, err := r.Read()
		if err != nil {
			log.Fatalln("could not read line", err)
		}
		for j := range record {
			if j > 1 {
				break
			}
			input, err := strconv.ParseInt(record[j], 10, 64)
			if err != nil {
				log.Fatalf("input is not valid")
			}
			sequences[i][j] = input
		}
	}

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

	var wg sync.WaitGroup
	var mux sync.Mutex

	big := int64(0)
	bigSequence := int64(0)

	for a := sequences[0][0]; a <= sequences[0][1]; a++ {
		for b := sequences[1][0]; b <= sequences[1][1]; b++ {
			for c := sequences[2][0]; c <= sequences[2][1]; c++ {
				for d := sequences[3][0]; d <= sequences[3][1]; d++ {
					for e := sequences[4][0]; e <= sequences[4][1]; e++ {
						junk := make(map[int64]bool)
						junk[a] = true
						junk[b] = true
						junk[c] = true
						junk[d] = true
						junk[e] = true

						if len(junk) < 5 {
							continue
						}

						wg.Add(1)

						go func(s []int64) {
							amps := make([]*intcode.Amplifier, 5)

							amps[0] = intcode.NewAmplifier([]int64{s[0], 0})
							amps[1] = intcode.NewChainAmplifier(s[1], amps[0])
							amps[2] = intcode.NewChainAmplifier(s[2], amps[1])
							amps[3] = intcode.NewChainAmplifier(s[3], amps[2])
							amps[4] = intcode.NewChainAmplifier(s[4], amps[3])

							for i := 0; i < 5; i++ {
								go func(amp *intcode.Amplifier) {
									code := make([]int64, len(program))
									copy(code, program)
									p, err := intcode.RunWithIO(code, amp)
									if err != nil {
										log.Println(amp.Phase(), "error", err, p)
									}
								}(amps[i])
							}

							mux.Lock()
							if amps[4].Value() > big {
								big = amps[4].Value()
								value := s[0]*10_000 + s[1]*1_000 + s[2]*100 + s[3]*10 + s[4]
								bigSequence = value
							}
							mux.Unlock()
							wg.Done()
						}([]int64{a, b, c, d, e})
					}
				}
			}
		}
	}
	wg.Wait()

	log.Println("Max thrust", big, "sequence", bigSequence)

}
