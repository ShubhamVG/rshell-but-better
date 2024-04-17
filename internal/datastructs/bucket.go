package datastructs

// =====================EXPORTABLES=====================

type Bucket[T any] struct {
	Len int

	capacity     int
	currEmptyIdx int
	items        []T
}

func NewBucket[T any](capacity int) Bucket[T] {
	storage := make([]T, capacity)

	return Bucket[T]{capacity: capacity, items: storage, Len: 0, currEmptyIdx: 0}
}

func (bkt *Bucket[T]) Append(item T) {
	if bkt.Len == bkt.capacity {
		bkt.tossHalf()
		return
	}

	bkt.items[bkt.currEmptyIdx] = item
	bkt.currEmptyIdx++
}

// Returns bkt.items[:Len] so DO NOT EDIT IT
func (bkt *Bucket[T]) Items() []T {
	return bkt.items[:bkt.Len]
}

// ====================INTERNALS==========================

// Make new slice and copy the last half items to it
// And make the new slice the bucket's storage
func (bkt *Bucket[T]) tossHalf() {
	newStorage := make([]T, bkt.capacity)
	halfCapacity := bkt.capacity / 2

	// if odd, mimic ceil division to make life easier
	if (bkt.capacity & 1) == 1 {
		halfCapacity++
	}

	for i := halfCapacity; i < bkt.capacity; i++ {
		newStorage[i-halfCapacity] = bkt.items[i]
	}

	bkt.items = newStorage
	bkt.currEmptyIdx = bkt.capacity - halfCapacity
}
