package main

import (
	"fmt"
)

/*

The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.

Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144

The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

https://projecteuler.net/problem=25

*/

func main() {
	var number []int
	var result []int
	var tmp []int

	number = []int{1}
	result = []int{1}

	index := 1

    for ; len(result) < 1000; {
		tmp = result
		result = sum(number, result)
		number = tmp

		index++;
    }

	fmt.Println(index + 1)
}

func sum(number1 []int, number2 []int) []int {
	var result []int
	var remainder int
	var num int

	remainder = 0

	j := len(number1) - 1

	for i := len(number2) - 1; i >= 0; i-- {
		if( j >= 0 ) {
			num = number1[j]
		} else {
			num = 0
		}

		result = append([]int{((num + number2[i]) + remainder) % 10}, result...)
		remainder =  ((num + number2[i]) + remainder) / 10

		j--
	}

	if( remainder > 0 ) {
		result = append(numberToArray(remainder), result...)
	}

	return result
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
