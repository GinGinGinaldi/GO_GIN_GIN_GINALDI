package main

import (
	"fmt"
)

func angkaKeRomawi(angka int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	simbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	hasil := ""
	for i := 0; i < len(val); i++ {
		for angka >= val[i] {
			hasil += simbol[i]
			angka -= val[i]
		}
	}
	return hasil
}
func main() {
	input := []int{4, 9, 23, 2021, 1646}
	for _, angka := range input {
		fmt.Printf("Angka: %d -> Romawi: %s\n", angka, angkaKeRomawi(angka))
	}
}
