package main

import (
	"fmt"
)

func merge(x []int, l, mid, r int) {
	temp := make([]int, r-l+1)
	var i, j, k int
	for i, j, k = l, mid+1, 0; i <= mid && j <= r; k++ {
		if x[i] < x[j] {
			temp[k] = x[i]
			i++
		} else {
			temp[k] = x[j]
			j++
		}
	}
	if j > r {

		for ; i <= mid; i, k = i+1, k+1 {
			temp[k] = x[i]
		}
	} else {

		for ; j <= r; j, k = j+1, k+1 {
			temp[k] = x[j]
		}
	}

	for i, j = 0, l; i < len(temp); i, j = i+1, j+1 {
		x[j] = temp[i]
	}
}
func mergeSort(x []int, l, r int) {

	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergeSort(x, l, mid)
	mergeSort(x, mid+1, r)
	merge(x, l, mid, r)
}
func main() {
	x := []int{7, 0, 1, 4, 2}
	fmt.Println("Input: ", x)
	mergeSort(x, 0, len(x)-1)
	fmt.Println("Sorted: ", x)
}
