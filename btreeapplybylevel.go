package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	var height int
	if root != nil {
		height = BTreeLevelCount(root)

		for i := 1; i <= height; i++ {
			BTreeApplyByLevel(root, f)
		}
	}
}
