package main

import (
	"fmt"
)

func main() {
	data := []int{5, 2, 8, 3, 9, 1, 6, 7, 4, 0}
	fmt.Println("Unsorted data:", data)
	sorted := MergeSort(data)
	fmt.Println("Sorted data:", sorted)
}
