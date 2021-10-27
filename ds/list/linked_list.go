package list

import "fmt"

type LinkedList struct {
	First *LinkedListNode
	Last  *LinkedListNode

	count int
}

type LinkedListNode struct {
	Next *LinkedListNode
}

// NewLinkedList - returns new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// NewLinkedListNode - returns new linked list node
func NewLinkedListNode() *LinkedListNode {
	return &LinkedListNode{}
}

// Add - add new node to the list tail
func (l *LinkedList) Add(node *LinkedListNode) {
	if l.First == nil {
		l.First = node
	} else {
		l.Last.Next = node
	}

	l.Last = node

	l.count++
}

// AddFirst - add new node at the beginning of the list
func (l *LinkedList) AddFirst(node *LinkedListNode) {
	node.Next = l.First
	l.First = node

	l.count++
}

// RemoveNode - remove node from the list
func (l *LinkedList) RemoveNode(node *LinkedListNode) bool {
	removed := false

	if node == l.First {
		l.First = l.First.Next
		node.Next = nil
		removed = true
	} else {
		for n := l.First; n != nil; {
			if n.Next == node {
				if node == l.Last {
					l.Last = n
				}
				n.Next = node.Next
				removed = true
				break
			}
			n = n.Next
		}
	}
	if removed {
		l.count--
	}

	return removed
}

// Print - print list nodes to stdout
func (l *LinkedList) Print() {
	for n := l.First; n != nil; n = n.Next {
		fmt.Printf("node: %p=%+v\n", n, n) // output for debug

	}
}

// Length - returns number of element in the list
func (l *LinkedList) Length() int {
	return l.count
}
