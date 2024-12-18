package main

import "fmt"

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(arr []int) PosPeaks {
	posarr := []int{}
	peakarr := []int{}
	for i := 1; i < len(arr)-1; i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] && arr[i] > arr[i-1] {
				posarr = append(posarr, i)
				peakarr = append(peakarr, arr[i])
				break

			} else if arr[i] < arr[j] {
				break
			}
		}
	}

	return PosPeaks{
		Pos:   posarr,
		Peaks: peakarr,
	}

}

func main() {
	a := []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1}
	result := PickPeaks(a)
	fmt.Println(result)
}
