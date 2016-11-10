package main

import (
    "os"
    "io/ioutil"
    "strings"
    "sort"
    "fmt"
)

/*

Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it
into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the
list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the
list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?

https://projecteuler.net/problem=22

*/

func main() {
    arg := os.Args[1]

  	content, err := ioutil.ReadFile(arg + "p022_names.txt")

  	if err != nil {
  	    fmt.Println(err)
  	}

  	names := strings.Split(string(content), "\",\"")

  	names[0] = strings.TrimPrefix(names[0], "\"")
  	names[len(names) - 1] = strings.TrimSuffix(names[len(names) - 1], "\"")

    sort.Strings(names)

    sum := 0

  	for position, name := range names {
        sum += countNameScore(name, position + 1)
    }

    fmt.Println(sum)
}

func countNameScore(name string, position int) int {
    letters := []byte(name)
    var sum int
    sum = 0

    for _, letter := range letters {
        sum += int(letter - 'A') + 1
    }

    return sum * position
}
