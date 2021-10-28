package list_test

import (
	"testing"

	"github.com/mshogin/katas/ds/list"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList1(t *testing.T) {
	a := assert.New(t)

	llist := list.NewLinkedList()
	a.NotNil(llist)
	a.Equal(0, llist.Length())

	node := list.NewLinkedListNode()
	a.NotNil(node)

	llist.Add(node)
	a.Equal(1, llist.Length())
	a.Equal(node, llist.First)
	a.Equal(node, llist.Last)

	node = list.NewLinkedListNode()
	llist.Add(node)
	a.Equal(2, llist.Length())
	a.NotEqual(llist.First, node)
	a.Equal(llist.First.Next, node)
	a.Equal(llist.Last, node)

	prevFirst := llist.First
	node = list.NewLinkedListNode()
	llist.AddFirst(node)
	a.Equal(3, llist.Length())
	a.Equal(llist.First, node)
	a.Equal(llist.First.Next, prevFirst)

	a.True(llist.RemoveNode(node)) // this is first node
	a.Equal(2, llist.Length())
	a.NotEqual(llist.First, node)
	a.NotNil(llist.Last)

	node = llist.Last
	a.True(llist.RemoveNode(node)) // this is first node
	a.NotNil(llist.Last)
	a.Equal(1, llist.Length())
	a.False(llist.Last == node)
}
