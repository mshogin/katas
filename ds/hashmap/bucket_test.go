package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicBucket(t *testing.T) {
	a := assert.New(t)

	b := &basicBucket{}
	a.NotNil(b)

	a.True(b.Add("key1", 1))
	a.Equal(1, b.Get("key1"))

	a.False(b.Add("key1", 2))
	a.Equal(2, b.Get("key1"))

	a.True(b.Add("key2", 2))
	a.Equal(2, b.Get("key2"))

	a.Equal([]string{"key1", "key2"}, b.GetKeys())
}
