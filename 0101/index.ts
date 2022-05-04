class TreeNode {
  val: number
  left: TreeNode | null
  right: TreeNode | null
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = val === undefined ? 0 : val
    this.left = left === undefined ? null : left
    this.right = right === undefined ? null : right
  }
}

function isSymmetric(root: TreeNode | null): boolean {
  return helper(root, root)
}
function helper(left: TreeNode, right: TreeNode) {
  if (left == null && right == null) {
    return true
  }
  if ((left == null && right != null) || (right == null && left != null)) {
    return false
  }
  return (
    left?.val == right?.val &&
    helper(left.left, right.right) &&
    helper(left.right, right.left)
  )
}
