package main

import (
	"fmt"

	"github.com/abhishek2966/repo1/internal/sortutility"
)

func main() {
	fmt.Println("from example repo")
	s3 := []int{5, 4, 2, 3}
	//sorting.SelectionSort(s3)
	s := sortutility.SortedSlice(s3)
	fmt.Println(s)
}
