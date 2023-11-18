package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil {
		if elem < root.Data {
			BTreeSearchItem(root.Left, elem)
			root = root.Left.Parent
		} else {
			BTreeSearchItem(root.Right, elem)
			root = root.Right.Parent
		}
		return root
	}
	return nil
}
