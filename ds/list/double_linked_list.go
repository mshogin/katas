package list

import (
	"errors"
	"fmt"
	"strings"
)

type doubleLinkedListNode struct {
	Prev *doubleLinkedListNode
	Next *doubleLinkedListNode

	Value interface{}
}

type DoubleLinkedList struct {
	First *doubleLinkedListNode
	Last  *doubleLinkedListNode

	count int
}

// NewDoubleLinkedList - create new double linked list
func NewDoubleLinkedList() *DoubleLinkedList {
	l := &DoubleLinkedList{}
	l.Last = l.First
	return l
}

// AddLast - adds new node at the end of the list and returs added node
func (l *DoubleLinkedList) AddLast(n *doubleLinkedListNode) *doubleLinkedListNode {
	l.count++
	if l.Last == nil {
		l.First = n
	} else {
		l.Last.Next = n
		n.Prev = l.Last
		l.Last = n
	}
	l.Last = n
	return n
}

// AddFirst - adds new node at the beginning of the list and returs added node
func (l *DoubleLinkedList) AddFirst(n *doubleLinkedListNode) *doubleLinkedListNode {
	l.count++
	if l.Last == nil {
		l.Last = n
	} else {
		l.First.Prev = n
		n.Next = l.First
	}
	l.First = n
	return n
}

// Length - returns length of the list
func (l *DoubleLinkedList) Length() int {
	return l.count
}

// AddAfter - adds new node after existing node
func (l *DoubleLinkedList) AddAfter(inListNode *doubleLinkedListNode, newNode *doubleLinkedListNode) (*doubleLinkedListNode, error) {
	if inListNode == nil {
		return nil, errors.New("The node after which we need to add a node is nil")
	}
	if !l.HasNode(inListNode) {
		return nil, fmt.Errorf("The node(%v) does not belongs to the list", inListNode)
	}
	l.count++

	nextNode := inListNode.Next

	inListNode.Next = newNode
	newNode.Prev = inListNode

	if nextNode != nil {
		newNode.Next = nextNode
		nextNode.Prev = newNode
	}
	if l.Last == inListNode {
		l.Last = newNode
	}

	return newNode, nil
}

// AddBefore - adds new node before existing node
func (l *DoubleLinkedList) AddBefore(inListNode *doubleLinkedListNode, newNode *doubleLinkedListNode) (*doubleLinkedListNode, error) {
	if inListNode == nil {
		return nil, errors.New("The node after which we need to add a node is nil")
	}
	if !l.HasNode(inListNode) {
		return nil, fmt.Errorf("The node(%v) does not belongs to the list", inListNode)
	}
	l.count++

	nodePrev := inListNode.Prev

	inListNode.Prev = newNode
	newNode.Next = inListNode

	if nodePrev != nil {
		nodePrev.Next = newNode
		newNode.Prev = nodePrev
	}
	if l.First == inListNode {
		l.First = newNode
	}
	return newNode, nil
}

func (l *DoubleLinkedList) HasNode(node *doubleLinkedListNode) bool {
	found := false
	l.ForwardTraverse(
		func(nextNode *doubleLinkedListNode) {
			tmpFound := nextNode == node
			found = found || tmpFound
		},
	)
	return found
}

// ForwardTraverse - traverse the list from first to last
func (l *DoubleLinkedList) ForwardTraverse(callback func(*doubleLinkedListNode)) {
	item := l.First
	for item != nil {
		callback(item)
		item = item.Next
	}
}

// BackwardTraverse - traverse the list from the last to the first
func (l *DoubleLinkedList) BackwardTraverse(callback func(*doubleLinkedListNode)) {
	item := l.Last
	for item != nil {
		callback(item)
		item = item.Prev
	}
}

// BackwardTraverse - traverse the list from the last to the first
func (l *DoubleLinkedList) Remove(node *doubleLinkedListNode) {
	prevNode := node.Prev
	nextNode := node.Next
	node.Next = nil
	node.Prev = nil

	if prevNode != nil {
		prevNode.Next = nextNode
	}
	if nextNode != nil {
		nextNode.Prev = prevNode
	}
	if node == l.First {
		l.First = nextNode
	}
	if node == l.Last {
		l.Last = nextNode
	}

	l.count--
}

func (l *DoubleLinkedList) Print() {
	l.ForwardTraverse(func(item *doubleLinkedListNode) {
		fmt.Printf("%p %+v\n", item, item)
	})
}

func (l *DoubleLinkedList) ForwardTraverseToString() string {
	r := ""
	l.ForwardTraverse(func(item *doubleLinkedListNode) {
		r = fmt.Sprintf("%s %v", r, item.Value)
	})
	return strings.Trim(r, " ")
}

func (l *DoubleLinkedList) BackwardTraverseToString() string {
	r := ""
	l.BackwardTraverse(func(item *doubleLinkedListNode) {
		r = fmt.Sprintf("%s %v", r, item.Value)
	})
	return strings.Trim(r, " ")
}
