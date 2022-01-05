package main

import "fmt"

func main() {
	//long declaration
	//var mapLanguageProficiency map[string]int
	//mapLanguageProficiency = map[string]int{}
	//
	//mapLanguageProficiency["english"] = 9
	//mapLanguageProficiency["japanese"] = 5
	//mapLanguageProficiency["chinese"] = 2

	//short declaration
	mapLanguageProficiency := map[string]int{
		"english":  9,
		"japanese": 5,
		"chinese":  2,
	}

	for key, value := range mapLanguageProficiency {
		fmt.Printf("My languange proficiency in %s is %d\n", key, value)
	}

	fmt.Println("==============================")

	delete(mapLanguageProficiency, "chinese")

	for key, value := range mapLanguageProficiency {
		fmt.Printf("My languange proficiency in %s is %d\n", key, value)
	}
}
