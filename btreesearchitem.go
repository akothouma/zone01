package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil && root.Data != elem {
		if elem < root.Data {
			BTreeSearchItem(root.Left, elem)
			root = root.Left
			BTreeSearchItem(root.Left, elem)
		} else if elem > root.Data {
			BTreeSearchItem(root.Right, elem)
			root = root.Right
			BTreeSearchItem(root.Right, elem)
		}
		return root
	}
	return nil
}
