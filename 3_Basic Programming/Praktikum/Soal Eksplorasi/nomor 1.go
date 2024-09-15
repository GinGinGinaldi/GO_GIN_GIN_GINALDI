package main

import (
	"fmt"
	"sort"
	"strings"
)

// Fungsi untuk mengurutkan huruf dalam kata
func urutkanHuruf(kata string) string {
	huruf := strings.Split(kata, "") // Memecah kata menjadi huruf-huruf
	sort.Strings(huruf)              // Mengurutkan huruf
	return strings.Join(huruf, "")   // Menggabungkan kembali huruf yang sudah diurutkan
}

func main() {
	var kataPertama, kataKedua string

	// Meminta input dari pengguna
	fmt.Print("Masukkan kata pertama: ")
	fmt.Scan(&kataPertama)
	fmt.Print("Masukkan kata kedua: ")
	fmt.Scan(&kataKedua)

	// Mengubah huruf menjadi huruf kecil agar tidak case-sensitive
	kataPertama = strings.ToLower(kataPertama)
	kataKedua = strings.ToLower(kataKedua)

	// Memeriksa apakah kedua kata merupakan anagram
	if urutkanHuruf(kataPertama) == urutkanHuruf(kataKedua) {
		fmt.Println("Anagram")
	} else {
		fmt.Println("Bukan Anagram")
	}
}
