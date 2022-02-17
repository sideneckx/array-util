package array_util_test

import (
	"strconv"
	"testing"

	array "github.com/sideneckx/array_util"
)

func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	output := array.Map(input, func(item int) string {
		return strconv.Itoa(item)
	})
	if output[0] == "1" && output[1] == "2" && output[2] == "3" {
		t.Log("Map test passed")
	} else {
		t.Error("Map test failed")
	}
}

func TestReduce(t *testing.T) {
	input := []int{1, 2, 3}
	output := array.Reduce(0, input, func(result int, item int) int {
		return result + item
	})
	if output == 6 {
		t.Log("Reduce test passed")
	} else {
		t.Error("Reduce test failed")
	}
}

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	output := array.Filter(input, func(item int) bool {
		return item > 3
	})
	if output[0] == 4 && output[1] == 5 {
		t.Log("Filter test passed")
	} else {
		t.Error("Filter test failed")
	}
}

func TestFirstIndex(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	output := array.FirstIndex(input, func(item int) bool {
		return item == 3
	})
	if output == 2 {
		t.Log("FirstIndex test passed")
	} else {
		t.Error("FirstIndex test failed")
	}
}

func TestFirstIndexOfItem(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	output := array.FirstIndexOfItem(input, &input[2])
	if output == 2 {
		t.Log("FirstIndexOfItem test passed")
	} else {
		t.Error("FirstIndexOfItem test failed")
	}
}

func TestFirst(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	output := array.First(input, func(item int) bool {
		return item == 3
	})
	if *output == 3 {
		t.Log("First test passed")
	} else {
		t.Error("First test failed")
	}
}

func TestInitSet(t *testing.T) {
	set := array.NewSet([]int{1, 2, 3})
	if set.Size() == 3 {
		t.Log("InitSet test passed")
	} else {
		t.Error("InitSet test failed")
	}
}
func TestInitSet2(t *testing.T) {
	set := array.Set[int]{}
	if set.Size() == 0 {
		t.Log("InitSet test passed")
	} else {
		t.Error("InitSet test failed")
	}
}
func TestSetContains(t *testing.T) {
	set := array.NewSet([]int{1, 2, 3})
	if set.Contains(2) {
		t.Log("SetContains test passed")
	} else {
		t.Error("SetContains test failed")
	}
}
func TestInsetSet(t *testing.T) {
	set := array.NewSet([]int{1, 2, 3})
	set.Insert(4)
	if set.Size() == 4 && set.Contains(4) {
		t.Log("InsetSet test passed")
	} else {
		t.Error("InsetSet test failed")
	}
}
func TestRemoveSet(t *testing.T) {
	set := array.NewSet([]int{1, 2, 3})
	set.Remove(2)
	if set.Size() == 2 && !set.Contains(2) {
		t.Log("RemoveSet test passed")
	} else {
		t.Error("RemoveSet test failed")
	}
}

func TestFilterSet(t *testing.T) {
	set := array.NewSet([]int{1, 2, 3})
	result := set.Filter(func(item int) bool {
		return item > 1
	})
	if result.Size() == 2 && result.Contains(3) && result.Contains(2) {
		t.Log("FilterSet test passed")
	} else {
		t.Error("FilterSet test failed")
	}
}

func TestUnionSet(t *testing.T) {
	set1 := array.NewSet([]int{1, 2, 3})
	set2 := array.NewSet([]int{2, 3, 4})
	result := set1.Union(*set2)
	if result.Size() == 4 && result.Contains(1) && result.Contains(2) && result.Contains(3) && result.Contains(4) && (*set1).Size() == 3 {
		t.Log("UnionSet test passed")
	} else {
		t.Error("UnionSet test failed")
	}
}
