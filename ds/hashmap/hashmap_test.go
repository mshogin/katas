package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMap(t *testing.T) {
	a := assert.New(t)

	h := NewHashMap()
	a.NotNil(h)

	h.Add("key1", "value1")
	a.Equal(1, h.Size)
	a.Equal("value1", h.Get("key1"))

	h.Add("key1", "value12")
	a.Equal(1, h.Size)
	a.Equal("value12", h.Get("key1"))

	h.Add("key2", "value2")
	a.Equal(2, h.Size)
	a.Equal("value2", h.Get("key2"))

	a.Contains(h.GetKeys(), "key1")
	a.Contains(h.GetKeys(), "key2")

	a.Contains(h.GetValues(), "value12")
	a.Contains(h.GetValues(), "value2")
}
