package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDequeBasicInt(t *testing.T) {
	d := NewDeque[int]()

	d.PushFront(2, 1)
	assert.Equal(t, 2, d.Len())
	assert.Equal(t, minCapacity, d.Cap())

	d.PushBack(3, 4)
	assert.Equal(t, 4, d.Len())
	assert.Equal(t, minCapacity, d.Cap())

	assert.Equal(t, 1, d.Get(0))
	assert.Equal(t, 2, d.Get(1))
	assert.Equal(t, 3, d.Get(2))
	assert.Equal(t, 4, d.Get(3))

	assert.Equal(t, 1, d.PopFront())
	assert.Equal(t, 3, d.Len())
	assert.Equal(t, minCapacity, d.Cap())

	assert.Equal(t, 4, d.PopBack())
	assert.Equal(t, 2, d.Len())
	assert.Equal(t, minCapacity, d.Cap())

	d.Set(0, 3)
	d.Set(1, 2)
	assert.Equal(t, 3, d.Get(0))
	assert.Equal(t, 2, d.Get(1))
}

func TestDequeBasicType(t *testing.T) {
	type User struct {
		name string
	}
	d := NewDeque[User]()
	d.PushFront(User{"1"}, User{"2"})
	assert.Equal(t, 2, d.Len())
	assert.Equal(t, minCapacity, d.Cap())
}

func TestDequeWithCap(t *testing.T) {
	d := NewDequeWithCap[int](20)
	assert.Equal(t, minCapacity*2, d.Cap())

	d = NewDequeWithCap[int](8)
	assert.Equal(t, minCapacity, d.Cap())
}

func TestDequeResize(t *testing.T) {
	slice := make([]int, 0)
	for i := 0; i < 64; i++ {
		slice = append(slice, i)
	}

	deque := NewDeque[int]()

	deque.PushBack(slice...)
	assert.Equal(t, 128, deque.Cap())
	deque.PushFront(slice...)
	assert.Equal(t, 256, deque.Cap())

	for i := 0; i < 64; i++ {
		deque.PopFront()
	}
	assert.Equal(t, 128, deque.Cap())
	for i := 0; i < 32; i++ {
		deque.PopBack()
	}
	assert.Equal(t, 64, deque.Cap())

	deque = NewDeque[int]()

	deque.PushBack(-1, 0, 1, 2, 3, 4)
	deque.PopFront()
	deque.resize(50)
	assert.Equal(t, 128, deque.Cap())
	for i := 0; i < 5; i++ {
		assert.Equal(t, i, deque.Get(i))
	}

	deque.PushFront(-1, -2, -3, -4, -5)
	deque.resize(10)
	assert.Equal(t, 32, deque.Cap())
	for i := 0; i < 10; i++ {
		assert.Equal(t, i-5, deque.Get(i))
	}

}

func TestPanic(t *testing.T) {
	t.Run("PopFrontPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("PopFront() panic!")
			}
		}()
		deque := NewDeque[int]()
		deque.PopFront()
	})

	t.Run("PopBackPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("PushBack() panic!")
			}
		}()
		deque := NewDeque[int]()
		deque.PopBack()
	})

	t.Run("GetPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Get() panic!")
			}
		}()
		deque := NewDeque[int]()
		deque.Get(1)
	})

	t.Run("SetPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Set() panic!")
			}
		}()
		deque := NewDeque[int]()
		deque.Set(1, 0)
	})
}

func TestLowPower(t *testing.T) {
	assert.Equal(t, 1, lowPower(1))
	assert.Equal(t, 2, lowPower(2))
	assert.Equal(t, 4, lowPower(3))
}

func BenchmarkPushFront(b *testing.B) {
	deque := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		deque.PushFront(i)
	}
}

func BenchmarkPushBack(b *testing.B) {
	deque := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		deque.PushBack(i)
	}
}

func BenchmarkPushPopBack(b *testing.B) {
	deque := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		deque.PushBack(i)
		deque.PopBack()
	}
}

func BenchmarkPushPopFront(b *testing.B) {
	deque := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		deque.PushFront(i)
		deque.PopFront()
	}
}
