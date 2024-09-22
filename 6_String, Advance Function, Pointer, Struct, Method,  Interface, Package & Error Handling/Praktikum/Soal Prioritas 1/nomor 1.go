package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pickVocals("alterra academy"))     // aea aae
	fmt.Println(pickVocals("sepulsa mantap jiwa")) // eua aa ia
	fmt.Println(pickVocals("kopi susu"))           // oi uu
}

func pickVocals(sentence string) string {
	vowels := "aeiou"
	words := strings.Split(sentence, " ")
	result := ""

	for _, word := range words {
		vocalWord := ""
		for _, char := range word {
			if strings.Contains(vowels, string(char)) {
				vocalWord += string(char)
			}
		}
		result += vocalWord + " "
	}

	return strings.TrimSpace(result)
}
