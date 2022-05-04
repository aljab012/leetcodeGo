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

function hasPathSum(root: TreeNode | null, targetSum: number): boolean {
  return helper(root, targetSum)
}

function helper(root: TreeNode, targetSum: number): boolean {
  if (root == null) return false
  let sum = targetSum - root.val
  if (sum == 0 && root.left == null && root.right == null) return true
  return helper(root.left, sum) || helper(root.right, sum)
}
