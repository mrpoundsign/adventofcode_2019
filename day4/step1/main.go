package main

import "log"

func main() {
	count := 0
	Finder:
	for i := 158126; i <= 624574; i++ {
		double := false
		lDigit := 10
		number := i
		for j := 0; j < 6; j++ {
			digit := number % 10
			// log.Println(digit)
			if digit > lDigit {
				continue Finder
			}
			if digit == lDigit {
				double = true
			}
			lDigit = digit
			number = number / 10
		}
		if double {
			count ++
		}
		// break
	}
	log.Println(count)
}