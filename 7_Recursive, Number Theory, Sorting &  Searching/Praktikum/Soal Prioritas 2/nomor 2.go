package main

import "fmt"

// Fungsi untuk mencari faktor prima dari bilangan
func primeFactors(n int) {
	// Loop untuk faktor 2
	for n%2 == 0 {
		fmt.Print(2, " ")
		n /= 2
	}

	// Loop untuk faktor ganjil mulai dari 3
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			fmt.Print(i, " ")
			n /= i
		}
	}

	// Jika n adalah bilangan prima lebih besar dari 2
	if n > 2 {
		fmt.Print(n)
	}
	fmt.Println()
}

func main() {
	// Test case
	primeFactors(20)   // Output: 2 2 5
	primeFactors(75)   // Output: 3 5 5
	primeFactors(1024) // Output: 2 2 2 2 2 2 2 2 2 2
}
