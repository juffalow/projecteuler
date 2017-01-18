# [Number spiral diagonals](https://projecteuler.net/problem=28)

The four corners can be written as :
n ^ 2 - 3 * n + 3
n ^ 2 - 2 * n + 2
n ^ 2 - n + 1
n ^ n

After summing them together :
4 * n ^ 2 - 6 * n + 6

With this formula, we can calculate sum of the corners for any odd-length square.

So if we want to calculate the sum of the diagonals, we have to sum up
the sum of the corners for every odd-length square.

Example for 5 by 5 spiral :

1
4 * 3 ^ 2 - 6 * 3 + 6 = 24
4 * 5 ^ 2 - 6 * 5 + 6 = 76

1 + 24 + 76 = 101

That means, we have to use the formula only ( n - 1 ) / 2 times.

Next step...

Now we have to sum up _s_ = ( n - 1 ) / 2 times number _i_ starting from 1 to _s_ :

4 * (2i +1) ^ 2 - 6(2i + 1) + 6 ( and + 1 as it is the first square )

Simplified :

16i2 + 4i + 4 ( and + 1)

Using summation of polynomial expressions :

(16s(s + 1)(2s + 1)) / 6 + (4s(s+1) / 2) + 4s ( and + 1)

Simplified :

(16s3 + 30s2 + 26s + 3) / 3

That means :

length := 1001
s := (length - 1) / 2
(16 * s * s * s + 30 * s * s + 26 * s + 3) / 3
