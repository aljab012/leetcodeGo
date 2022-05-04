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

function inorderTraversal(root: TreeNode | null): number[] {
  let arr = []

  arr = helper(root, arr)
  return arr
}
function helper(root: TreeNode, arr: number[]): number[] {
  if (root == null) {
    return arr
  }
  if (root.left) {
    arr = helper(root.left, arr)
  }
  arr.push(root.val)
  if (root.right) {
    arr = helper(root.right, arr)
  }
  return arr
}
