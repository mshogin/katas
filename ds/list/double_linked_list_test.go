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

	firstNode := &doubleLinkedListNode{Value: 1}
	l.AddLast(firstNode)
	a.Equal(firstNode, l.Last)
	a.Equal(1, l.Length())
	a.NotNil(l.First)
	a.NotNil(l.Last)
	a.Equal(l.First, l.Last)
	a.Nil(firstNode.Prev)
	a.Nil(firstNode.Next)

	secondNode := &doubleLinkedListNode{Value: 2}
	l.AddLast(secondNode)
	a.Equal(2, l.Length())
	a.NotEqual(l.First, l.Last)
	a.Equal(firstNode, l.First)
	a.Equal(secondNode, l.Last)
	a.NotNil(secondNode.Prev)
	a.Nil(secondNode.Next)
	a.Equal(secondNode, firstNode.Next)
	a.Equal(firstNode, secondNode.Prev)

	thirdNode := &doubleLinkedListNode{Value: 3}
	l.AddLast(thirdNode)
	a.Equal(3, l.Length())
	a.NotEqual(l.First, l.Last)
	a.Equal(firstNode, l.First)
	a.Equal(thirdNode, l.Last)
	a.NotNil(thirdNode.Prev)
	a.Nil(thirdNode.Next)
	a.Equal(thirdNode, secondNode.Next)
	a.Equal(secondNode, thirdNode.Prev)
}

func TestDoubleLinkedListAddFirst(t *testing.T) {
	a := assert.New(t)
	l := NewDoubleLinkedList()

	firstNode := &doubleLinkedListNode{Value: 1}
	listNode := l.AddFirst(firstNode)
	a.Equal(firstNode, listNode)
	a.Equal(1, l.Length())
	a.NotNil(l.First)
	a.NotNil(l.Last)
	a.Equal(l.First, l.Last)
	a.Nil(firstNode.Prev)
	a.Nil(firstNode.Next)

	secondNode := &doubleLinkedListNode{Value: 2}
	l.AddFirst(secondNode)
	a.Equal(2, l.Length())
	a.NotEqual(l.First, l.Last)
	a.Equal(secondNode, l.First)
	a.Equal(firstNode, l.Last)
	a.NotNil(firstNode.Prev)
	a.Nil(secondNode.Prev)
	a.Equal(secondNode, firstNode.Prev)
	a.Equal(firstNode, secondNode.Next)

	thirdNode := &doubleLinkedListNode{Value: 3}
	l.AddFirst(thirdNode)
	a.Equal(3, l.Length())
	a.NotEqual(l.First, l.Last)
	a.Equal(thirdNode, l.First)
	a.Equal(firstNode, l.Last)
	a.NotNil(thirdNode.Next)
	a.Nil(thirdNode.Prev)
	a.Equal(thirdNode, secondNode.Prev)
	a.Equal(secondNode, thirdNode.Next)
}

func TestDoubleLinkedListAddAfter(t *testing.T) {
	a := assert.New(t)
	l := NewDoubleLinkedList()

	nextNode := &doubleLinkedListNode{Value: 1}
	_, err := l.AddAfter(l.First, nextNode)
	a.Error(err)
	_, err = l.AddAfter(l.First, &doubleLinkedListNode{Value: 100})
	a.Error(err)

	node1 := l.AddLast(&doubleLinkedListNode{Value: 1})
	node3 := l.AddLast(&doubleLinkedListNode{Value: 3})

	node2, err := l.AddAfter(node1, &doubleLinkedListNode{Value: 2})
	a.NoError(err)
	a.NotNil(node2)
	a.Equal(node1, node2.Prev)
	a.Equal(node3, node2.Next)
	a.Equal(node2, node1.Next)
	a.Equal(node2, node3.Prev)

	node4, err := l.AddAfter(node3, &doubleLinkedListNode{Value: 4})
	a.NoError(err)
	a.NotNil(node4)
	a.Equal(node3, node4.Prev)
	a.Equal(node4, node3.Next)
	a.Nil(node4.Next)
}

func TestDoubleLinkedListAddBefore(t *testing.T) {
	a := assert.New(t)
	l := NewDoubleLinkedList()

	nextNode := &doubleLinkedListNode{Value: 1}
	for _, node := range []*doubleLinkedListNode{l.First, l.Last} {
		_, err := l.AddBefore(node, nextNode)
		a.Error(err)
		_, err = l.AddBefore(node, &doubleLinkedListNode{Value: 100})
		a.Error(err)
	}

	node2 := l.AddLast(&doubleLinkedListNode{Value: 2})
	node4 := l.AddLast(&doubleLinkedListNode{Value: 4})

	node3, err := l.AddBefore(node4, &doubleLinkedListNode{Value: 3})
	a.NoError(err)
	a.NotNil(node3)
	a.Equal(node2, node3.Prev)
	a.Equal(node4, node3.Next)
	a.Equal(node3, node2.Next)
	a.Equal(node3, node4.Prev)

	node1, err := l.AddBefore(node2, &doubleLinkedListNode{Value: 1})
	a.NoError(err)
	a.NotNil(node1)
	a.Equal(node1, node2.Prev)
	a.Equal(node2, node1.Next)
	a.Nil(node1.Prev)
}

func TestHasNode(t *testing.T) {
	a := assert.New(t)
	l := NewDoubleLinkedList()
	for _, v := range []int{1, 2, 3, 4, 5} {
		node := l.AddLast(&doubleLinkedListNode{Value: v})
		a.True(l.HasNode(node))
	}
}

func TestForwardTraverse(t *testing.T) {
	a := assert.New(t)
	l := NewDoubleLinkedList()
	for _, v := range []int{1, 2, 3, 4, 5} {
		l.AddLast(&doubleLinkedListNode{Value: v})
	}

	counter := 1
	l.ForwardTraverse(func(item *doubleLinkedListNode) {
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
		l.AddLast(&doubleLinkedListNode{Value: v})
	}

	counter := 5
	l.BackwardTraverse(func(item *doubleLinkedListNode) {
		value, ok := item.Value.(int)
		a.True(ok)
		a.Equal(counter, value)
		counter--
	})
}

func TestDoubleLinkedListRemove(t *testing.T) {
	a := assert.New(t)
	l := NewDoubleLinkedList()

	node1 := l.AddLast(&doubleLinkedListNode{Value: 1})
	node2 := l.AddLast(&doubleLinkedListNode{Value: 2})
	node3 := l.AddLast(&doubleLinkedListNode{Value: 3})

	a.Equal(3, l.Length())

	l.Remove(node2)
	a.Equal(2, l.Length())
	a.True(l.HasNode(node1))
	a.False(l.HasNode(node2))
	a.True(l.HasNode(node3))

	l.Remove(node1)
	a.Equal(1, l.Length())
	a.False(l.HasNode(node1))
	a.True(l.HasNode(node3))

	l.Remove(node3)
	a.Equal(0, l.Length())
}
