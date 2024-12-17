package main

import "fmt"

func main() {
	fmt.Print(deadfish("zizlwisipigelkgiibsghis"))
}
func deadfish(s string) []int {
	r := 0
	a := []int{}

	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "i":
			r++
			break
		case "s":
			r = r * r
			break
		case "d":
			r--
			break
		case "o":
			a = append(a, r)
			break

		}

	}

	return a
}
