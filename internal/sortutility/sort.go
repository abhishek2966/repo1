package sortutility

import (
	"examplerepo/pkg/sorting"
)

func SortedSlice(s []int) []int {
	sorting.SelectionSort(s)
	return s
}
