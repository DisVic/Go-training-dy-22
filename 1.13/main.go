package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	fmt.Print(findDigitSquare(n))
}

func findDigitSquare(number int) int {
	for {
		summ := 0
		for number > 0 {
			summ += number % 10
			number /= 10
		}
		if summ < 10 {
			return summ
		} else {
			number = summ
		}
	}
}
