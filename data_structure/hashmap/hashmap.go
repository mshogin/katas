package hashmap

type hashMap struct {
	Size int

	hash         hashFunk
	buckets      []bucket
	createBucket createBucketFunc
}

var defaultCreateBucketFunc = func() bucket {
	return &basicBucket{}
}

func (m *hashMap) getBucket(key string) bucket {
	bucketID := m.hash(key)
	if bucketID > len(m.buckets) {
		m.createBuckets(bucketID)
	}
	return m.buckets[bucketID]
}

func (m *hashMap) createBuckets(size int) {
	buckets := make([]bucket, size*2)
	for i := 0; i < len(buckets); i++ {
		if i < len(m.buckets) {
			buckets[i] = m.buckets[i]
		} else {
			buckets[i] = m.createBucket()
		}
	}
	m.buckets = buckets
}

// NewHashMap - creates new hash map
func NewHashMap() *hashMap {
	m := &hashMap{}
	m.hash = defaultHashFunc
	m.createBucket = defaultCreateBucketFunc
	return m
}

// Add - add new pair (key and value) to the hash
func (m *hashMap) Add(key string, value interface{}) {
	b := m.getBucket(key)
	if b.Add(key, value) {
		m.Size++
	}
}

// Get - get value by key
func (m *hashMap) Get(key string) interface{} {
	b := m.getBucket(key)
	return b.Get(key)
}

// GetKeys - return keys
func (m *hashMap) GetKeys() []string {
	keys := []string{}
	for _, b := range m.buckets {
		keys = append(keys, b.GetKeys()...)
	}
	return keys
}

// GetValues - return values
func (m *hashMap) GetValues() []interface{} {
	values := make([]interface{}, m.Size)
	for i, k := range m.GetKeys() {
		values[i] = m.Get(k)
	}
	return values
}
