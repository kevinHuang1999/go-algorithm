package main

import (
	"algorithm/sorts"
	"fmt"
)

func main() {
	arr1 := []int{4, 2, 5, 1}
	arr1 = sorts.HeapSort(arr1)
	fmt.Println(arr1)

	arr2 := []int{4, 2, 5, 1}
	arr2 = sorts.InsertionSort(arr2)
	fmt.Println(arr2)

	arr3 := []int{4, 2, 5, 1, 9, 2, 3}
	arr3 = sorts.MergeSort(arr3)
	fmt.Println(arr3)

	arr4 := []int{4, 2, 5, 1, 9, 2, 3}
	arr4 = sorts.QuickSort(arr4)
	fmt.Println(arr4)
}
