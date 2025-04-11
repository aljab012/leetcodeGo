package main

type MyQueue struct {
	pushStack []int
	popStack  []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.pushStack = append(this.pushStack, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}

	this.shift()
	ret := this.popStack[len(this.popStack)-1]
	this.popStack = this.popStack[:len(this.popStack)-1]
	return ret
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	this.shift()
	return this.popStack[len(this.popStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.pushStack)+len(this.popStack) == 0
}

func (this *MyQueue) shift() {
	if len(this.popStack) == 0 {
		for len(this.pushStack) > 0 {
			this.popStack = append(this.popStack, this.pushStack[len(this.pushStack)-1])
			this.pushStack = this.pushStack[:len(this.pushStack)-1]
		}
	}
}
