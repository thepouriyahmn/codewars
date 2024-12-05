package main

import "fmt"

func main() {
	fmt.Print(fact(5))
}
func fact(num int) int {
	res := 1
	for i := num; i > 0; i-- {
		res = res * i
	}
	return res
}
