# [Maximum path sum I](https://projecteuler.net/problem=18)

[Dijkstra's algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)

Example on [youtube](https://www.youtube.com/watch?v=0nVYi3o161A)

## Little change

We don't need the shortest path, but its opposite. So the default "distance" value
is set to 0 and the greater the value will be, the better for us.

## Process

I created one line of numbers. Each number equals maximum path sum for that node.

E.g.:

|node| |0|1|2|3|4|5|6|7|8|9|
|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
|MAXPATHSUM| |0|0|0|0|0|0|0|0|0|0|

Node index is in parentheses :

|row| | | | | | | | |
|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
|0| | | | |3 (0)| | | |
|1| | | |7 (1)| |4 (2)| | |
|2| | |2 (3)| |4 (4)| |6 (5)| |
|3| |8 (6)| |5 (7)| |9 (8)| |3 (9)|

Max path sum for node(0) is on index 0. Path sum for each node is calculated :

`sum = node(i) + MAXPATHSUM[i - row]` _if it is not the last number in row_

`sum = node(i) + MAXPATHSUM[i - row - 1]` _if it is not the first number in row_

Then if `sum` is greater than MAXPATHSUM[i], this value is overwritten.

_There is a little exception for the very first number, because neither first
calculation condition, nor second could be applied._

## Example

Initial state :

|node| |0|1|2|3|4|5|6|7|8|9|
|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
|MAXPATHSUM| |0|0|0|0|0|0|0|0|0|0|

i = 0

row = 0

sum = 3 + MAXPATHSUM[i - row] = 3 + MAXPATHSUM[0 - 0] = 3 + MAXPATHSUM[0] = 3 + 0 = 3

3 > 0, so MAXPATHSUM[i] / MAXPATHSUM[0] is set to 3

Next step :

|node| |0|1|2|3|4|5|6|7|8|9|
|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
|MAXPATHSUM| |3|0|0|0|0|0|0|0|0|0|

i = 1

row = 1

sum = 7 + MAXPATHSUM[1 - 1] = 7 + 3 = 10

10 > 0, so MAXPATHSUM[1] is set to 10

...

Step for node 4 :

|node| |0|1|2|3|4|5|6|7|8|9|
|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
|MAXPATHSUM| |3|10|7|12|0|0|0|0|0|0|

i = 4

row = 2

sum = 4 + MAXPATHSUM[4 - 2] = 4 + 7 = 11

sum = 4 + MAXPATHSUM[4 - 2 - 1] = 4 + 10 = **14**

14 > 0, so MAXPATHSUM[4] is set to 14
