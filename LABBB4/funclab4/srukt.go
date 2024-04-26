package funclab4

import "sync"

// Структура кольцевого буфера
type RingBufer struct {
	mu         sync.Mutex
	buf        []int
	size       int
	readPoint  int
	writePoint int
	count      int
	notEmpty   *sync.Cond
	notFull    *sync.Cond
}

// Создание нового буфера
func NewBuf(size int) *RingBufer {
	ringBuf := &RingBufer{
		buf:        make([]int, size),
		size:       size,
		readPoint:  0,
		writePoint: 0,
		count:      0,
	}

	ringBuf.notEmpty = sync.NewCond(&ringBuf.mu)
	ringBuf.notFull = sync.NewCond(&ringBuf.mu)

	return ringBuf
}
