package main

import "fmt"

// Fungsi untuk menjumlahkan semua bilangan Fibonacci hingga ke-n
func fibX(n int) int {
	if n < 0 {
		return 0 // Jika n negatif, return 0
	}

	a, b := 0, 1 // Menginisialisasi dua bilangan Fibonacci pertama
	sum := 0     // Variabel untuk menyimpan hasil penjumlahan

	// Menghitung jumlah bilangan Fibonacci hingga ke-n
	for i := 0; i <= n; i++ {
		sum += a      // Menambahkan bilangan Fibonacci ke-sum
		a, b = b, a+b // Memperbarui nilai a dan b
	}

	return sum // Mengembalikan hasil penjumlahan
}

func main() {
	fmt.Println(fibX(5))  // 12
	fmt.Println(fibX(10)) // 143
}
