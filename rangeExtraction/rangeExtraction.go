package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}
	fmt.Println(extract(a))
}
func extract(list []int) string {
	var c int
	var temp int
	var j int

	save := []string{}

	for {

		j = temp

		if j == len(list) {
			break
		}
		for ; j < len(list)-1; j++ {

			if list[j+1] == list[j]+1 {
				c++
			} else {
				break
			}

		}
		if c >= 2 {
			save = append(save, strconv.Itoa(list[j-c])+"-"+strconv.Itoa(list[j]))

		} else if j != 0 {
			if list[j-1] == list[j]-1 {
				save = append(save, strconv.Itoa(list[j-1]))
				save = append(save, strconv.Itoa(list[j]))
			} else {
				save = append(save, strconv.Itoa(list[j]))
				fmt.Println(list[j])
			}
		} else if j == 0 {
			save = append(save, strconv.Itoa(list[j]))
		}

		temp = j + 1

		c = 0

	}
	fmt.Println(save)
	var sum string
	for i := 0; i < len(save); i++ {
		sum += save[i]
		if i != len(save)-1 {
			sum += ","
		}
	}

	return sum

}
