package bst

type tree struct {
	data                      int
	leftSubTree, rightSubTree *tree
}

func (t *tree) insert(v int) *tree {
	if t == nil {
		t = &tree{data: v}
		return t
	}
	if v <= t.data {
		if t.leftSubTree == nil {
			t.leftSubTree = &tree{data: v}
			return t
		}
		t.leftSubTree = t.leftSubTree.insert(v)
		return t
	}
	if t.rightSubTree == nil {
		t.rightSubTree = &tree{data: v}
		return t
	}
	t.rightSubTree = t.rightSubTree.insert(v)
	return t
}

func Search(t *tree, val int) (present bool) {
	if t == nil {
		return
	}
	if val == t.data {
		return true
	}
	if val < t.data {
		return Search(t.leftSubTree, val)
	}
	if val > t.data {
		return Search(t.rightSubTree, val)
	}
	return
}

func (t *tree) traversalInOrder() (result []int) {
	if t == nil {
		return
	}
	result = append(result, t.leftSubTree.traversalInOrder()...)
	result = append(result, t.data)
	result = append(result, t.rightSubTree.traversalInOrder()...)
	return
}

func CreateBST(values ...int) *tree {
	var t *tree
	for _, val := range values {
		t = t.insert(val)
	}
	return t
}

func InsertInBST(t *tree, val int) *tree {
	t = t.insert(val)
	return t
}

func GetSortedArray(t *tree) []int {
	return t.traversalInOrder()
}
