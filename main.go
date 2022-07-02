package main

import (
	"fmt"

	"github.com/abhishek2966/repo1/pkg/bst"
)

func main() {
	s3 := []int{15, 14, 12, 23}
	t := bst.CreateBST(s3...)
	fmt.Println(bst.Search(t, 24))
}
