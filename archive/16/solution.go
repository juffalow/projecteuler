package main

import (
	"fmt"
)

/*

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
What is the sum of the digits of the number 2^1000?

https://projecteuler.net/problem=16

*/

func main() {
    var number []int
    
    number = make([]int, 1)
    number[0] = 2
    
    for i := 1; i < 1000; i++ {
        number = sum(number, number)
    }
    
    var length int
    var sum int
    
    length = len(number)
    sum = 0
    
    for i := 0; i < length; i++ {
        sum += number[i]
    }
     
	fmt.Println(sum)
}

func sum(n1 []int, n2 []int) []int {
    var length int
    var result []int
    var sum int
    var remainder int
    
    length = len(n1) - 1
    result = make([]int, length + 1)

    for i := length; i >= 0; i-- {
        sum = n1[i] + n2[i] + remainder
        result[i] = sum - ((sum / 10) * 10)
        remainder = sum / 10
    }
	
    for remainder > 0 {
        result = append([]int{(remainder - ((remainder / 10) * 10))}, result...)
        remainder /= 10
        length++
    }
    
    return result
}