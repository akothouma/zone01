package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	height := BTreeLevelCount(root)
	if height > 1 {
		f(root.Data)
		f(root.Left.Data)
		f(root.Right.Data)
	}
}
