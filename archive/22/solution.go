package main

import (
    "io/ioutil"
	"strings"
    "fmt"
)

func main() {
	content, err := ioutil.ReadFile("/home/matej/Documents/GoLang/src/projecteuler/archive/22/p022_names.txt")
	
	if err != nil {
	    //Do something
	}

	names := strings.Split(string(content), "\",\"")

	names[0] = strings.TrimPrefix(names[0], "\"")
	names[len(names) - 1] = strings.TrimSuffix(names[len(names) - 1], "\"")

	fmt.Println(names[0])
}
