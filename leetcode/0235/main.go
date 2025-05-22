package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * Recursive solution
 */
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor1(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor1(root.Right, p, q)
	} else {
		return root
	}
}

/*
 * Iterative solution
 *
 */
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}
	return nil
}
