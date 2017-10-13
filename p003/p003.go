package main

import (
	"fmt"
	"math"
)

func check_prime(x int, divs []int) bool {
	for i := 0; i < len(divs); i += 1 {
		if x%divs[i] == 0 {
			return false
		}
	}
	return true
}

func factors(x int) []int {
	var divs []int
	for i := 2; i < int(math.Sqrt(float64(x))); i += 1 {
		if x%i == 0 && check_prime(i, divs) {
			divs = append(divs, i)
		}
	}
	return divs
}

func main() {

	// Find factors
	divs := factors(600851475143)

	fmt.Println("Answer:", divs[len(divs)-1])
}
