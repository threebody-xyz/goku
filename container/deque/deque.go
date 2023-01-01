package deque

const minCapacity = 16

type Deque[T any] struct {
	buffer []T
	head   int
	tail   int
	cap    int
}

// NewDeque create a deque with capacity to store 16 elements.
// For example, to create a deque that element type is int:
// d := NewDeque[int]()
func NewDeque[T any]() *Deque[T] {
	d := &Deque[T]{}
	d.cap = minCapacity
	d.buffer = make([]T, d.cap)
	return d
}

// NewDequeWithCap create a deque with capacity to store elements.
// Cap must be power of 2, and the number of elements is more than 16.
// For example, to create a deque that element type is int:
// d := NewDequeWithCap[int](10)
func NewDequeWithCap[T any](cap int) *Deque[T] {
	if cap < minCapacity {
		cap = minCapacity
	}
	d := &Deque[T]{}
	d.cap = lowPower(cap)
	d.buffer = make([]T, d.cap)
	return d
}

func (d *Deque[T]) resize(len int) {
	newCap := lowPower(len * 2)
	newBuffer := make([]T, newCap)
	newTail := d.Len()

	if d.tail > d.head {
		copy(newBuffer, d.buffer[d.head:d.tail])
	} else {
		n := copy(newBuffer, d.buffer[d.head:])
		copy(newBuffer[n:], d.buffer[:d.tail])
	}

	d.head = 0
	d.tail = newTail
	d.cap = newCap
	d.buffer = newBuffer
}

// PushFront appends an element to the front of deque.
func (d *Deque[T]) PushFront(values ...T) {
	if d.Len()+len(values) >= d.cap {
		d.resize(d.Len() + len(values))
	}

	for _, v := range values {
		d.head = d.prev(d.head)
		d.buffer[d.head] = v
	}
}

// PushBack appends an element to the back of deque.
func (d *Deque[T]) PushBack(values ...T) {
	if d.Len()+len(values) >= d.cap {
		d.resize(d.Len() + len(values))
	}

	for _, v := range values {
		d.buffer[d.tail] = v
		d.tail = d.next(d.tail)
	}
}

// PopFront returns the first element of deque or calls panic if the deque is empty.
func (d *Deque[T]) PopFront() T {
	if d.Len() <= 0 {
		panic("PopFront() called on empty deque")
	}

	result := d.buffer[d.head]
	d.head = d.next(d.head)

	if d.Cap() > minCapacity && d.Len() <= d.cap/4 {
		d.resize(d.Len())
	}
	return result
}

// PopBack returns the last element of deque or calls panic if the deque is empty.
func (d *Deque[T]) PopBack() T {
	if d.Len() <= 0 {
		panic("PopBack() called on empty deque")
	}

	d.tail = d.prev(d.tail)
	result := d.buffer[d.tail]

	if d.Cap() > minCapacity && d.Len() <= d.cap/4 {
		d.resize(d.Len())
	}
	return result
}

func (d *Deque[T]) Get(i int) T {
	if i < 0 || i >= d.Len() {
		panic("Get() uses invalid index!")
	}
	return d.buffer[(d.head+i)&(d.cap-1)]
}

func (d *Deque[T]) Set(i int, v T) {
	if i < 0 || i >= d.Len() {
		panic("Set() uses invalid index!")
	}
	d.buffer[(d.head+i)&(d.cap-1)] = v
}

// Len returns the number of elements of deque.
func (d *Deque[T]) Len() int {
	return (d.tail + d.cap - d.head) & (d.cap - 1)
}

func (d *Deque[T]) Cap() int {
	return d.cap
}

func (d *Deque[T]) prev(pos int) int {
	return (pos + d.cap - 1) & (d.cap - 1)
}

func (d *Deque[T]) next(pos int) int {
	return (pos + 1) & (d.cap - 1)
}

func lowPower(cap int) int {
	n := cap - 1
	n = n | (n >> 1)
	n = n | (n >> 2)
	n = n | (n >> 4)
	n = n | n>>8
	n = n | n>>16

	return n + 1
}
