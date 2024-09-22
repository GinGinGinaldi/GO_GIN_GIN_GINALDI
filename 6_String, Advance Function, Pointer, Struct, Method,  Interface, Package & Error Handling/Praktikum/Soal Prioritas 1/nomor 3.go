package main

import (
	"fmt"
)

func main() {
	fmt.Println(groupPalindrome([]string{"katak", "civic", "kasur", "rusak"}))                  // [[katak, civic], kasur, rusak]
	fmt.Println(groupPalindrome([]string{"racecar", "seru", "kasur", "civic", "bilik", "kak"})) // [[racecar, civic, kak], seru, kasur, bilik]
	fmt.Println(groupPalindrome([]string{"masuk", "civic", "hahah", "garam"}))                  // [[civic, hahah], masuk, garam]
}

func isPalindrome(word string) bool {
	n := len(word)
	for i := 0; i < n/2; i++ {
		if word[i] != word[n-i-1] {
			return false
		}
	}
	return true
}

func groupPalindrome(words []string) []interface{} {
	palindromes := []string{}
	nonPalindromes := []string{}

	for _, word := range words {
		if isPalindrome(word) {
			palindromes = append(palindromes, word)
		} else {
			nonPalindromes = append(nonPalindromes, word)
		}
	}

	result := []interface{}{palindromes}
	for _, nonPalindrome := range nonPalindromes {
		result = append(result, nonPalindrome)
	}
	return result
}
