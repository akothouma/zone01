package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root != nil {
		for root.Right != nil {
			root = root.Right
		}
		return root
	}
	return nil
}
