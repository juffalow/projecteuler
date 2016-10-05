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
		factor := numberToArray(j)

		number = multiply(number, factor)
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

		result = append(result, remainder)

		number = number / 10
	}

	for i := len(result) / 2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}

	return result
}

func multiply(number []int, factor []int) []int {
	var result [][]int
	var zeros []int

	for i := len(factor) - 1; i >= 0; i-- {
		result = append(result, multiplyBy(number, factor[i], zeros))
		zeros = append(zeros, 0)
	}

	return sum(result)
}

func multiplyBy(number []int, factor int, zeros []int) []int {
	var result []int
	var remainder int
	remainder = 0

	result = append(result, zeros...)

	for i := len(number) - 1; i >= 0; i-- {
		result = append(result, ((number[i] * factor) + remainder) % 10)
		remainder = ((number[i] * factor) + remainder) / 10
	}
	if( remainder > 0 ) {
		result = append(result, remainder)
	}

	return revertArray(result)
}

func revertArray(arr []int) []int {
	for i := len(arr) / 2 - 1; i >= 0; i-- {
		opp := len(arr) - 1 - i
		arr[i], arr[opp] = arr[opp], arr[i]
	}
	return arr
}

func sum(numbers [][]int) []int {
	var result []int
	var index int
	var sum int
	var remainder int
	var exists bool

	for index = 0; ; index++ {
		for i := 0; i < len(numbers); i++ {
			if len(numbers[i]) - index - 1 >= 0 {
				exists = true
				sum += numbers[i][len(numbers[i])-index-1]
			}
		}
		if !exists {
			break
		}
		result = append(result, (sum + remainder) % 10)
		remainder = (sum + remainder) / 10
		sum = 0
		exists = false
	}

	if( remainder > 0 ) {
		result = append(result, remainder)
	}

	return revertArray(result)
}
