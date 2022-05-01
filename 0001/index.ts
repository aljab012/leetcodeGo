function twoSum(nums: number[], target: number): number[] {
  let myMap = new Map<number, number>()

  for (var i = 0; i < nums.length; i++) {
    const complement = target - nums[i]
    if (myMap.has(complement)) {
      return [i, Number(myMap.get(complement))]
    }
    myMap.set(nums[i], i)
  }
  return [-1, -1]
}

const ret = twoSum([2, 7, 11, 15], 9)
console.log(ret)
