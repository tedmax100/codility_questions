package main

func main() {

}

func Process(args Input_Struct) []*int {
	var cache = Constructor(args.Values[0][0])
	result := make([]*int, 0)

	inputData := args.Actions[1:]
	for idx, action := range inputData {
		if action == "put" {
			kv := args.Values[idx+1]
			cache.Put(kv[0], kv[1])
			result = append(result, nil)
		}
		if action == "get" {
			value := cache.Get(args.Values[idx+1][0])
			result = append(result, &value)
		}
	}
	return result
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	cap  int
	head *Node
	tail *Node
	data map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{Prev: head}
	head.Next = tail
	return LRUCache{
		cap:  capacity,
		head: head,
		tail: tail,
		data: make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.data[key]; ok {
		this.unlink(node)
		this.prepend(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.data[key]; ok {
		this.unlink(node)
		node.Val = value
		this.prepend(node)
		return
	}

	if this.cap < 1 {
		lastNode := this.tail.Prev
		if lastNode == this.head {
			return
		}
		this.unlink(lastNode)
		delete(this.data, lastNode.Key)
		this.cap++
	}

	newNode := &Node{Key: key, Val: value}
	this.data[key] = newNode
	this.prepend(newNode)
	this.cap--
}

func (this *LRUCache) prepend(node *Node) {
	node.Next = this.head.Next
	node.Prev = this.head
	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) unlink(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Next = nil
	node.Prev = nil
}
