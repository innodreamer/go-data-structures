/*
Package minheap provides an implementation of a Min Heap data structure
*/
package minheap

import (
	"log"

	"github.com/dorin131/go-data-structures/heap"
)

// MinHeap : represents a Min Heap data structure
type MinHeap struct {
	*heap.Heap
}

// New : returns a new instance of a Heap
func New(input []int) *MinHeap {
	h := &MinHeap{
		&heap.Heap{
			Size:     len(input),
			Elements: len(input),
			Items:    input,
		},
	}

	if len(h.Items) > 0 {
		h.buildMinHeap()
	}

	return h
}

// ExtractMin : removes top element of the heap and returns it
func (h *MinHeap) ExtractMin() int {
	if len(h.Items) == 0 {
		log.Fatal("No items in the heap")
	}
	item := h.Items[0]
	h.Items[0] = h.Items[len(h.Items)-1]
	h.Items = h.Items[:len(h.Items)-1]

	h.minHeapifyDown()

	return item
}

// Insert : adds a new element to the heap
func (h *MinHeap) Insert(item int) *MinHeap {
	h.Items = append(h.Items, item)
	lastElementIndex := len(h.Items) - 1
	h.minHeapifyUp(lastElementIndex)

	return h
}

func (h *MinHeap) buildMinHeap() {
	for i := len(h.Items) / 2; i >= 0; i-- {
		h.minHeapifyUp(i)
	}
}

// minHeapifyDown : takes the top element and moves it down until
// the min-heap properties are satisfied
func (h *MinHeap) minHeapifyDown() {
	index := 0
	for (h.HasLeft(index) && h.Items[index] > h.Left(index)) ||
		(h.HasRight(index) && h.Items[index] > h.Right(index)) {
		if h.HasLeft(index) && h.Items[index] > h.Left(index) {
			h.Swap(index, h.GetLeftIndex(index))
			index = h.GetLeftIndex(index)
		} else {
			h.Swap(index, h.GetRightIndex(index))
			index = h.GetRightIndex(index)
		}
	}
}

// minHeapUp : takes the last element and moves it up until
// the min-heap properties are satisfied
func (h *MinHeap) minHeapifyUp(index int) {
	for h.HasParent(index) && h.Parent(index) > h.Items[index] {
		h.Swap(h.GetParentIndex(index), index)
		index = h.GetParentIndex(index)
	}
}
