package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	for root != nil {
		f(root.Left.Data, root.Data, root.Right.Data)
	}
}
