package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil || root.Data != elem {
		if elem < root.Data {
			root.Left = root
			BTreeSearchItem(root.Left, elem)
		} else if elem > root.Data {
			root.Right = root
			BTreeSearchItem(root.Right, elem)
		}
		return root
	}
	return nil
}
