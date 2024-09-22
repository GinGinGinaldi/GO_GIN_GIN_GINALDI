package main

import (
	"fmt"
	"math"
)

func cekBilanganPrima(angka int) bool {
	if angka < 2 {
		return false
	}
	batas := int(math.Sqrt(float64(angka)))
	for i := 2; i <= batas; i++ {
		if angka%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(cekBilanganPrima(1000000007)) // true
	fmt.Println(cekBilanganPrima(13))         // true
	fmt.Println(cekBilanganPrima(17))         // true
	fmt.Println(cekBilanganPrima(20))         // false
	fmt.Println(cekBilanganPrima(35))         // false
}
