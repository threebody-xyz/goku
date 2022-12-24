package deque

type Element interface{}

type Deque struct {
}

// PushFront inserts a new element e with value v at the front of deque d.
func (d *Deque) PushFront(v Element) {
}

// PushBack inserts a new element e with value v at the back of deque d.
func (d *Deque) PushBack(v Element) {
}

// PopFront
func (d *Deque) PopFront() Element {
	return nil
}

// PopBack
func (d *Deque) PopBack() Element {
	return nil
}

// Front returns the first element of deque d or nil if the deque is empty.
func (d *Deque) Front() Element {
	return nil
}

// Back returns the last element of deque d or nil if the deque is empty.
func (d *Deque) Back() Element {
	return nil
}

func (d *Deque) Get(i int) Element {
	return nil
}

func (d *Deque) Set(i int, v Element) {
}

// Len returns the number of elements of deque d.
func (d *Deque) Len() int {
	return 0
}

func (d *Deque) Less(i, j int) bool {
	return false
}

// Swap swaps the elements with indexes i and j.
func (d *Deque) Swap(i, j int) {

}
