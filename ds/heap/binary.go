package heap

type binaryHeap struct {
	size int

	items []int
}

// NewBinaryHeap - create new binary heap
func NewBinaryHeap() Heap {
	return &binaryHeap{}
}

func (h *binaryHeap) siftup(childIdx int) {
	parentIdx := (childIdx - 1) / 2
	if h.items[childIdx] > h.items[parentIdx] {
		h.swap(childIdx, parentIdx)
		h.siftup(parentIdx)
	}
}

func (h *binaryHeap) siftdown(parentIdx int) {
	leftChildIdx := parentIdx*2 + 1
	rightChildIdx := parentIdx*2 + 2

	for _, childIdx := range []int{leftChildIdx, rightChildIdx} {
		if h.hasItem(childIdx) && h.items[childIdx] > h.items[parentIdx] {
			h.swap(childIdx, parentIdx)
			h.siftdown(childIdx)
			h.siftdown(parentIdx)
			break
		}
	}
}

func (h *binaryHeap) swap(idx1, idx2 int) {
	h.items[idx1], h.items[idx2] = h.items[idx2], h.items[idx1]
}

func (h *binaryHeap) resize(delta int) {
	items := make([]int, (h.size+delta)*2)
	copy(items[:], h.items)
	h.items = items
}

func (h *binaryHeap) hasItem(idx int) bool {
	return idx < h.GetSize()
}

// IsValid - validate the structure is heap
// any parent element should be greater then any of child
func (h *binaryHeap) IsValid() bool {
	for i := 0; i < h.size/2; i++ {
		left := 2*i + 1
		right := 2*i + 2
		if h.hasItem(left) && h.items[left] > h.items[i] {
			return false
		}
		if h.hasItem(right) && h.items[right] > h.items[i] {
			return false
		}
	}
	return true
}

// IsEmpty - inspect if heap is empty
func (h *binaryHeap) IsEmpty() bool {
	return h.GetSize() == 0
}

// GetSize - return heap size
func (h *binaryHeap) GetSize() int {
	return h.size
}

// Insert - insert new elements to the heap
func (h *binaryHeap) Insert(items ...int) {
	if h.size+len(items) >= len(h.items) {
		h.resize(len(items))
	}
	for _, element := range items {
		h.items[h.size] = element
		h.size++
		h.siftup(h.size - 1)
	}
}

// Index - find index of element by value
func (h *binaryHeap) Index(val int) int {
	for i, v := range h.items {
		if v == val {
			return i
		}
	}
	return -1
}

// Replace - replace value of element
func (h *binaryHeap) Replace(itemIdx int, val int) {
	oldVal := h.items[itemIdx]
	h.items[itemIdx] = val
	if oldVal > val {
		h.siftdown(itemIdx)
	} else {
		h.siftup(itemIdx)
	}
}

// GetMax - get max value from the heap
func (h *binaryHeap) GetMax() int {
	max := h.items[0]
	h.items[0], h.items[h.size-1] = h.items[h.size-1], 0
	h.size--
	h.siftdown(0)
	return max
}
