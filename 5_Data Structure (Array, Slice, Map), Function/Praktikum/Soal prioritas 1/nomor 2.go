package main

import "fmt"

func main() {
	fmt.Println(jumlahBilPrima([]int{1, 2, 4, 5, 12, 19, 30})) // 26
	fmt.Println(jumlahBilPrima([]int{45, 17, 44, 67, 11}))     // 95
	fmt.Println(jumlahBilPrima([]int{15, 12, 9, 10}))          // 0
}

func cekPrima(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func jumlahBilPrima(angka []int) int {
	total := 0
	for _, bilangan := range angka {
		if cekPrima(bilangan) {
			total += bilangan
		}
	}
	return total
}
