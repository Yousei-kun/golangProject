package main

import (
	"errors"
	"fmt"
)

func main() {
	//Problem1
	//scores := []int{10, 5, 8, 9, 7}
	//
	//total := Sum(scores)
	//
	//fmt.Printf("Total: %d\n", total)

	//Problem2
	answer, errorVerdict := Calculation(100, 2, "/")

	if errorVerdict != nil {
		fmt.Printf("%s", errorVerdict)
	}
	fmt.Printf("%d", answer)
}

func Sum(scores []int) int {

	total := 0
	for _, score := range scores {
		total += score
	}

	return total
}

func Calculation(number1 int, number2 int, operation string) (int, error) {
	var calculation int
	var errorVerdict error

	switch operation {
	case "+":
		calculation = number1 + number2
	case "-":
		calculation = number1 - number2
	case "*":
		calculation = number1 * number2
	case "/":
		calculation = number1 / number2
	default:
		errorVerdict = errors.New("unknown operation")
	}

	return calculation, errorVerdict
}
