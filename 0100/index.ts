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

function isSameTree(p: TreeNode | null, q: TreeNode | null): boolean {
  return helper(p, q)
}
function helper(p: TreeNode, q: TreeNode): boolean {
  if (p == null && q == null) return true

  if (p == null && q != null && q == null && p != null) {
    return false
  }
  return p?.val == q?.val && helper(p.left, q.left) && helper(p.right, q.right)
}
