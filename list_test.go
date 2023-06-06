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

func caseRunner[T comparable](t *testing.T, c testCase[T]) {
	println("caseRunner: Running test case: ", c.name, "...")
	panicked := false
	defer func() {
		if r := recover(); r != nil {
			if c.expectPanic {
				panicked = true
				fmt.Printf("caseRunner: test case %v panics as expected. %v\n", c.name, r)
			}
			return
		}
	}()
	result := c.runnable(t, c.input, c.parameters)
	if c.expectPanic != panicked {
		t.Errorf("caseRunner: %v: Expected panic %v. Got: %v", c.name, c.expectPanic, panicked)
	}
	if !c.nilTypeComparison && fmt.Sprintf("%v%v", result, c.expected) == "<nil><nil>" {
		return
	}
	if result != c.expected && !c.expectPanic {
		t.Errorf("caseRunner: %v: Expected %v. Got: %v", c.name, c.expected, result)
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
		input:       oneTwoThreeList.Clone(),
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
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.FirstIndexWhere(func(a any) bool {
				return a.(int)%2 == 1
			}) == 0
		},
	},
	{
		name:        "List.LastIndexWhere.Odd",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.LastIndexWhere(func(a any) bool {
				return a.(int)%2 == 1
			}) == 2
		},
	},
	{
		name:        "List.IndexWhere.Odd",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			odds := list.IndexWhere(func(a any) bool {
				return a.(int)%2 == 1
			})
			return odds.ElementAt(0) == 0 && odds.ElementAt(1) == 2
		},
	},
}
var whereCases = []testCase[bool]{
	{
		name:        "List.Where.Odd",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			odds := list.Where(func(candidate any) bool {
				return candidate.(int)%2 == 1
			})
			return odds.ElementAt(0).(int) == 1 && odds.ElementAt(1).(int) == 3 && odds.Length() == 2
		},
	},
	{
		name:        "List.Where.Even",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			odds := list.Where(func(candidate any) bool {
				return candidate.(int)%2 == 0
			})
			return odds.ElementAt(0).(int) == 2 && odds.Length() == 1
		},
	},
}
var mapCases = []testCase[bool]{
	{
		name:        "List.Map.String",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			strs := list.Map(func(candidate any) any {
				return fmt.Sprint(candidate)
			})
			return strs.ElementAt(0).(string) == "1" && strs.ElementAt(1).(string) == "2" && strs.ElementAt(2).(string) == "3" && strs.Length() == 3
		},
	},
}

var reduceCases = []testCase[bool]{
	{
		name:        "List.Reduce.Sum",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.Reduce(func(acc any, candidate any, idx int) any {
				return acc.(int) + candidate.(int)
			}, 0) == 6
		},
	},
	{
		name:        "List.Reduce.Map",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			strs := list.Reduce(func(acc any, candidate any, idx int) any {
				return acc.(*List[string]).Push(fmt.Sprint(candidate))
			}, NewList[string]()).(*List[string])
			return strs.ElementAt(0) == "1" && strs.ElementAt(1) == "2" && strs.ElementAt(2) == "3" && strs.Length() == 3
		},
	},
}

var everyCases = []testCase[bool]{
	{
		name:        "List.Every.lt",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.Every(func(candidate any) bool {
				return candidate.(int) < 4
			})
		},
	},
	{
		name:        "List.Every.gt",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.Every(func(candidate any) bool {
				return candidate.(int) > 0
			})
		},
	},
	{
		name:        "List.Every.gt.false",
		input:       oneTwoThreeList.Clone(),
		expected:    false,
		expectPanic: false,
		runnable: func(t *testing.T, list Iterable[any], parameters []any) bool {
			return list.Every(func(candidate any) bool {
				return candidate.(int) > 1
			})
		},
	},
}

func TestLength(t *testing.T) {
	for _, v := range lengthCases {
		caseRunner[int](t, v)
	}
}

func TestEmpty(t *testing.T) {
	for _, v := range emptyCases {
		caseRunner[bool](t, v)
	}
}

func TestAt(t *testing.T) {
	for _, v := range atCases {
		caseRunner[any](t, v)
	}
}

func TestIndexes(t *testing.T) {
	for _, v := range indexCases {
		caseRunner[bool](t, v)
	}
}

func TestWhere(t *testing.T) {
	for _, v := range whereCases {
		caseRunner[bool](t, v)
	}
}

func TestMap(t *testing.T) {
	for _, v := range mapCases {
		caseRunner[bool](t, v)
	}
}

func TestReduce(t *testing.T) {
	for _, v := range reduceCases {
		caseRunner[bool](t, v)
	}
}

func TestEvery(t *testing.T) {
	for _, v := range everyCases {
		caseRunner[bool](t, v)
	}
}
