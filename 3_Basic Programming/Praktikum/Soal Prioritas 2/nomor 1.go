package main

import (
	"fmt"
)

func main() {
	var angka int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	// cetakin pola berdasarkan faktor bilangan
	for i := 1; i <= angka; i++ {
		if angka%i == 0 { // Jika i adalah faktor dari angka
			if i%2 == 0 { // Jika faktor adalah bilangan genap
				fmt.Print("I")
			} else { // Jika faktor adalah bilangan ganjil
				fmt.Print("O")
			}
		}
	}
	fmt.Println() // Pindah baris setelah pola udah dicetak
}
