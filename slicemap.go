package main

import "fmt"

func main() {
	sliceStudent := []map[string]string{
		{"name": "Jessica", "score": "A"},
		{"name": "Ivan", "score": "AB"},
		{"name": "Steven", "score": "C"},
	}

	fmt.Println(sliceStudent)

	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	var good_scores []int

	total_sum := 0

	for _, value := range scores {
		total_sum += value
		if value >= 90 {
			good_scores = append(good_scores, value)
		}
	}

	mean := float64(total_sum) / float64(len(scores))
	fmt.Printf("Mean dari ujian adalah: %f\n", mean)

	fmt.Printf("Terdapat %d nilai bagus:\n", len(good_scores))
	for _, value := range good_scores {
		fmt.Printf("%d\n", value)
	}
}
