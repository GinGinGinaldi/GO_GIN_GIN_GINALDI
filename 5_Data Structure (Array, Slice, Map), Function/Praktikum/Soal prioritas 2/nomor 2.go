package main

import "fmt"

// Fungsi untuk memeriksa apakah bilangan prima
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Fungsi untuk mengelompokkan bilangan prima dan non-prima
func groupPrime(numbers []int) []any {
	primes := []int{}
	nonPrimes := []int{}

	for _, num := range numbers {
		if isPrime(num) {
			primes = append(primes, num)
		} else {
			nonPrimes = append(nonPrimes, num)
		}
	}

	if len(primes) == 0 {
		return []any{numbers} // Tidak ada bilangan prima, kembalikan input asli
	}

	return []any{primes, nonPrimes}
}

func main() {
	fmt.Println(groupPrime([]int{2, 3, 4, 5, 6, 7, 8})) // Output: [[2,3,5,7],4,6,8]
	fmt.Println(groupPrime([]int{15, 24, 3, 9, 5}))     // Output: [[3,5],15,24,9]
	fmt.Println(groupPrime([]int{4, 8, 9, 12}))         // Output: [4,8,9,12]
}
