package main

import "fmt"

func main() {
	a := []int{1, 2, 5, 0, 6, 0, 4}
	fmt.Println(moveZeros(a))
}
func moveZeros(arr []int) []int {
	arr2 := make([]int, len(arr))

	var j int

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr2[j] = arr[i]
			j++
		}
	}
	return arr2
}
