package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operands = map[string]func(int64, int64) int64{
	"+": func(a, b int64) int64 { return a + b },
	"-": func(a, b int64) int64 { return a - b },
	"*": func(a, b int64) int64 { return a * b },
	"/": func(a, b int64) int64 { return a / b },
	//	'C': 100,
	//	'D': 500,
	//	'M': 1000,
}

func bop(parsedPointer *[]int64, element string) error {
	if operand, ok := operands[element]; ok {
		l := len(*parsedPointer) - 1
		if l > 0 {
			if (*parsedPointer)[l-1] == 0 && element == "/" {
				return errors.New("Division by zero")
			}
			(*parsedPointer)[l-1] = operand((*parsedPointer)[l], (*parsedPointer)[l-1])
			*parsedPointer = (*parsedPointer)[:l]
			return nil
		} else {
			return errors.New("Not enough arguments for operand")
		}

	} else {
		return errors.New("Unknown operand")
	}

}

func evalString(s string) {
	parsed := []int64{}
	halfParsed := strings.Split(s, " ")
	if len(halfParsed) == 0 {
		return
	}
	for _, element := range halfParsed {
		if number, err := strconv.ParseInt(element, 10, 64); err == nil {
			parsed = append(parsed, number)
		} else {
			err := bop(&parsed, element)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	if len(parsed) > 1 {
		fmt.Println("Too many numbers left after operands")
		return
	}
	fmt.Println(parsed[0])

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		if text == "q" {
			return
		}
		evalString(text)
	}

}
