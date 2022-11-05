package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFromFile(fileStrings []string, count map[string]int) {
	for _, word := range fileStrings {
		if _, ok := count[word]; ok {
			count[word]++
		} else {
			count[word] = 1
		}
	}
}

func writeBigAmounts(count map[string]int) {
	for word, amount := range count {
		if amount > 1 {
			fmt.Println(word, amount)
		}
	}
}

func main() {
	count := make(map[string]int)
	args := os.Args[1:]
	for _, arg := range args {
		dat, err := os.ReadFile(`./` + arg)
		check(err)
		readFromFile((strings.Split(string(dat), "\n")), count)
	}
	writeBigAmounts(count)
}
