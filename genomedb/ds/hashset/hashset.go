package hashset

type HashSet[T comparable] map[T]struct{}

func NewHashSet[T comparable]() HashSet[T] {
	return make(HashSet[T])
}

func (s HashSet[T]) Contains(value T) bool {
	_, found := s[value]
	return found
}

func (s *HashSet[T]) Remove(value T) {
	delete((*s), value)
}

func (s *HashSet[T]) Add(value T) {
	(*s)[value] = struct{}{}
}
