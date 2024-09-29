package main

import (
	"fmt"
	"math"
)

// Fungsi Frog untuk menghitung biaya minimum yang diperlukan katak untuk mencapai Batu N.
func Frog(jumps []int) int {
	// Inisialisasi array untuk menyimpan biaya minimum setiap batu
	n := len(jumps)
	if n == 0 {
		return 0
	}

	// Array untuk menyimpan total biaya minimum hingga batu i
	biaya := make([]int, n)
	biaya[0] = 0 // Biaya awal di Batu 1 adalah 0

	// Iterasi untuk menghitung biaya minimum dari Batu 2 hingga Batu N
	for i := 1; i < n; i++ {
		// Inisialisasi biaya minimum ke batu saat ini dengan nilai tak terhingga
		biayaMin := math.MaxInt32

		// Lompat dari batu sebelumnya (Batu i-1) ke batu saat ini (Batu i)
		biayaDariBatuSebelumnya := biaya[i-1] + int(math.Abs(float64(jumps[i]-jumps[i-1])))
		biayaMin = min(biayaMin, biayaDariBatuSebelumnya)

		// Jika batu ini bisa dijangkau dari dua batu sebelumnya (Batu i-2)
		if i > 1 {
			biayaDariDuaBatuSebelumnya := biaya[i-2] + int(math.Abs(float64(jumps[i]-jumps[i-2])))
			biayaMin = min(biayaMin, biayaDariDuaBatuSebelumnya)
		}

		// Update biaya minimum ke batu ini
		biaya[i] = biayaMin
	}

	// Biaya total minimum untuk mencapai Batu N
	return biaya[n-1]
}

// Fungsi min untuk mendapatkan nilai minimum dari dua angka
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Uji fungsi Frog dengan beberapa contoh kasus
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // Output: 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // Output: 40
}
