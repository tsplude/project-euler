package main

import (
    "fmt"
    "strconv"
    "sort"
)

func check_palindrome(x string) bool {
     for i := 0; i < len(x)/2; i += 1 {
         if string(x[i]) != string(x[len(x) - i - 1]) {
	     return false
	 }
     }
     return true
}

func main() {

     var pals []int

     for a := 100; a < 1000; a += 1 {
         for b := a; b < 1000; b += 1 {
     	     if check_palindrome(strconv.Itoa(a * b)) {
     	         pals = append(pals, a * b)
      	     }
         }
     }

     sort.Ints(pals)
     fmt.Println("Answer:", pals[len(pals)-1])
}