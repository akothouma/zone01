package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil {
		if elem < root.Data && root.Data != elem {
			BTreeSearchItem(root.Left, elem)
			if root.Data == elem {
				return root
			}
		} else if elem > root.Data && root.Data != elem {
			BTreeSearchItem(root.Right, elem)
			if root.Data == elem {
				return root
			}
		}
	}
	return nil
}
