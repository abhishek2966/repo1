package sorting

// selection sort, time O(n^2), space O(1)
// find the min element from unsorted part and put it at the beginning
func SelectionSort(s []int) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s); i++ {
		minVal, minIndex := s[i], i
		for j := i; j < len(s); j++ {
			//find min from unsorted array
			if s[j] < minVal {
				minVal, minIndex = s[j], j
			}
		}
		if minIndex != i {
			swap(s, i, minIndex)
		}
	}
	return
}

// bubble sort, time O(n^2), space O(1)
// repeatedly swap adjacent elements if they are in wrong order and put max elem at right end
func BubbleSort(s []int) {
	if len(s) == 0 {
		return
	}
	n := len(s)
	for n > 1 {
		for i := 1; i < n; i++ {
			if s[i-1] > s[i] {
				swap(s, i-1, i)
			}
		}
		n--
	}
}

func BubbleSortRecursive(s []int) {
	recursiveBubble(s, len(s))
}
func recursiveBubble(s []int, sizeSubArray int) {
	if sizeSubArray <= 1 {
		return
	}
	for i := 1; i < sizeSubArray; i++ {
		if s[i-1] > s[i] {
			swap(s, i-1, i)
		}
	}
	recursiveBubble(s, sizeSubArray-1)
}

// insertion sort, time O(n^2), space O(1)
// the way we sort playing cards; pick arr[i] and insert into sorted arr[0...i-1]
func InsertionSort(s []int) {
	for i := 1; i < len(s); i++ {
		j := i
		for j > 0 {
	    		if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
	    		}
	    		j = j - 1
		}
	}
}

func swap(s []int, i, j int) {
	if len(s) == 0 || len(s) <= i || len(s) <= j {
		return
	}
	s[i], s[j] = s[j], s[i]
}
