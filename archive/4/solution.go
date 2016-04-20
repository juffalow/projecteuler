package main

import (
	"fmt"
    "math"
)

/*

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.

https://projecteuler.net/problem=4

*/

func main() {
    digits := 3
    min := int(math.Pow(10, float64(digits - 1)))
    max := int(math.Pow(10, float64(digits)))
    
    var i int
    var j int
    
    var number1 int
    var number2 int
    var palindrome int
    palindrome = 0
    
    for i = min; i < max; i++ {
        for j = min; j < max; j++ {
            if isPalindrome(i * j) == true && i * j > palindrome {
                number1 = i
                number2 = j
                palindrome = i * j
            }
        }
    }
    
    fmt.Println(number1)
    fmt.Println(number2)
    fmt.Println(number1 * number2)
}

func isPalindrome(number int) bool {
    palindrome := number
    reverse := 0
    
    for palindrome != 0 {
        remainder := palindrome % 10
        reverse = reverse * 10 + remainder
        palindrome = palindrome / 10
    }
    
    return number == reverse
}
