package tree

import (
	"fmt"
	"math"
)

type bstTree struct {
	Data int

	Left  *bstTree
	Right *bstTree
}

// NewBst - create new binary search tree
func NewBst(val int) *bstTree {
	return &bstTree{Data: val}
}

// GetHeight - get tree heightn
func (m *bstTree) GetHeight() int {
	if m == nil {
		return 0
	}
	leftHeight := 1 + m.Left.GetHeight()
	rightHeight := 1 + m.Right.GetHeight()
	return int(
		math.Max(float64(leftHeight),
			float64(rightHeight)),
	)
}

// Add - add new item to the tree
func (m *bstTree) Add(val int) {
	var insert func(**bstTree, *bstTree)
	insert = func(testNode **bstTree, newNode *bstTree) {
		if *testNode == nil {
			*testNode = newNode
			return
		}
		if newNode.Data > (*testNode).Data {
			insert(&(*testNode).Right, newNode)
		} else {
			insert(&(*testNode).Left, newNode)
		}
	}
	insert(&m, &bstTree{Data: val})
}

// IsValid - test binary search tree rools
func (m *bstTree) IsValid() bool {
	if m == nil {
		return true
	}
	if m.Left != nil && m.Left.Data > m.Data {
		return false
	}
	if m.Right != nil && m.Right.Data < m.Data {
		return false
	}
	return m.Left.IsValid() && m.Right.IsValid()
}

// Search - search the node by given value
func (m *bstTree) Search(val int) *bstTree {
	n, _ := m.search(nil, val)
	return n
}

func (m *bstTree) search(parent *bstTree, val int) (*bstTree, *bstTree) {
	if m.Data == val {
		return m, parent
	}
	if m.Right != nil && val > m.Data {
		return m.Right.search(m, val)
	}
	if m.Left != nil && val < m.Data {
		return m.Left.search(m, val)
	}
	return nil, nil
}

func (m *bstTree) isLeaf() bool {
	return m.Left == nil && m.Right == nil
}

func (m *bstTree) getRightMin(parent *bstTree) (*bstTree, *bstTree) {
	if m.Left == nil {
		return m, parent
	}
	return m.Left.getRightMin(m)
}

func (m *bstTree) getLeftMax(parent *bstTree) (*bstTree, *bstTree) {
	if m.Right == nil {
		return m, parent
	}
	return m.Right.getLeftMax(m)
}

// Delete - remove element from the tree
func (m *bstTree) Delete(val int) bool {
	unlink := func(node *bstTree, parent *bstTree) {
		if node == parent.Left {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	}
	var remove func(*bstTree, *bstTree) bool
	remove = func(node *bstTree, parent *bstTree) bool {
		if node == nil {
			return false
		}

		if node.isLeaf() {
			unlink(node, parent)
			return true
		}

		if m == node {
			if node.Right != nil {
				n, p := m.Right.getRightMin(nil)
				unlink(n, p)
				node.Data = n.Data

			} else {
				m = m.Left
			}
			return true
		}

		var n, p *bstTree
		if node.Data > m.Data {
			n, p = node.getRightMin(nil)
		} else {
			n, p = node.getLeftMax(nil)
		}
		unlink(n, p)
		node.Data = n.Data

		return true
	}
	node, parent := m.search(nil, val)
	return remove(node, parent)

}

// Print - print the tree
func (m *bstTree) Print() {
	if m == nil {
		return
	}
	m.Left.Print()
	fmt.Printf(" %+v \n", m) // output for debug
	m.Right.Print()
}
