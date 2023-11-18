package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	for root.Data != elem {
		if root != nil {
			if elem < root.Data {
				root = root.Left.Parent
				BTreeSearchItem(root.Left, elem)
			} else if elem > root.Data {
				root = root.Right.Parent
				BTreeSearchItem(root.Right, elem)
			}

		}
	}
	if root == nil {
		return nil
	}
	return root
}
