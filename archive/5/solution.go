package main

import (
	"fmt"
)

/*

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

https://projecteuler.net/problem=5

*/

func main() {
    var number int
    
    number = 10
    
    for true {
        if isDivisible(number, 1, 20) == true  {
            break
        }
        number += 10
    }
    
    fmt.Println(number)
}

func isDivisible(number int, from int, to int) bool {
   for i := from; i <= to; i++ {
       if number % i != 0 {
           return false
       }
    }
    
    return true
}
