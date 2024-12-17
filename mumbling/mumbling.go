package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mumbling("abcdhsjso"))
}
func mumbling(a string) string {
	var c string
	var b string

	s := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		c += strings.ToUpper(string(a[i]))

	}
	for i := 0; i < len(c); i++ {
		s[i] += string(c[i])
	}
	for i := 0; i < len(a); i++ {

		for j := 1; j <= i; j++ {

			s[i] += strings.ToLower(string(a[i]))

		}

	}

	for i, v := range s {
		b += v
		if i != len(s)-1 {
			b += "-"
		}
	}

	return b
}
