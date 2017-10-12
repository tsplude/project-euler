package main

import "fmt"

func main() {

     // Collect all positive integers < 1000 that are divided by 3 or 5
     var divisors []int
     for i := 1; i < 1000; i += 1 {
     	 if i % 3 == 0 || i % 5 == 0 {
	    divisors = append(divisors, i)
	 }
     }

     // Find the sum of the divisors
     sum := 0
     for i := 0; i < len(divisors); i += 1 {
     	 sum += divisors[i]
     }

     fmt.Println("Answer:", sum)
}
