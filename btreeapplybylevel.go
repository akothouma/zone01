package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyByLevel(root.Left, f)
		f(root.Data)
		BTreeApplyByLevel(root.Right, f)
	}
}

