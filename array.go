package main

import "fmt"

func main() {
	lang := [...]string{
		"Python",
		"C++",
		"Ruby",
		"Go",
	}

	for index, row := range lang {
		fmt.Printf("Element-%d is: %s\n", index, row)
	}
}
