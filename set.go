package array_util

type Set[T comparable] struct {
	items map[T]bool
}

func NewSet[T comparable](array []T) *Set[T] {
	var item = map[T]bool{}
	for i := 0; i < len(array); i++ {
		item[array[i]] = true
	}
	return &Set[T]{items: item}
}

func (s *Set[T]) Insert(item T) {
	s.items[item] = true
}
func (s Set[T]) Remove(item T) {
	delete(s.items, item)
}
func (s Set[T]) Contains(item T) bool {
	return s.items[item]
}
func (s Set[T]) Size() int {
	return len(s.items)
}
func (s Set[T]) Items() []T {
	var items = []T{}
	for key, value := range s.items {
		if value {
			items = append(items, key)
		}
	}
	return items
}

func (s Set[T]) Filter(closure func(T) bool) Set[T] {
	var result = NewSet([]T{})
	for key, value := range s.items {
		if value {
			if closure(key) {
				result.Insert(key)
			}
		}
	}
	return *result
}

//combine two set
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := NewSet([]T{})
	for key, value := range s.items {
		if value {
			result.Insert(key)
		}
	}
	for key, value := range other.items {
		if value {
			result.Insert(key)
		}
	}
	return *result
}
