package sortutility

import (
	"github.com/abhishek2966/repo1/pkg/sorting"
)

func SortedSlice(s []int) []int {
	sorting.SelectionSort(s)
	return s
}
