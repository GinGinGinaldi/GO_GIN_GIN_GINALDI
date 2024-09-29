package main

import (
	"fmt"
)

// SimpleEquations adalah fungsi yang mencari tiga bilangan x, y, dan z
// yang memenuhi tiga persamaan:
// x + y + z = A
// x * y * z = B
// x^2 + y^2 + z^2 = C
func SimpleEquations(a, b, c int) {
	// Inisialisasi loop untuk mencari nilai x, y, dan z
	for x := 1; x < a; x++ {
		for y := 1; y < a; y++ {
			for z := 1; z < a; z++ {
				// Memastikan x, y, dan z adalah bilangan yang berbeda
				if x != y && y != z && x != z {
					// Memeriksa apakah nilai x, y, dan z memenuhi ketiga persamaan
					if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
						fmt.Println(x, y, z)
						return
					}
				}
			}
		}
	}
	// Jika tidak ada solusi, cetak "No solution"
	fmt.Println("No solution.")
}

func main() {
	SimpleEquations(1, 2, 3)  // Output: No solution
	SimpleEquations(6, 6, 14) // Output: 1 2 3
}
