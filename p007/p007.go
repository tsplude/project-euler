package main

import "fmt"

func is_prime(x int, factors []int) bool {
	for _, v := range factors {
		if x%v == 0 {
			return false
		}
	}
	return true
}

func main() {

	target_prime := 10001

	primes := []int{2, 3}
	curr := 5
	for {
		if len(primes) == target_prime {
			break
		}
		if is_prime(curr, primes) {
			primes = append(primes, curr)
		}
		curr += 2
	}

	fmt.Println("Answer:", primes[len(primes)-1])
}
