package main

import (
    "os"
    "io/ioutil"
    "strings"
    "sort"
    "fmt"
)

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
