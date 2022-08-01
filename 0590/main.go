/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	return helper(root, []int{})
}

func helper(root *Node, ret []int) []int {
	if root == nil {
		return ret
	}
	for i := range root.Children {
		ret = helper(root.Children[i], ret)
	}
	ret = append(ret, root.Val)
	return ret
}
