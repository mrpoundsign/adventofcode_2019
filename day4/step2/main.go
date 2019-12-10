package main

import "log"

func main() {
	count := 0
Finder:
	for i := 158126; i <= 624574; i++ {
		// for i := 223444; i <= 223444; i++ {
		double := false
		repeatDigitCount := 0
		lDigit := 10
		number := i
		for number > 0 {
			digit := number % 10

			if digit == lDigit {
				repeatDigitCount++
			} else {
				if repeatDigitCount == 1 {
					double = true
				}
				repeatDigitCount = 0
			}

			if digit > lDigit {
				continue Finder
			}

			lDigit = digit
			number = number / 10
		}

		if double || repeatDigitCount == 1 {
			log.Println(i)
			count++
		}
		// break
	}
	log.Println(count)
}
