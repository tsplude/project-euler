package main

import (
    "fmt"
)

func check_solution(x int, factors []int) bool {
     for _, f := range factors {
         if x % f != 0 {
	     return false
	 }
     }
     return true
}

func main() {

     factors := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

     start := 20
     for {
         if check_solution(start, factors) {
	     break
	 }
	 start += 2
     }

     fmt.Println("Answer:", start)
}