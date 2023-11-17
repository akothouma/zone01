package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		_,err:=f(root.Left.Data,root.Parent.Data,root.Right.Data)

		if err!=nil{
			return
		}else{
			f(root.Left.Data)
			f(root.Parent.Data)
			f(root.Right.Data)
		}
		
	}
}
