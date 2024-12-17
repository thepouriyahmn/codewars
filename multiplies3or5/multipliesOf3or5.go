package main

import "fmt"

func main() {
	fmt.Println(multiply3or5(10))
}
func multiply3or5(number int) int {
	var result int
	for i := 1; i < number; i++ {
		if i%5 == 0 || i%3 == 0 {
			result = result + i
		}
	}
	return result
}
