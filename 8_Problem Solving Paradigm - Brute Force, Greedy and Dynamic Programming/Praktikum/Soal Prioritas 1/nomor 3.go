package main

import "fmt"

var memo = map[int]int{}

func fibonacci(number int) int {
	if val, ok := memo[number]; ok {
		return val
	}
	if number == 0 || number == 1 {
		return number
	}
	memo[number] = fibonacci(number-1) + fibonacci(number-2)
	return memo[number]
}
func main() {
	fmt.Println(fibonacci(0))  // Output: 0
	fmt.Println(fibonacci(2))  // Output: 1
	fmt.Println(fibonacci(9))  // Output: 34
	fmt.Println(fibonacci(10)) // Output: 55
	fmt.Println(fibonacci(12)) // Output: 144
}
