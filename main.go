package main

func main() {
}

type Node struct {
	Value int
	Link  *Node
}

type LinkedList struct {
	Head *Node
}

func reverseLinkedList(list LinkedList) {
	var old = list.Head
	if old == nil {
		return
	}

	var pos = list.Head.Link
	var next *Node

	for pos != nil {
		next = pos.Link
		pos.Link = old
		old = pos
		pos = next
	}

}
