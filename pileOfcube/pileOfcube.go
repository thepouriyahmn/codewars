package main

import (
	"fmt"
)

func main() {
	fmt.Println(listSquared(42, 250))

}
func listSquared(m int, n int) [][]int {

	var arr [][]int = [][]int{}

	for i := m; i <= n; i++ {

		if isSquare(dividors(i)) {

			arr = append(arr, []int{i, dividors(i)})
		}

	}
	return arr
}
func dividors(num int) int {

	var l int
	var temp int
	j := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			l++
		}
	}
	arr := make([]int, l)
	for i := num; i >= 1; i-- {
		if num%i == 0 {

			arr[j] = i * i
			j++
		}
	}
	for _, v := range arr {
		temp += v
	}

	return temp
}
func isSquare(num int) bool {
	for i := 1; i*i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
