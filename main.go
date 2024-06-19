package main

import "fmt"

func main() {
	// Slices
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}

	// Maps
	capitals := map[string]string{"France": "Paris", "Italy": "Rome"}
	for country, capital := range capitals {
		fmt.Printf("The capital of %s is %s\n", country, capital)
	}
}
