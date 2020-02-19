package heap

type Heap interface {
	GetMax() int
	GetSize() int
	IsEmpty() bool
	Insert(...int)
	Index(int) int
	Replace(int, int)
	IsValid() bool

	Print(int, int)
}
