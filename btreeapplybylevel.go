package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	height := BTreeLevelCount(root)
	if root!=nil{
	for i := 1; i <= height; i++ {
		f(root.Data)
		f(root.Left.Data)
		f(root.Right.Data)
	}
}
}
