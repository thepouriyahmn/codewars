package main

import (
	"fmt"
	"strconv"
)

func main() {
	println(validISBN10("048665088x"))
}
func validISBN10(isbn string) bool {
	var sum int
	var c int
	var d int
	n := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	nx := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "X", "x"}
	if len(isbn) != 10 {
		return false
	}
	for i := 0; i < len(isbn)-1; i++ {
		for j := 0; j < len(n); j++ {
			if string(isbn[i]) == n[j] {
				c++
			}
		}
	}
	if c != 9 {
		return false
	}
	for i := 0; i < len(nx); i++ {
		if string(isbn[9]) == nx[i] {
			d++
			break
		}
	}
	if d != 1 {
		return false
	}

	if isbn[9] != 'X' && isbn[9] != 'x' {
		for i := 0; i < len(isbn); i++ {
			num, err := strconv.Atoi(string(isbn[i]))
			if err != nil {
				fmt.Println(err)

			}
			sum += num * (i + 1)
		}

	} else if isbn[9] == 'X' || isbn[9] == 'x' {
		for i := 0; i < len(isbn)-1; i++ {
			num, err := strconv.Atoi(string(isbn[i]))
			if err != nil {
				fmt.Println(err)

			}
			sum += num * (i + 1)
		}
		sum += 10 * 10
	}
	if sum%11 == 0 {
		return true
	}

	return false

}
