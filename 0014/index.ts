function longestCommonPrefix(strs: string[]): string {
  let ret = ''
  if (strs.length == 0 || strs == undefined || strs == null) {
    return ret
  }
  for (let firstIndex = 0; firstIndex < strs[0].length; firstIndex++) {
    for (let otherIndex = 1; otherIndex < strs.length; otherIndex++) {
      if (
        firstIndex == strs[otherIndex].length ||
        strs[0][firstIndex] != strs[otherIndex][firstIndex]
      ) {
        return ret
      }
    }
    ret += strs[0][firstIndex]
  }
  return ret
}

const strs = ['flower', 'flow', 'flight']

console.log(longestCommonPrefix(strs))
