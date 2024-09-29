package main

import (
	"fmt"
)

// Fungsi untuk mengkonversi angka biasa ke Angka Romawi
func angkaKeRomawi(angka int) string {
	// Pemetaan nilai integer ke simbol Romawi
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	simbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	hasil := "" // Inisialisasi string hasil untuk menyimpan Angka Romawi

	// Perulangan untuk memeriksa setiap nilai dan simbol Romawi
	for i := 0; i < len(val); i++ {
		// Selama nilai angka lebih besar atau sama dengan nilai Romawi saat ini
		for angka >= val[i] {
			hasil += simbol[i] // Tambahkan simbol Romawi ke hasil
			angka -= val[i]    // Kurangi angka dengan nilai Romawi yang telah ditambahkan
		}
	}
	return hasil // Kembalikan hasil konversi
}

func main() {
	// Contoh penggunaan program dengan beberapa input
	input := []int{4, 9, 23, 2021, 1646}
	for _, angka := range input {
		fmt.Printf("Angka: %d -> Romawi: %s\n", angka, angkaKeRomawi(angka))
	}
}
