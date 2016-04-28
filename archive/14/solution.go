package main

import (
	"fmt"
    "time"
)

/*

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.


https://projecteuler.net/problem=14

*/

func main() {
    var max int
    var steps int
    var startingNumber int
    
    max = 0
    
    start := time.Now()
    
	for i := 1; i < 1000000; i++ {
        steps = countSteps(i)
        if steps > max {
            max = steps
            startingNumber = i
        }
	}
    
    fmt.Printf("Starting number : %d\n", startingNumber)
    fmt.Printf("Steps : %d\n", max)
    
    elapsed := time.Since(start)
    fmt.Printf("It took : %f\n", elapsed.Seconds())
}

func countSteps(number int) int {
    var count int
    count = 1
    
    for number > 1 {        
        if number & 1 == 1 {
            number = (3 * number) + 1
        } else {
            number /= 2
        }
        count++
    }
    
    return count
}