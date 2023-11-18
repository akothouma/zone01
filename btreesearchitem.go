package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil {
		if elem < root.Data && root.Left != nil {
			root.Left = BTreeSearchItem(root.Left, elem)
			root = root.Left.Parent
		} else if elem > root.Data && root.Right != nil {
			root.Right = BTreeSearchItem(root.Right, elem)
			root = root.Right.Parent
		}
		return root
	}
	return nil
}
