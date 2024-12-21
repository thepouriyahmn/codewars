package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(weightOrder("56 65 74 100 99 68 86 180 90"))
}
func weightOrder(list string) string {
	var sum int
	var sum2 string
	var str string
	arr := []int{}
	arrstr := []string{}
	for _, v := range list {
		if v != ' ' {
			num, err := strconv.Atoi(string(v))
			sum += num
			sum2 += string(v)
			if err != nil {
				fmt.Println(err)

			}
		} else {
			arr = append(arr, sum)
			arrstr = append(arrstr, sum2)
			sum = 0
			sum2 = ""
		}

	}
	arr = append(arr, sum)
	arrstr = append(arrstr, sum2)

	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] || (arr[j] == arr[j+1] && arrstr[j] > arrstr[j+1]) {

				arr[j], arr[j+1] = arr[j+1], arr[j]

				arrstr[j], arrstr[j+1] = arrstr[j+1], arrstr[j]

			}
		}
	}
	for i := 0; i < len(arrstr); i++ {
		str += arrstr[i]
		if i != len(arrstr)-1 {
			str += " "
		}
	}
	return str
}
