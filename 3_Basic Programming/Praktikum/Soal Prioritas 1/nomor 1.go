package main

import (
	"fmt"
)

func main() {
	var jariJari, tinggi, volume float64
	const phi = 3.14

	// minta input jari-jari dan tinggi
	fmt.Print("Masukkan jari-jari tabung: ")
	fmt.Scan(&jariJari)
	fmt.Print("Masukkan tinggi tabung: ")
	fmt.Scan(&tinggi)

	// hitungin volume tabung
	volume = phi * jariJari * jariJari * tinggi

	// tampilin hasil volume tabung
	fmt.Printf("Volume tabung adalah: %.2f\n", volume)
}
