package main

import (
	"fmt"
)

/*

If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.

https://projecteuler.net/problem=17

*/

var numbers = map[int]string{}

func main() {
	initializeNumbers()

	lettersCount := 0

	for i := 1; i <= 1000; i++ {
		number := createStringNumber(i)
		fmt.Println(number)
		lettersCount += len(number)
    }

	fmt.Println(lettersCount)
}

func initializeNumbers() {
	numbers[1] = "one"
	numbers[2] = "two"
	numbers[3] = "three"
	numbers[4] = "four"
	numbers[5] = "five"
	numbers[6] = "six"
	numbers[7] = "seven"
	numbers[8] = "eight"
	numbers[9] = "nine"
	numbers[10] = "ten"
	numbers[11] = "eleven"
	numbers[12] = "twelve"
	numbers[13] = "thirteen"
	numbers[14] = "fourteen"
	numbers[15] = "fifteen"
	numbers[16] = "sixteen"
	numbers[17] = "seventeen"
	numbers[18] = "eighteen"
	numbers[19] = "nineteen"
	numbers[20] = "twenty"
	numbers[30] = "thirty"
	numbers[40] = "forty"
	numbers[50] = "fifty"
	numbers[60] = "sixty"
	numbers[70] = "seventy"
	numbers[80] = "eighty"
	numbers[90] = "ninety"
	numbers[100] = "hundred"
	numbers[1000] = "thousand"

	return
}

func createStringNumber(n int) string {
	var number string
	number = ""

	if( n >= 1000 ) {
		number = numbers[n / 1000] + numbers[1000]
		n = n - ( ( n / 1000 ) * 1000 )
	}

	if n >= 100 {
		number = numbers[n / 100] + numbers[100]
		n = n - ( ( n / 100 ) * 100 )
	}

	if( n > 0 && n < 100 ) {
		if len(number) > 0 {
			number += "and"
		}

		if _, ok := numbers[n]; ok {
			number += numbers[n]
		} else {
			number += numbers[(n / 10) * 10] + numbers[n % 10]
		}
	}

	return number
}
