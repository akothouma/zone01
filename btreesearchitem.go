package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil || root.Data != elem {
		if elem < root.Data {
			BTreeSearchItem(root.Left, elem)
			if root.Data == elem {
				return root
			}
		} else if elem > root.Data {
			BTreeSearchItem(root.Right, elem)
			if root.Data == elem {
				return root
			}
		}
	}
	return nil
}
