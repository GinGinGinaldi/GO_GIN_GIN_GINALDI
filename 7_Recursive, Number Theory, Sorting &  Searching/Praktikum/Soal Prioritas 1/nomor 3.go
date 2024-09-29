package main

import "fmt"

// Fungsi getSequence menghitung jumlah total berdasarkan pola deret segitiga
func getSequence(n int) int {
	// Menggunakan rumus deret aritmatika untuk menghitung jumlah
	return n * (n + 1) / 2
}

func main() {
	fmt.Println(getSequence(4))   // 10
	fmt.Println(getSequence(15))  // 120
	fmt.Println(getSequence(100)) // 5050
}
