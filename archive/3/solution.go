package main

import (
	"fmt"
)

/*

The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?

https://projecteuler.net/problem=3

*/

func main() {
    var number int64
    number = 600851475143
    var i int64
    for i = 3; i < number; i += 2 {
        if isPrime(i) == true && gcd(number, i) == i {
            fmt.Println(i)
        }
    }
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

func gcd(u int64, v int64) int64 {
    var t int64
    t = 0
    for v != 0 {
        t = v
        v = u % v
        u = t
    }
    return u
}