package main

import "fmt"

// Fungsi untuk menghitung kuadrat dari bilangan
func getDiamondSeq(n int) int {
	return n * n
}

func main() {
	// Test case
	fmt.Println(getDiamondSeq(4))   // Output: 16
	fmt.Println(getDiamondSeq(25))  // Output: 625
	fmt.Println(getDiamondSeq(100)) // Output: 10000
}
