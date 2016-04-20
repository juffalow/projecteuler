package main

import (
	"fmt"
)

/*

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

https://projecteuler.net/problem=7

*/

func main() {
    var number int
    
    number = 10001
    
    var i int
    i = 3
    var nth int
    nth = 1
    
    for true {
        if isPrime(i) == true {
            nth++
        }
        
        if  nth == number {
            break
        }

	    i += 2
    }
    
    fmt.Println(i)
}

func isPrime(number int) bool {
    if(number == 2) || (number == 3) {
        return true
    }
    
    if(number % 2 == 0) || (number % 3 == 0) {
        return false
    }
    
    var i int
    for i = 3; i * i <= number; i += 2 {
        if number % i == 0 {
            return false
        }
    }
    return true
}