package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	for root != nil {
		f(root.Left)
		f(root.Parent)
		f(root.Right)
	}
}
