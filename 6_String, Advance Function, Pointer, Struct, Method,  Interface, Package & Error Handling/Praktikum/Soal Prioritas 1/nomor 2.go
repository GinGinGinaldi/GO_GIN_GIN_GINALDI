package main

import (
	"fmt"
)

func main() {
	fmt.Println(spinString("alta"))    // aalt
	fmt.Println(spinString("alterra")) // aalrtre
	fmt.Println(spinString("sepulsa")) // saesplu
}

func spinString(sentence string) string {
	n := len(sentence)
	if n == 0 {
		return ""
	}

	rotated := make([]byte, n)
	for i := 0; i < n; i++ {
		rotated[i] = sentence[(i+2)%n]
	}

	return string(rotated)
}
