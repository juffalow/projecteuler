package main

import (
	"fmt"
)

/*

Euler discovered the remarkable quadratic formula:

n ^ 2 + a * n + b

It turns out that the formula will produce 40 primes for the consecutive integer values 0<=n<=39. However, when
n=40,40^2+40+41=40(40+1)+41 is divisible by 41, and certainly when n=41,41^2+41+41 is clearly divisible by 41.

The incredible formula n^2−79n+1601 was discovered, which produces 80 primes for the consecutive values 0<=n<=79. The
product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:
n2+an+b, where |a|<1000 and |b|<=1000

where |n| is the modulus/absolute value of n
e.g. |11|=11 and |−4|=4

Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for
consecutive values of n, starting with n=0.

https://projecteuler.net/problem=27

*/

func main() {
	primes := sieveOfEratosthenes(1000)

	max := 0
	var n, aMax, bMax int
    for _, b := range primes {
		for a := -999; a < 1000; a += 2 {
			for n = 0; true ; n++ {
				if !isPrime( (n * n) + (a * n) + b) {
					break;
				}
			}
			if n > max {
				max = n
				aMax = a
				bMax = b
			}
		}
	}
	// fmt.Println(max - 1)
	// fmt.Println(aMax)
	// fmt.Println(bMax)
	fmt.Println(aMax * bMax)
}

func sieveOfEratosthenes(n int) []int {
	prime := make([]bool, n + 1)

	for p := 2; p * p <= n; p++ {
		if prime[p] == true {
			for i := p * 2; i <= n; i+= p {
				prime[i] = true
			}
		}
	}

	sieve := make([]int, 0)

	for i := 2; i <= n; i++ {
		if prime[i] == false {
			sieve = append(sieve, i)
		}
	}

	return sieve
}

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }

    if n <= 3 {
        return true
    }

    if n % 2 == 0 || n % 3 == 0 {
        return false
    }

    for i := 5; i * i <= n; i = i + 6 {
        if n % i == 0 || n % (i + 2) == 0 {
            return false
        }
    }
    return true
}
