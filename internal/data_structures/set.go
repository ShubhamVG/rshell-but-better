package datastructures

type Nothing struct{}

type Set[T comparable] struct {
	hashMap map[T]Nothing
}

func NewSet[T comparable]() Set[T] {
	hashMap := map[T]Nothing{}
	return Set[T]{hashMap: hashMap}
}

// Non-destuctive
func NewSetFromArray[T comparable](arr []T) Set[T] {
	set := NewSet[T]()

	for _, data := range arr {
		set.Add(data)
	}

	return set
}

// Destructive
func NewSetFromQueue[T comparable](q Queue[T]) Set[T] {
	set := NewSet[T]()
	first := q.First

	for first != nil {
		set.Add(first.Data)
		first = first.Next
	}

	return set
}

func (set *Set[T]) Add(elem T) {
	set.hashMap[elem] = Nothing{}
}

func (set *Set[T]) Contains(elem T) bool {
	if _, isPresent := set.hashMap[elem]; isPresent {
		return true
	}

	return false
}
