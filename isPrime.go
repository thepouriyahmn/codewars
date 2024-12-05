package main

import "fmt"

func main() {

	fmt.Println(isPrime(0))
	fmt.Println(fact(5))

}
func isPrime(num int) bool {
	prime := true
	if num < 2 {
		return false

	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			prime = false
			break
		}
	}
	return prime
}
func fact(num int) int {
	res := 1
	for i := num; i > 0; i-- {
		res = res * i
	}
	return res
}
