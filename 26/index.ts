function removeDuplicates(nums: number[]): number {
  let index = 0

  for (let i = 1; i < nums.length; i++) {
    if (nums[i] == nums[index]) {
      continue
    }
    index++
    nums[index] = nums[i]
  }
  return index + 1
}
