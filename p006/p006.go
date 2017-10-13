package main

import (
	"fmt"
)

func main() {
	end := 100.

	var sum1, sum2 uint64 = 0., 0.
	for i := 0.; i <= end; i += 1. {
		sum1 += uint64(i * i)
		sum2 += uint64(i)
	}

	fmt.Println("Answer:", (sum2*sum2 - sum1))
}
