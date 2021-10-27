package heap

import "fmt"

const (
	newLine      = "\n"
	emptySpace   = "    "
	middleItem   = "├── "
	continueItem = "│   "
	lastItem     = "└── "
)

func (h *binaryHeap) Print(idx int, offset int) {
	var hasItem = func(idx int) bool {
		return idx < h.GetSize()
	}
	if !hasItem(idx) {
		fmt.Printf("%+v\n", "Wrong element index")
		return
	}

	fmt.Printf("%+v", h.items[idx])
	fmt.Printf("%+v", newLine)
	if hasItem(2*idx + 1) {
		if offset > 0 {
			fmt.Printf("%+v", continueItem)
		}
		for i := 0; i < offset-1; i++ {
			fmt.Printf("%+v", continueItem)
		}
		fmt.Printf("%+v", middleItem)
		h.Print(2*idx+1, offset+1)
	}
	if hasItem(2*idx + 2) {
		if offset > 0 {
			fmt.Printf("%+v", continueItem)
		}
		for i := 0; i < offset-1; i++ {
			fmt.Printf("%+v", continueItem)
		}
		fmt.Printf("%+v", middleItem)
		h.Print(2*idx+2, offset+1)
	}
}
