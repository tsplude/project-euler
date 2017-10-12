package main

import "fmt"


func next_fib() func() int {
     prev, curr := 0, 1
     return func() int {
     	    ret := prev
	    prev, curr = curr, prev + curr
	    return ret
     }
}

func main() {

     sum := 0
     fibs := next_fib()
     for {
     	 next := fibs()
	 if next > 4000000 {
	    break
	 } else if next % 2 == 0 {
	    sum += next
	 }
     }

     fmt.Println("Answer:", sum)
}