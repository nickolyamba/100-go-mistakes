package main

import "fmt"

func main() {
	s := make([]int, 3, 6)
	print("s", s)

	s[1] = 1
	print("s", s)

	s = append(s, 2)
	print("s", s)

	s = append(s, 3, 4, 5)
	print("s", s)

	s1 := make([]int, 3, 6)
	s2 := s1[1:3]
	s1[1] = 1
	print("s1", s1)
	print("s2", s2)

	s2 = append(s2, 2)
	print("s1", s1)
	print("s2", s2)

	s2 = append(s2, 3)
	s2 = append(s2, 4)
	s2 = append(s2, 5)
	print("s1", s1)
	print("s2", s2)
}

func print(name string, s []int) {
	fmt.Printf("slice: %s, slice start addr: %p, underlying array addr: %p, len=%d, cap=%d: %v\n",
		name, s, &s[0], len(s), cap(s), s)
}
