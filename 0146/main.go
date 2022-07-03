package main

import "fmt"

type LRUCache struct {
	Capacity   int
	NMap       map[int]*Node
	Head, Tail *Node
}
type Node struct {
	Key, Value int
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := Node{}, Node{}
	head.Next = &tail
	tail.Prev = &head
	return LRUCache{Capacity: capacity, NMap: make(map[int]*Node, capacity), Head: &head, Tail: &tail}
}

func (this *LRUCache) Get(key int) int {
	node, found := this.NMap[key]
	if found {
		this.remove(node)
		this.insert(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, found := this.NMap[key]
	if found {
		this.remove(node)
	}
	newNode := Node{Key: key, Value: value}
	this.insert(&newNode)
	if len(this.NMap) > this.Capacity {
		last := this.Tail.Prev
		this.remove(last)
	}
}

func (this *LRUCache) remove(node *Node) {
	delete(this.NMap, node.Key)
	next := node.Next
	prev := node.Prev
	next.Prev = prev
	prev.Next = next
}

func (this *LRUCache) insert(node *Node) {
	this.NMap[node.Key] = node

	next := this.Head.Next

	this.Head.Next = node
	node.Prev = this.Head

	next.Prev = node
	node.Next = next
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
}
