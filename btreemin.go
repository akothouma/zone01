package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root != nil {
		for root.Left != nil {
			BTreeMin(root.Left)
			root = root.Left
		}
		return root
	}
	return nil
}
