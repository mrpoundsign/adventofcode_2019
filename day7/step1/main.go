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
			sequences[i][j] = input
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

						go func(sequence []int) {
							var amp intcode.Amplifier

							for i := 0; i < len(sequence); i++ {

								code := make([]int, len(program))
								copy(code, program)

								switch i {
								case 0:
									amp = intcode.NewAmplifier(sequence[i], 0)
								default:
									amp = intcode.NewAmplifier(sequence[i], amp.Value())
								}

								_, err := intcode.Run(code, &amp)
								if err != nil {
									log.Println(err)
								}

								mux.Lock()
								if amp.Value() > big {
									big = amp.Value()
									value := sequence[0]*10_000 + sequence[1]*1_000 + sequence[2]*100 + sequence[3]*10 + sequence[4]
									bigSequence = value
								}
								mux.Unlock()
							}
							wg.Done()
						}([]int{a, b, c, d, e})
					}
				}
			}
		}
	}
	wg.Wait()

	log.Println("Max thrust", big, "sequence", bigSequence)

}
