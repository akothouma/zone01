package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	for root.Data != elem {
		if root != nil {
			if elem < root.Data {
				BTreeSearchItem(root.Left, elem)
				root = root.Left.Parent
			} else if elem > root.Data {
				BTreeSearchItem(root.Right, elem)
				root = root.Right.Parent
			}
			return root
		}
	}
	return nil
}
