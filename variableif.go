package main

import (
	"fmt"
	"golangProject/calculation"
)

func main() {
	result := calculation.Add(5, 1)
	fmt.Println(result)

	result2 := calculation.Multiply(10, 4)
	fmt.Println(result2)

	if result2 > 10 {
		fmt.Println("Kamu tua ya")
	} else {
		fmt.Println("Ndak")
	}
}
