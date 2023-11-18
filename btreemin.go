package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root != nil {
		BTreeMin(root.Left)
		root = root.Left
		return root
	}
	return nil
}
