package main

import "fmt"

func hitungPangkat(basis, eksponen int) int {
	if eksponen == 0 {
		return 1
	}
	setengahPangkat := hitungPangkat(basis, eksponen/2)
	if eksponen%2 == 0 {
		return setengahPangkat * setengahPangkat
	} else {
		return basis * setengahPangkat * setengahPangkat
	}
}

func main() {
	fmt.Println(hitungPangkat(2, 3))  // 8
	fmt.Println(hitungPangkat(5, 3))  // 125
	fmt.Println(hitungPangkat(10, 2)) // 100
	fmt.Println(hitungPangkat(2, 5))  // 32
	fmt.Println(hitungPangkat(7, 3))  // 343
}
