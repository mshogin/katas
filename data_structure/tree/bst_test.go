package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBst(t *testing.T) {
	a := assert.New(t)

	tree := NewBst(5)
	a.NotNil(tree)
	a.True(tree.IsValid())
	a.Nil(tree.Left)
	a.Nil(tree.Right)
	a.Equal(1, tree.GetHeight())

	tree.Add(7)
	a.True(tree.IsValid())
	a.Nil(tree.Left)
	a.NotNil(tree.Right)
	a.Equal(2, tree.GetHeight())

	tree.Add(3)
	a.True(tree.IsValid())
	a.NotNil(tree.Left)
	a.NotNil(tree.Right)
	a.Equal(2, tree.GetHeight())

	tree.Add(2)
	a.True(tree.IsValid())
	a.Equal(3, tree.GetHeight())

	tree.Add(4)
	a.True(tree.IsValid())
	a.Equal(3, tree.GetHeight())

	tree.Add(6)
	a.True(tree.IsValid())
	a.Equal(3, tree.GetHeight())

	tree.Add(9)
	a.True(tree.IsValid())
	a.Equal(3, tree.GetHeight())

	tree.Add(8)
	a.True(tree.IsValid())
	a.Equal(4, tree.GetHeight())

	/*
			       5
			   3       7
			 2   4   6   9
		                    8
	*/

	node7 := tree.Search(7)
	a.NotNil(node7)
	a.Equal(7, node7.Data)
	a.Equal(tree.Right, node7)

	node9 := tree.Search(9)
	a.NotNil(node9)
	a.Equal(9, node9.Data)
	a.Equal(node7.Right, node9)

	node6 := tree.Search(6)
	a.NotNil(node6)
	a.Equal(6, node6.Data)
	a.Equal(node7.Left, node6)

	node3 := tree.Search(3)
	node4 := tree.Search(4)
	a.Equal(node3.Right, node4)

	notFound := tree.Search(1001)
	a.Nil(notFound)

	n, p := tree.getRightMin(nil)
	a.Equal(tree.Search(2), n)
	a.Equal(tree.Search(3), p)
	n, p = tree.Search(7).getRightMin(nil)
	a.Equal(tree.Search(6), n)
	a.Equal(tree.Search(7), p)

	a.True(tree.Delete(3))
	a.Nil(tree.Search(3))
	a.Equal(tree.Left, tree.Search(4))
	a.Equal(tree.Search(2), tree.Search(4).Left)

	a.True(tree.Delete(8))
	a.Nil(tree.Search(8))
	a.Equal(tree.Right, tree.Search(7))
	a.Equal(tree.Search(6), tree.Search(7).Left)
	a.Equal(tree.Search(9), tree.Search(7).Right)

	a.True(tree.Delete(5)) // remove root
	a.Nil(tree.Search(5))
	a.Equal(tree, tree.Search(6))
	a.Equal(tree.Search(4), tree.Left)
	a.Equal(tree.Search(7), tree.Right)

}
