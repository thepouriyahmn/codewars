package main

import "fmt"

func main() {

	a := []int{2, 2, 3}
	b := []int{4, 9, 9}
	fmt.Println(comp(a, b))
}
func comp(a []int, b []int) bool {
	var c int
	if a == nil || b == nil || len(a) != len(b) {
		return false

	}

	for i := 0; i < len(a); i++ {
		sum := a[i] * a[i]
		for j := 0; j < len(b); j++ {
			if sum == b[j] {
				b[j] = 0
				c++
				break
			}
		}
	}
	if c == len(a) {
		return true
	}
	return false
}
