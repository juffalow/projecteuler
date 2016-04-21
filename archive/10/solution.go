package main

import (
	"fmt"
)

/*

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.

https://projecteuler.net/problem=10

*/

func main() {
    var max int64
    var sum int64
    var i int64
    
    max = 2000000
    sum = 2
    
    for i : 3; i < max; i += 2 {
        if isPrime(i) == true {
            sum += i
        }
    }
    
    fmt.Println(sum)
}

func isPrime(number int64) bool {
    if(number == 2) || (number == 3) {
        return true
    }
    
    if(number & 1 == 0) || (number % 3 == 0) {
        return false
    }
    
    var i int64
    for i = 3; i * i <= number; i += 2 {
        if number % i == 0 {
            return false
        }
    }
    return true
}