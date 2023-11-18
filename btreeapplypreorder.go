package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root,f)
		BTreeApplyPreorder(root.Left, f)
		BTreeApplyPreorder(root.Right,f)
	}
}
