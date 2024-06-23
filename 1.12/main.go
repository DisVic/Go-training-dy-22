package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	_, _ = fmt.Scan(&a, &b, &c)
	checkTriangle(a, b, c)
}

func checkTriangle(a, b, c int) {
	if a*a+b*b == c*c {
		fmt.Print("Прямоугольный")
	} else {
		fmt.Print("Непрямоугольный")
	}
}
