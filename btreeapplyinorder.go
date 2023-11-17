package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
			BTreeApplyInorder(root.Left,f)
			BTreeApplyInorder(root.Parent,f)
			BTreeApplyInorder(root.Right,f)
		}
	}
