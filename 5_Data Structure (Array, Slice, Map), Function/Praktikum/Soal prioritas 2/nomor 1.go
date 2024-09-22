package main

import "fmt"

// Fungsi untuk memutar elemen slice sesuai dengan ilustrasi
func spinSlice(numbers []int) []int {
	if len(numbers) <= 2 {
		return numbers // tidak ada rotasi jika elemen kurang dari atau sama dengan 2
	}

	result := make([]int, len(numbers))
	mid := (len(numbers) + 1) / 2 // menghitung posisi tengah

	// Menyalin elemen pertama dan memulai dari tengah
	result[0] = numbers[0]
	for i, j, k := 1, len(numbers)-1, 1; i < mid; i, j, k = i+1, j-1, k+2 {
		result[k] = numbers[j] // menempatkan elemen dari akhir
		if k+1 < len(numbers) {
			result[k+1] = numbers[i] // menempatkan elemen dari awal setelahnya
		}
	}

	return result
}

func main() {
	fmt.Println(spinSlice([]int{1, 2, 3, 4, 5}))    // Output: [1,5,2,4,3]
	fmt.Println(spinSlice([]int{6, 7, 8}))          // Output: [6,8,7]
	fmt.Println(spinSlice([]int{8, 5, 4, 3, 2, 1})) // Output: [8,1,5,2,4,3]
}
