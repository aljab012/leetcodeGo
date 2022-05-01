function searchInsert(nums: number[], target: number): number {
  let p1 = 0
  let p2 = nums.length - 1

  while (p1 <= p2) {
    let mid = Math.floor((p1 + p2) / 2)
    if (nums[mid] == target) {
      return mid
    } else if (nums[mid] < target) {
      p1 = mid + 1
    } else {
      p2 = mid - 1
    }
  }
  return p1
}
