function isValid(s: string): boolean {
  let stack = []
  let cMap = new Map<String, String>()
  cMap.set('(', ')')
  cMap.set('{', '}')
  cMap.set('[', ']')
  for (let char of s) {
    if (char === '(' || char === '{' || char === '[') {
      stack.push(char)
    } else {
      const popChar = '' + stack.pop()
      if (char != cMap.get(popChar)) {
        return false
      }
    }
  }
  return stack.length == 0
}

console.log(isValid('({}'))
