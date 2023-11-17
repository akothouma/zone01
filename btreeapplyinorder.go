package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		f(&root.Left)
		f(&root)
		f(&root.Right)
	}
}
