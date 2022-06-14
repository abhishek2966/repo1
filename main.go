package main

import (
	"fmt"

	"github.com/abhishek2966/repo1/pkg/sorting"
)

func main() {
	s1 := []int{5}
	sorting.SelectionSort(s1)
	fmt.Println(s1)
	s2 := []int{55}
	sorting.BubbleSort(s2)
	fmt.Println(s2)
	s3 := []int{15, 14, 12, 13}
	s3 = sorting.InsertionSort(s3)
	fmt.Println(s3)
}
