package piscine

func BTreeLevelCount(root *TreeNode) int {
	var lcount, rcount int
	if root == nil {
		return 0
	} else {
		lcount = BTreeLevelCount(root.Left)
		rcount = BTreeLevelCount(root.Right)

		if lcount > rcount {
			return lcount + 1
		}
		return rcount + 1
	}
}
