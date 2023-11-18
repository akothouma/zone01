package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil {
		if elem < root.Data && root.Data != elem {
			BTreeSearchItem(root.Left, elem)
			root = root.Left
		} else if elem > root.Data && root.Data != elem {
			BTreeSearchItem(root.Right, elem)
			root = root.Right
		}
		return root
	}
	return nil
}
