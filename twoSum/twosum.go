package main

import "fmt"

func main() {
	array := []int{2, 3, 4}
	targetNum := 5
	fmt.Println(twosum(array, targetNum))
}

func twosum(numbers []int, target int) []int {
	for i, value := range numbers {
		complimet := target - value
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] == complimet {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
