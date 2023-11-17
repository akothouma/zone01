package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		f(root.Left.Data)
		f(root.Parent.Data)
		f(root.Right.Data)
	}
}
