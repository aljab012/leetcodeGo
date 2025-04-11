package main

type Node struct {
	Val int
	Min int
}
type MinStack struct {
	stack []Node
}

func Constructor() MinStack {
	return MinStack{stack: []Node{}}
}

func (this *MinStack) Push(val int) {
	min := val
	if len(this.stack) > 0 {
		min = Min(val, this.stack[len(this.stack)-1].Min)
	}
	this.stack = append(this.stack, Node{Val: val, Min: min})
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1].Val
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1].Min
	}
	return -1
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
