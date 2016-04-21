package main

import (
	"fmt"
)

/*

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a2 + b2 = c2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

https://projecteuler.net/problem=9

*/

func main() {
    var c int
    for a := 1; a < 400; a++ {
        for b := a + 1; b < 400; b++ {
            c = 1000 - a - b
            if (a * a) + (b * b) == (c * c) {
                fmt.Println(a)
                fmt.Println(b)
                fmt.Println(c)
                fmt.Println(a * b * c)
                return
            }
        }
    }
}

/* Very bad solution but working one */
/*
func main() {
    var perfectSquare int
    for i := 1; i < 400; i++ {
        for j := 1; j < 400; j++ {
            perfectSquare = getPerfectSquare((i * i) + (j * j))
            if perfectSquare != 0 && i + j + perfectSquare == 1000 {
                fmt.Println(i)
                fmt.Println(j)
                fmt.Println(perfectSquare)
                fmt.Println(i * j * perfectSquare)
                return
            }
        }
    }
}

func getPerfectSquare(number int) int {
    var current int
    var i int
    
    current = 0
    i = 0
    
    for current < number {
        i++
        current = i * i
    }
    
    if current == number {
        return i
    }
    return 0
}
*/