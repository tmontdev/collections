package iterable

import (
	"fmt"
	"testing"
)

type testCase[T comparable] struct {
	name              string
	input             Iterable[any]
	parameters        []any
	expectPanic       bool
	expected          T
	runnable          func(*testing.T, Iterable[any], []any) T
	nilTypeComparison bool
}

var emptyList = NewList[any]()

var oneTwoThreeList = NewList[any](1, 2, 3)

func CaseRunner[T comparable](t *testing.T, c testCase[T]) {
	println("CaseRunner: Running test case: ", c.name, "...")
	panicked := false
	defer func() {
		if r := recover(); r != nil {
			if c.expectPanic {
				panicked = true
				fmt.Printf("CaseRunner: test case %v panics as expected. %v\n", c.name, r)
			}
			return
		}
	}()
	result := c.runnable(t, c.input, c.parameters)
	if c.expectPanic != panicked {
		t.Errorf("CaseRunner: %v: Expected panic %v. Got: %v", c.name, c.expectPanic, panicked)
	}
	if !c.nilTypeComparison && fmt.Sprintf("%v%v", result, c.expected) == "<nil><nil>" {
		return
	}
	if result != c.expected && !c.expectPanic {
		t.Errorf("CaseRunner: %v: Expected %v. Got: %v", c.name, c.expected, result)
	}
	return
}

var lengthCases = []testCase[int]{
	{
		name:        "List.Length.Empty",
		input:       emptyList.Clone(),
		parameters:  nil,
		expectPanic: false,
		expected:    0,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) int {
			return list.Length()
		},
	},
	{
		name:        "List.Length.Value",
		input:       oneTwoThreeList.Clone(),
		parameters:  nil,
		expectPanic: false,
		expected:    len(oneTwoThreeList.Elements()),
		runnable: func(t *testing.T, list Iterable[any], parameters []any) int {
			return list.Length()
		},
	},
}

var emptyCases = []testCase[bool]{
	{
		name:        "List.IsEmpty.Empty",
		input:       emptyList.Clone(),
		parameters:  nil,
		expectPanic: false,
		expected:    true,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.IsEmpty()
		},
	},
	{
		name:        "List.IsNotEmpty.Empty",
		input:       emptyList.Clone(),
		parameters:  nil,
		expectPanic: false,
		expected:    false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.IsNotEmpty()
		},
	},
	{
		name:        "List.IsNotEmpty.Filled",
		input:       oneTwoThreeList.Clone(),
		parameters:  nil,
		expectPanic: false,
		expected:    true,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.IsNotEmpty()
		},
	},
	{
		name:        "List.IsEmpty.Filled",
		input:       oneTwoThreeList.Clone(),
		parameters:  nil,
		expectPanic: false,
		expected:    false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.IsEmpty()
		},
	},
}

var atCases = []testCase[any]{
	{
		name:              "List.At.Empty",
		input:             emptyList.Clone(),
		parameters:        nil,
		expectPanic:       false,
		expected:          nil,
		nilTypeComparison: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) any {
			return list.At(0)
		},
	},
	{
		name:              "List.At.MutableValue",
		input:             oneTwoThreeList.Clone(),
		parameters:        nil,
		expectPanic:       false,
		expected:          true,
		nilTypeComparison: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) any {
			n := list.At(0)
			*n = 4
			return list.ElementAt(0).(int) == 4
		},
	},
	{
		name:              "List.ElementAt.Empty",
		input:             emptyList.Clone(),
		parameters:        nil,
		expectPanic:       true,
		expected:          false,
		nilTypeComparison: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) any {
			return list.ElementAt(0).(int) == 1
		},
	},
	{
		name:              "List.ElementAt.NonMutableValue",
		input:             oneTwoThreeList.Clone(),
		parameters:        nil,
		expectPanic:       false,
		expected:          true,
		nilTypeComparison: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) any {
			n := list.ElementAt(0)
			if n != 1 {
				t.Error("List.ElementAt is different from expected")
			}
			n = 4
			a := list.ElementAt(0)
			if a != 1 {
				t.Error("List.ElementAt is different from expected")
			}
			return a != n
		},
	},
}

var indexCases = []testCase[bool]{
	{
		name:        "List.FirstElement.Empty",
		input:       emptyList.Clone(),
		expected:    false,
		expectPanic: true,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.FirstElement() == 1
		},
	},
	{
		name:        "List.First.Empty",
		input:       emptyList.Clone(),
		expected:    false,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.First() == nil
		},
	},
	{
		name:        "List.LastElement.Empty",
		input:       emptyList.Clone(),
		expected:    false,
		expectPanic: true,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.LastElement() == 1
		},
	},
	{
		name:        "List.Last.Empty",
		input:       emptyList.Clone(),
		expected:    false,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.Last() == nil
		},
	},
	{
		name:        "List.FirstIndexWhere.NotSatisfied",
		input:       emptyList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.FirstIndexWhere(func(a any) bool {
				return true
			}) == -1
		},
	},
	{
		name:        "List.FirstIndexWhere.Even",
		input:       emptyList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.FirstIndexWhere(func(a any) bool {
				return a.(int)%2 == 0
			}) == 1
		},
	},
	{
		name:        "List.FirstIndexWhere.Odd",
		input:       emptyList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.FirstIndexWhere(func(a any) bool {
				return a.(int)%2 == 0
			}) == 1
		},
	},
}

func TestLength(t *testing.T) {
	for _, v := range lengthCases {
		CaseRunner[int](t, v)
	}
}

func TestEmpty(t *testing.T) {
	for _, v := range emptyCases {
		CaseRunner[bool](t, v)
	}
}

func TestAt(t *testing.T) {
	for _, v := range atCases {
		CaseRunner[any](t, v)
	}
}

func TestIndexes(t *testing.T) {
	for _, v := range indexCases {
		CaseRunner[bool](t, v)
	}
}
