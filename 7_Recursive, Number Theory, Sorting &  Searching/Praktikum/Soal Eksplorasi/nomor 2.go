package main

import "fmt"

// Fungsi getItem untuk menemukan kombinasi item yang jumlahnya mendekati target
func getItem(items []int, target int) []int {
	var result []int
	sum := 0

	for _, item := range items {
		if sum+item <= target {
			result = append(result, item)
			sum += item
		}
		if sum == target {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(getItem([]int{10, 10, 5, 30, 40, 10, 5}, 80)) // [40 30 10]
	fmt.Println(getItem([]int{50, 20, 10, 25, 25}, 100))      // [50 25 25]
}
