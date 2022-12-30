package deque

import (
	"fmt"
	"testing"
)

func TestDeque(t *testing.T){
	deque := NewDeque()
	for i:=0; i<minCapacity+1; i++ {
		deque.PushFront(i)
	}
	fmt.Printf("deque: %+v\n", deque)

	for i:=0; i<minCapacity+1; i++ {
		fmt.Println(deque.PopFront())
	}
	fmt.Printf("deque: %+v\n", deque)

	for i:=0; i<minCapacity+1; i++ {
		deque.PushBack(i)
	}
	for i:=0; i<minCapacity+1; i++ {
		fmt.Println(deque.PopBack())
	}
	for i:=0; i<(minCapacity+1)/2; i++ {
		deque.PushFront(i)
	}
	for i:=0; i<(minCapacity+1)/2; i++ {
		deque.PushBack(i)
	}
	fmt.Printf("deque: %+v\n", deque)
	dequeItem0 := deque.Get(0)
	fmt.Println(dequeItem0)
	deque.Set(0, 999)
	dequeItem0 = deque.Get(0)
	fmt.Println(dequeItem0)
}

func BenchmarkPushFront(b *testing.B) {
	deque := NewDeque()
	for i:=0; i<b.N; i++ {
		deque.PushFront(i)
	}
}

func BenchmarkPushBack(b *testing.B) {
	deque := NewDeque()
	for i:=0; i<b.N; i++ {
		deque.PushBack(i)
	}
}

func BenchmarkPushPopBack(b *testing.B) {
	deque := NewDeque()
	for i:=0; i<b.N; i++ {
		deque.PushBack(i)
		deque.PopBack()
	}
}
func BenchmarkPushPopFront(b *testing.B) {
	deque := NewDeque()
	for i:=0; i<b.N; i++ {
		deque.PushFront(i)
		deque.PopFront()
	}
}