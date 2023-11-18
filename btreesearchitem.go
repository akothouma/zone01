package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil || root.Data != elem {
		if elem < root.Data {
			BTreeSearchItem(root.Left, elem)
			root.Left.Parent=root
		} else if elem > root.Data {
			BTreeSearchItem(root.Right, elem)
			root.Right.Parent=root
		}
		return root
	}
	return nil
}
