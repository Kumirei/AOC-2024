package util

import (
	"container/heap"
)

// An MHItem is something we manage in a priority queue.
type MHItem[T any] struct {
	Value    T   // The value of the item; arbitrary.
	Priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A MinHeap implements heap.Interface and holds Items.
type MinHeap[T any] []*MHItem[T]

func (pq MinHeap[T]) Len() int { return len(pq) }

func (pq MinHeap[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq MinHeap[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *MinHeap[T]) Push(x any) {
	n := len(*pq)
	item := x.(*MHItem[T])
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *MinHeap[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an MHItem in the queue.
func (pq *MinHeap[T]) update(item *MHItem[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
