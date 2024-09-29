package main

import (
	"fmt"
)

func buatSegitigaPascal(jumlahBaris int) [][]int {
	segitiga := make([][]int, jumlahBaris)
	for i := 0; i < jumlahBaris; i++ {
		segitiga[i] = make([]int, i+1)
		segitiga[i][0], segitiga[i][i] = 1, 1

		for j := 1; j < i; j++ {
			segitiga[i][j] = segitiga[i-1][j-1] + segitiga[i-1][j]
		}
	}
	return segitiga
}
func main() {
	jumlahBaris := 5
	hasil := buatSegitigaPascal(jumlahBaris)
	for _, baris := range hasil {
		fmt.Println(baris)
	}
}
