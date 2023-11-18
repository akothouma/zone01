package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil || root.Data != elem {
		if elem < root.Data {
			root.Left.Parent = root
			BTreeSearchItem(root.Left.Parent, elem)
		} else if elem > root.Data {
			root.Right.Parent = root
			BTreeSearchItem(root.Right.Parent, elem)
		}
		return root
	}
	return nil
}
