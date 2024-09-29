package main

import "fmt"

// Fungsi faktorial untuk menghitung n!
func faktorial(n int) int {
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi catalan untuk menghitung bilangan Catalan
func catalan(n int) int {
	return faktorial(2*n) / (faktorial(n+1) * faktorial(n))
}

func main() {
	fmt.Println(catalan(7))  // 429
	fmt.Println(catalan(10)) // 16796
}
