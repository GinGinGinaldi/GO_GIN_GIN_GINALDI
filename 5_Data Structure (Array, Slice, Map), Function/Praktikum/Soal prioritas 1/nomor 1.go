package main

import "fmt"

func main() {
	hasil1 := gabungArray([][]int{
		{1, 2, 5, 4, 3},
		{1, 2, 7, 8},
	})
	fmt.Println(hasil1) // [1, 2, 5, 4, 3, 7, 8]

	hasil2 := gabungArray([][]int{
		{1, 2, 5, 4, 5},
		{7, 9, 10, 5},
	})
	fmt.Println(hasil2) // [1, 2, 5, 4, 7, 9, 10]

	hasil3 := gabungArray([][]int{
		{1, 4, 5},
		{7},
	})
	fmt.Println(hasil3) // [1, 4, 5, 7]
}

func gabungArray(data [][]int) []int {
	elemenUnik := make(map[int]bool)
	var hasil []int

	for _, baris := range data {
		for _, nilai := range baris {
			if !elemenUnik[nilai] {
				elemenUnik[nilai] = true
				hasil = append(hasil, nilai)
			}
		}
	}
	return hasil
}
