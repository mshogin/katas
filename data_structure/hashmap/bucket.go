package hashmap

type bucket interface {
	Add(key string, value interface{}) bool
	Get(key string) interface{}
	GetKeys() []string
}

type createBucketFunc func() bucket

type basicBucket struct {
	keys []string
	vals []interface{}
}

func (b *basicBucket) hasKey(key string) int {
	for i, k := range b.keys {
		if k == key {
			return i
		}
	}
	return -1
}

// Add - add key value to the bucket
func (b *basicBucket) Add(key string, val interface{}) bool {
	if i := b.hasKey(key); i >= 0 {
		b.vals[i] = val
		return false
	}
	b.keys = append(b.keys, key)
	b.vals = append(b.vals, val)
	return true
}

// Get - get key from the bucket
func (b *basicBucket) Get(key string) interface{} {
	for i, k := range b.keys {
		if key == k {
			return b.vals[i]
		}
	}
	return nil
}

// GetKeys ...
func (b *basicBucket) GetKeys() []string {
	return b.keys
}
