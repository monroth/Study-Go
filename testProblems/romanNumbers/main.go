package main

import "fmt"

var literals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func parseFromRoman(s string) []int {

	numbers := []int{}
	for _, char := range s {
		numbers = append(numbers, literals[char])
	}
	numbers = append(numbers, 0)
	return numbers
}

func calculateRoman(numbers []int) int {
	value := 0
	for i, number := range numbers {
		if number == 0 {
			return value
		}
		if number < numbers[i+1] {
			value = value - number
		} else {
			value = value + number
		}

	}
	return -1
}

func romanToInt(s string) int {
	return calculateRoman(parseFromRoman(s))
}

func main() {
	fmt.Println(romanToInt("XXIV"))

}
