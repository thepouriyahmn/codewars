package main

import "fmt"

func main() {

	fmt.Println(isPrime(0))

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
