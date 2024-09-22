package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(item int) {
	if !s.contains(item) {
		s.items = append(s.items, item)
	}
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s Stack) values() []int {
	return s.items
}

func (s Stack) contains(item int) bool {
	for _, v := range s.items {
		if v == item {
			return true
		}
	}
	return false
}

func main() {
	stack := &Stack{}

	stack.push(10)
	stack.push(20)
	stack.push(10) // duplicate, won't be added

	fmt.Println("Stack values:", stack.values()) // [10, 20]

	stack.pop()
	fmt.Println("Stack after pop:", stack.values()) // [10]
}
