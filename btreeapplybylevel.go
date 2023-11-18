package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	height := BTreeLevelCount(root)
	if height > 1 {
		f(root.Data)
		BTreeApplyByLevel(root.Left, f)
		BTreeApplyByLevel(root.Right, f)
	}
}
