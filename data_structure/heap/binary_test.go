package heap

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlub.com/mshogin/katas/utils"
)

func TestBinaryHeap(t *testing.T) {
	a := assert.New(t)

	h := NewBinaryHeap()
	a.NotNil(h)
	a.True(h.IsValid())
	a.True(h.IsEmpty())
	a.Equal(0, h.GetSize())

	h.Insert(1)
	a.True(h.IsValid())
	a.False(h.IsEmpty())
	a.Equal(1, h.GetSize())

	h.Insert(2)
	a.True(h.IsValid())
	a.Equal(2, h.GetSize())

	values := make([]int, 10)
	for i := 0; i < len(values); i++ {
		values[i] = rand.Intn(100)
	}

	values = utils.Unique(values)
	for i := 0; i < len(values); i++ {
		h.Insert(values[i])
		a.True(h.IsValid(), "insert %v", values[i])
	}

	idx := rand.Intn(10)
	h.Replace(idx, 44)
	a.True(h.IsValid())

	h.Print(0, 0)

	max := h.GetMax()
	for _, v := range values {
		a.GreaterOrEqual(max, v)
	}
	a.True(h.IsValid())
	h.Print(0, 0)
}
