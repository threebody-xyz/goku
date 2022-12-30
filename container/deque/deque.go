package deque

type Element interface {}

const minCapacity = 16

type Deque struct {
	buffer []Element
	head int
	tail int
	len int
}

// NewDeque params[0] refer capacity
func NewDeque(params ...int) *Deque{
	var capacity int
	if len(params) > 0 {
		capacity = params[0]
	} else {
		capacity = minCapacity
	}
	buffer := make([]Element, capacity)
	return &Deque{
		buffer: buffer,
	}
}

func (d *Deque) resize() {
	newBuffer := make([]Element, d.len<<1)
	if d.tail > d.head {
		copy(newBuffer, d.buffer[d.head:d.tail])
	} else {
		n := copy(newBuffer, d.buffer[d.head:])
		copy(newBuffer[n:], d.buffer[:d.tail])
	}
	d.head = 0
	d.tail = d.len
	d.buffer = newBuffer
}

// PushFront inserts a new element e with value v at the front of deque d.
func (d *Deque) PushFront(v Element)  {
	if len(d.buffer) != 0 && d.len == len(d.buffer) {
		d.resize()
	}
	d.head = (d.head - 1) & (len(d.buffer)-1)
	d.buffer[d.head] = v
	d.len++
}

// PushBack inserts a new element e with value v at the back of deque d.
func (d *Deque) PushBack(v Element)  {
	if len(d.buffer) != 0 && d.len == len(d.buffer) {
		d.resize()
	}
	d.tail = (d.tail+1) & len(d.buffer)
	d.buffer[d.tail] = v
	d.len++
}
// PopFront returns the first element of deque d or nil if the deque is empty.
func (d *Deque) PopFront() Element {
	if d.len <= 0{
		return nil
	}
	result := d.buffer[d.head]
	var zero Element
	d.buffer[d.head] = zero
	d.head =(d.head+1) % len(d.buffer)
	d.len --
	if len(d.buffer) > minCapacity && len(d.buffer) == d.len<<2 {
		d.resize()
	}
	return result
}

// PopBack returns the last element of deque d or nil if the deque is empty.
func (d *Deque) PopBack() Element {
	if d.len <= 0{
		return nil
	}
	result := d.buffer[d.tail]
	var zero Element
	d.buffer[d.tail] = zero
	d.tail =(d.tail-1) & (len(d.buffer)-1)
	d.len --
	if len(d.buffer) > minCapacity && len(d.buffer) == d.len<<2 {
		d.resize()
	}
	return result
}

func (d *Deque) Get(i int) Element {
	if i < 0 || i >= d.len {
		return nil
	}
	return d.buffer[(d.head+1)%(len(d.buffer))]
}

func (d *Deque) Set(i int, v Element) {
	if i < 0 || i >= d.len {
		return
	}
	d.buffer[(d.head+1)%(len(d.buffer))] = v
}

// Len returns the number of elements of deque d.
func (d *Deque) Len() int {
	return d.len
}
