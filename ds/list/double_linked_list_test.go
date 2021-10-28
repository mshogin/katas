package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleLinkedListNew(t *testing.T) {
	a := assert.New(t)

	l := NewDoubleLinkedList()

	a.Nil(l.First)
	a.Nil(l.Last)
	a.Equal(l.First, l.Last)
}

func TestDoubleLinkedListAddLast(t *testing.T) {
	a := assert.New(t)

	l := NewDoubleLinkedList()

	firstNode := &DoubleLinkedListNode{Value: 1}
	l.AddLast(firstNode)
	a.Equal(firstNode, l.Last)
	a.Equal(1, l.Length())
	a.NotNil(l.First)
	a.NotNil(l.Last)
	a.Equal(l.First, l.Last)
	a.Nil(firstNode.Prev)
	a.Nil(firstNode.Next)
	a.Equal("1", l.ForwardTraverseToString())
	a.Equal("1", l.BackwardTraverseToString())

	secondNode := &DoubleLinkedListNode{Value: 2}
	l.AddLast(secondNode)
	a.Equal(2, l.Length())
	a.NotEqual(l.First, l.Last)
	a.Equal(firstNode, l.First)
	a.Equal(secondNode, l.Last)
	a.NotNil(secondNode.Prev)
	a.Nil(secondNode.Next)
	a.Equal(secondNode, firstNode.Next)
	a.Equal(firstNode, secondNode.Prev)
	a.Equal("1 2", l.ForwardTraverseToString())
	a.Equal("2 1", l.BackwardTraverseToString())

	thirdNode := &DoubleLinkedListNode{Value: 3}
	l.AddLast(thirdNode)
	a.Equal(3, l.Length())
	a.NotEqual(l.First, l.Last)

	a.Equal(firstNode, l.First)
	a.Equal(thirdNode, l.Last)

	a.NotNil(thirdNode.Prev)
	a.Nil(thirdNode.Next)

	a.Equal(thirdNode, secondNode.Next)
	a.Equal(secondNode, thirdNode.Prev)

	a.Equal(firstNode, secondNode.Prev)
	a.Equal(secondNode, firstNode.Next)
	a.Equal("1 2 3", l.ForwardTraverseToString())
	a.Equal("3 2 1", l.BackwardTraverseToString())
}

func TestDoubleLinkedListAddFirst(t *testing.T) {
	a := assert.New(t)
	l := NewDoubleLinkedList()

	firstNode := &DoubleLinkedListNode{Value: 1}
	listNode := l.AddFirst(firstNode)
	a.Equal(firstNode, listNode)
	a.Equal(1, l.Length())
	a.NotNil(l.First)
	a.NotNil(l.Last)
	a.Equal(l.First, l.Last)
	a.Nil(firstNode.Prev)
	a.Nil(firstNode.Next)
	a.Equal("1", l.ForwardTraverseToString())
	a.Equal("1", l.BackwardTraverseToString())

	secondNode := &DoubleLinkedListNode{Value: 2}
	l.AddFirst(secondNode)
	a.Equal(2, l.Length())
	a.NotEqual(l.First, l.Last)
	a.Equal(secondNode, l.First)
	a.Equal(firstNode, l.Last)
	a.NotNil(firstNode.Prev)
	a.Nil(secondNode.Prev)
	a.Equal(secondNode, firstNode.Prev)
	a.Equal(firstNode, secondNode.Next)
	a.Equal("2 1", l.ForwardTraverseToString())
	a.Equal("1 2", l.BackwardTraverseToString())

	thirdNode := &DoubleLinkedListNode{Value: 3}
	l.AddFirst(thirdNode)
	a.Equal(3, l.Length())
	a.NotEqual(l.First, l.Last)
	a.Equal(thirdNode, l.First)
	a.Equal(firstNode, l.Last)
	a.NotNil(thirdNode.Next)
	a.Nil(thirdNode.Prev)
	a.Equal(thirdNode, secondNode.Prev)
	a.Equal(secondNode, thirdNode.Next)

	a.Equal("3 2 1", l.ForwardTraverseToString())
	a.Equal("1 2 3", l.BackwardTraverseToString())
}

func TestDoubleLinkedListAddAfter(t *testing.T) {
	a := assert.New(t)
	l := NewDoubleLinkedList()

	nextNode := &DoubleLinkedListNode{Value: 1}
	_, err := l.AddAfter(l.First, nextNode)
	a.Error(err)
	_, err = l.AddAfter(l.First, &DoubleLinkedListNode{Value: 100})
	a.Error(err)

	node1 := l.AddLast(&DoubleLinkedListNode{Value: 1})
	node3 := l.AddLast(&DoubleLinkedListNode{Value: 3})
	a.NotNil(node3)
	a.NotNil(node1)
	a.Equal("1 3", l.ForwardTraverseToString())
	a.Equal("3 1", l.BackwardTraverseToString())
	a.Equal(l.Last, node3)
	a.Equal(l.Last.Prev, node1)

	node2, err := l.AddAfter(node1, &DoubleLinkedListNode{Value: 2})
	a.NoError(err)
	a.NotNil(node2)
	a.Equal("1 2 3", l.ForwardTraverseToString())
	a.Equal("3 2 1", l.BackwardTraverseToString())
	a.Equal(l.Last, node3)
	a.Equal(l.Last.Prev, node2)

	node4, err := l.AddAfter(node3, &DoubleLinkedListNode{Value: 4})
	a.NoError(err)
	a.Equal(l.Last, node4)
	a.Equal(l.Last.Prev, node3)
	a.Equal("1 2 3 4", l.ForwardTraverseToString())
	a.Equal("4 3 2 1", l.BackwardTraverseToString())
}

func TestDoubleLinkedListAddBefore(t *testing.T) {
	a := assert.New(t)
	l := NewDoubleLinkedList()

	nextNode := &DoubleLinkedListNode{Value: 1}
	for _, node := range []*DoubleLinkedListNode{l.First, l.Last} {
		_, err := l.AddBefore(node, nextNode)
		a.Error(err)
		_, err = l.AddBefore(node, &DoubleLinkedListNode{Value: 100})
		a.Error(err)
	}

	node2 := l.AddLast(&DoubleLinkedListNode{Value: 2})
	node4 := l.AddLast(&DoubleLinkedListNode{Value: 4})
	a.Equal("2 4", l.ForwardTraverseToString())
	a.Equal("4 2", l.BackwardTraverseToString())

	node3, err := l.AddBefore(node4, &DoubleLinkedListNode{Value: 3})
	a.NoError(err)
	a.NotNil(node3)
	a.Equal("2 3 4", l.ForwardTraverseToString())
	a.Equal("4 3 2", l.BackwardTraverseToString())

	node1, err := l.AddBefore(node2, &DoubleLinkedListNode{Value: 1})
	a.NoError(err)
	a.NotNil(node1)
	a.Equal("1 2 3 4", l.ForwardTraverseToString())
	a.Equal("4 3 2 1", l.BackwardTraverseToString())
}

func TestHasNode(t *testing.T) {
	a := assert.New(t)
	l := NewDoubleLinkedList()
	for _, v := range []int{1, 2, 3, 4, 5} {
		node := l.AddLast(&DoubleLinkedListNode{Value: v})
		a.True(l.HasNode(node))
	}
}

func TestForwardTraverse(t *testing.T) {
	a := assert.New(t)
	l := NewDoubleLinkedList()
	for _, v := range []int{1, 2, 3, 4, 5} {
		l.AddLast(&DoubleLinkedListNode{Value: v})
	}

	counter := 1
	l.ForwardTraverse(func(item *DoubleLinkedListNode) {
		value, ok := item.Value.(int)
		a.True(ok)
		a.Equal(counter, value)
		counter++
	})
}

func TestBackwardTraverse(t *testing.T) {
	a := assert.New(t)
	l := NewDoubleLinkedList()
	for _, v := range []int{1, 2, 3, 4, 5} {
		l.AddLast(&DoubleLinkedListNode{Value: v})
	}

	counter := 5
	l.BackwardTraverse(func(item *DoubleLinkedListNode) {
		value, ok := item.Value.(int)
		a.True(ok)
		a.Equal(counter, value)
		counter--
	})
}

func TestDoubleLinkedListRemove(t *testing.T) {
	a := assert.New(t)
	l := NewDoubleLinkedList()

	node1 := l.AddLast(&DoubleLinkedListNode{Value: 1})
	node2 := l.AddLast(&DoubleLinkedListNode{Value: 2})
	node3 := l.AddLast(&DoubleLinkedListNode{Value: 3})
	a.Equal("1 2 3", l.ForwardTraverseToString())
	a.Equal("3 2 1", l.BackwardTraverseToString())

	a.Equal(3, l.Length())

	l.Remove(node2)
	a.Equal(2, l.Length())
	a.Equal("1 3", l.ForwardTraverseToString())
	a.Equal("3 1", l.BackwardTraverseToString())

	l.Remove(node1)
	a.Equal("3", l.ForwardTraverseToString())
	a.Equal("3", l.BackwardTraverseToString())
	a.Equal(1, l.Length())
	a.False(l.HasNode(node1))
	a.True(l.HasNode(node3))

	l.Remove(node3)
	a.Equal(0, l.Length())
}
