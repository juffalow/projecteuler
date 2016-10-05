package main

import (
	"fmt"
    //"strings"
    //"strconv"
)

/*

n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!

https://projecteuler.net/problem=20

*/

func main() {
	var number []int

	number = []int{1,0,0}

    for j := 99; j >= 2; j-- {
		number = multiply(number, j)
    }

	var sum int
	sum = 0
	for i := 0; i < len(number); i++ {
		sum += number[i]
		fmt.Print(number[i])
	}
	fmt.Println()

	fmt.Println(sum)
}

func numberToArray(number int) []int {
	var result []int

	for ; number > 0; {
		remainder := number % 10

		result = append([]int{remainder}, result...)

		number = number / 10
	}

	return result
}

func multiply(number []int, factor int) []int {
	var result []int
	var remainder int

	remainder = 0

	for i := len(number) - 1; i >= 0; i-- {
		result = append([]int{((number[i] * factor) + remainder) % 10}, result...)
		remainder =  ((number[i] * factor) + remainder) / 10
	}

	if( remainder > 0 ) {
		result = append(numberToArray(remainder), result...)
	}

	return result
}
