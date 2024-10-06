package main

import "fmt"

type set struct {
	dt map[int]int
}

func (s set) add(i int) {
	s.dt[i] = 1
}

func (s set) get() []int {
	r := []int{}

	for k := range s.dt {
		r = append(r, k)
	}

	return r
}

func (s set) remove(k int) {
	delete(s.dt, k)
}

func main() {
	s := set{
		dt: map[int]int{},
	}

	s.add(1)
	s.add(2)
	s.add(1)
	s.add(3)
	s.add(4)
	s.add(5)

	s.remove(100)

	fmt.Println(s.get())
}
