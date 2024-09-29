package main

import (
	"fmt"
	"strconv"
)

func generateBinary(n int) []string {
	result := []string{}

	for i := 0; i <= n; i++ {
		binary := strconv.FormatInt(int64(i), 2)
		result = append(result, binary)
	}
	return result
}
func main() {
	n := 3
	result := generateBinary(n)
	fmt.Println(result) // Output: [0, 1, 10, 11]
}
