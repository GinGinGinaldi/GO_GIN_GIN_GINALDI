package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%4 == 0 && i%7 == 0 {
			fmt.Println("buzz")
		} else if i%4 == 0 {
			fmt.Println(i * i) // tampilin kuadrat
		} else if i%7 == 0 {
			fmt.Println(i * i * i) // tampilin pangkat tiga
		} else {
			fmt.Println(i) // tampilin angka aslinya
		}
	}
}
