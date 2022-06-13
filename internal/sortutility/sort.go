package sortutility

import (
	"repo1/pkg/sorting"
)

func SortedSlice(s []int) []int {
	sorting.SelectionSort(s)
	return s
}
