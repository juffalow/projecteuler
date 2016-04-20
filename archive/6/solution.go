package main

import (
	"fmt"
)

/*

The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 55^2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

https://projecteuler.net/problem=6

*/

func main() {
    var number int
    
    number = 10
    
    var sumOfTheSquares int
    sumOfTheSquares = sumSquareNumbers(number)
    
    var squareOfTheSum int
    squareOfTheSum = sumSequence(number)
    squareOfTheSum *= squareOfTheSum
    
    fmt.Println(squareOfTheSum - sumOfTheSquares)
}

func sumSquareNumbers(n int) int {
    return (n * (n + 1) * ((2 * n) + 1)) / 6
}

func sumSequence(n int) int {
    return (n * (n + 1)) / 2
}