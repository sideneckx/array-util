package array_util

func Map[P any, T any](list []P, closure func(P) T) []T {
	result := make([]T, len(list))
	for i, item := range list {
		result[i] = closure(item)
	}
	return result
}

func Reduce[P any, T any](initial T, list []P, closure func(T, P) T) T {
	result := initial
	for _, item := range list {
		result = closure(result, item)
	}
	return result
}

func Filter[T any](list []T, closure func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range list {
		if closure(item) {
			result = append(result, item)
		}
	}
	return result
}

func FirstIndex[T any](list []T, closure func(T) bool) int {
	for i, item := range list {
		if closure(item) {
			return i
		}
	}
	return -1
}

func FirstIndexOfItem[T any](list []T, item *T) int {
	for i := 0; i < len(list); i++ {
		if &list[i] == item {
			return i
		}
	}
	return -1
}

func First[T any](list []T, closure func(T) bool) *T {
	for i := 0; i < len(list); i++ {
		if closure(list[i]) {
			return &list[i]
		}
	}
	return nil
}
