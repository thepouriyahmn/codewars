package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(SumOfDevided([]int{15, 21, 24, 30, 45}))
}

func SumOfDevided(list []int) string {
	tempPrimes := []int{}
	prime := []int{}

	var save int
	for i := 0; i < len(list); i++ {

		tempPrimes = append(tempPrimes, dividedPrime(list[i])...)
	}

	for i := 0; i < len(tempPrimes); i++ {
		for j := 0; j < len(tempPrimes); j++ {
			if tempPrimes[i] == tempPrimes[j] && i != j {

				save = j
				tempPrimes[i] = 0
				break
			}
			save = j
		}

		if save == len(tempPrimes)-1 && tempPrimes[i] != 0 {
			prime = append(prime, tempPrimes[i])
		}

	}

	for i := 1; i < len(prime); i++ {
		for j := 0; j < len(prime)-i; j++ {
			if prime[j] > prime[j+1] {
				prime[j], prime[j+1] = prime[j+1], prime[j]
			}
		}
	}

	var sum int
	var res string
	for i := 0; i < len(prime); i++ {
		for j := 0; j < len(list); j++ {

			if list[j]%prime[i] == 0 {

				sum += list[j]

			}

		}

		res += "(" + strconv.Itoa(prime[i]) + " " + strconv.Itoa(sum) + ")"
		sum = 0
	}

	return res

}
func findPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func dividedPrime(num int) []int {

	if num < 0 {
		num = -num
	}
	arr := []int{}
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			if findPrime(i) {
				arr = append(arr, i)
			}
		}
	}
	if len(arr) == 0 {
		return nil
	}
	return arr
}
