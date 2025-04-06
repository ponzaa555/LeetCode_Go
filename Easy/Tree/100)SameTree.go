package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// break state
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	checkLeft := isSameTree(p.Left, q.Left)
	checkRight := isSameTree(p.Right, q.Right)

	return checkLeft && checkRight
}
