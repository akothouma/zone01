package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root.Parent == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Parent.Data {
		root.Left = BTreeInsertData(root.Parent.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Parent.Right, data)
	}
	return root.Parent
}
