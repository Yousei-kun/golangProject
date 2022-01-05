package main

import "fmt"

func main() {

	//simple for
	for i := 1; i <= 100; i++ {
		fmt.Printf("Coba ke-%d\n", i)
	}

	//foreach
	name := "Ivan is very cool"

	for index, letter := range name {
		fmt.Printf("index: %d letter: %s\n", index, string(letter))
	}

	//quiz
	title := "Golang the best language"

	//problem1
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Printf("%s\n", string(letter))
		}
	}

	fmt.Printf("------------------------------------\n")

	//underline is used to declare UNWANTED variable!!!

	//problem2 - long!!
	for _, letter := range title {
		if (string(letter) == "a") || (string(letter) == "i") || (string(letter) == "u") || (string(letter) == "e") || (string(letter) == "o") {
			fmt.Printf("%s\n", string(letter))
		}
	}

	//problem2 - short
	for _, letter := range title {
		letterInString := string(letter)

		switch letterInString {
		case "a", "i", "u", "e", "o":
			fmt.Printf("%s\n", string(letter))
		}
	}
}
