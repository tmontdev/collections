package collection

import (
	"fmt"
	"testing"
)

type testCase[T comparable] struct {
	name              string
	input             Collection[any]
	parameters        []any
	expectPanic       bool
	expected          T
	runnable          func(*testing.T, Collection[any], []any) T
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
				message := fmt.Sprintf("caseRunner: test case %v panics as expected. %v\n", c.name, r)
				println(message)
			} else {
				message := fmt.Sprintf(fmt.Sprintf("caseRunner: test case %v panics were not expected. %v\n", c.name, r))
				println(message)
				t.Error(message)
			}
			return
		}
	}()
	result := c.runnable(t, c.input, c.parameters)
	if c.expectPanic != panicked {
		message := fmt.Sprintf("caseRunner: %v: Expected panic %v. Got: %v", c.name, c.expectPanic, panicked)
		println(message)
		t.Error(message)
	}
	if !c.nilTypeComparison && fmt.Sprintf("%v%v", result, c.expected) == "<nil><nil>" {
		return
	}
	if result != c.expected && !c.expectPanic {
		message := fmt.Sprintf("caseRunner: %v: Expected %v. Got: %v", c.name, c.expected, result)
		println(message)
		t.Error(message)
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) int {
			return list.Length()
		},
	},
	{
		name:        "List.Length.Value",
		input:       oneTwoThreeList.Clone(),
		parameters:  nil,
		expectPanic: false,
		expected:    len(oneTwoThreeList.Elements()),
		runnable: func(t *testing.T, list Collection[any], parameters []any) int {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.IsEmpty()
		},
	},
	{
		name:        "List.IsNotEmpty.Empty",
		input:       emptyList.Clone(),
		parameters:  nil,
		expectPanic: false,
		expected:    false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.IsNotEmpty()
		},
	},
	{
		name:        "List.IsNotEmpty.Filled",
		input:       oneTwoThreeList.Clone(),
		parameters:  nil,
		expectPanic: false,
		expected:    true,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.IsNotEmpty()
		},
	},
	{
		name:        "List.IsEmpty.Filled",
		input:       oneTwoThreeList.Clone(),
		parameters:  nil,
		expectPanic: false,
		expected:    false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) any {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) any {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) any {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) any {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.FirstElement() == 1
		},
	},
	{
		name:        "List.First.Empty",
		input:       emptyList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.First() == nil
		},
	},
	{
		name:        "List.LastElement.Empty",
		input:       emptyList.Clone(),
		expected:    false,
		expectPanic: true,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.LastElement() == 1
		},
	},
	{
		name:        "List.Last.Empty",
		input:       emptyList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.Last() == nil
		},
	},
	{
		name:        "List.FirstIndexWhere.NotSatisfied",
		input:       emptyList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
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
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.Every(func(candidate any) bool {
				return candidate.(int) > 1
			})
		},
	},
}

var someCases = []testCase[bool]{
	{
		name:        "List.Some.lt",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.Some(func(candidate any) bool {
				return candidate.(int) < 4
			})
		},
	},
	{
		name:        "List.Some.gt",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.Some(func(candidate any) bool {
				return candidate.(int) > 0
			})
		},
	},
	{
		name:        "List.Every.eq",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.Some(func(candidate any) bool {
				return candidate.(int) == 1
			})
		},
	},

	{
		name:        "List.Every.eq.false",
		input:       oneTwoThreeList.Clone(),
		expected:    false,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.Some(func(candidate any) bool {
				return candidate.(int) == 10
			})
		},
	},
}

var noneCases = []testCase[bool]{
	{
		name:        "List.None.lt",
		input:       oneTwoThreeList.Clone(),
		expected:    false,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.None(func(candidate any) bool {
				return candidate.(int) < 4
			})
		},
	},
	{
		name:        "List.None.gt",
		input:       oneTwoThreeList.Clone(),
		expected:    false,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.None(func(candidate any) bool {
				return candidate.(int) > 0
			})
		},
	},
	{
		name:        "List.None.eq",
		input:       oneTwoThreeList.Clone(),
		expected:    false,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.None(func(candidate any) bool {
				return candidate.(int) == 1
			})
		},
	},

	{
		name:        "List.None.eq.false",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			return list.None(func(candidate any) bool {
				return candidate.(int) == 10
			})
		},
	},
}

var popCases = []testCase[bool]{
	{
		name:        "List.Pop",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			popped := list.Pop()
			return popped.ElementAt(0) == 1 && popped.ElementAt(1) == 2 && popped.Length() == 2

		},
	},
	{
		name:        "List.Pop.Double",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			popped := list.Pop().Pop()
			return popped.ElementAt(0) == 1 && popped.Length() == 1

		},
	},
}

var shiftCases = []testCase[bool]{
	{
		name:        "List.Shift",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			popped := list.Shift()
			return popped.ElementAt(0) == 2 && popped.ElementAt(1) == 3 && popped.Length() == 2

		},
	},
	{
		name:        "List.Shift.Double",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			popped := list.Shift().Shift()
			return popped.ElementAt(0) == 3 && popped.Length() == 1

		},
	},
}

var setCases = []testCase[bool]{
	{
		name:        "List.Set",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			set := list.Set(1, 4)
			return set.Length() == list.Length() && set.ElementAt(0) == list.ElementAt(0) && set.ElementAt(1) == 4 && set.ElementAt(2) == list.ElementAt(2)

		},
	},
}

var stringCases = []testCase[string]{
	{
		name:        "List.String",
		input:       oneTwoThreeList.Clone(),
		expected:    fmt.Sprint(oneTwoThreeList.Elements()),
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) string {
			return list.String()
		},
	},
}

var intervalCases = []testCase[bool]{
	{
		name:        "List.Interval",
		input:       oneTwoThreeList.Clone(),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) bool {
			interval := list.Interval(0, 1)
			return interval.ElementAt(0).(int) == 1 && interval.ElementAt(1).(int) == 2
		},
	},
}

var sortCases = []testCase[string]{
	{
		name:        "List.Sort.Int.Esc",
		input:       NewList[any](5, 1981, 3, 15, 142, 1, 23, 89, 67, 1203, 439, 24),
		expected:    fmt.Sprint([]any{1, 3, 5, 15, 23, 24, 67, 89, 142, 439, 1203, 1981}),
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) string {
			list.Sort(func(a, b any) int {
				return a.(int) - b.(int)
			})
			return list.String()
		},
	},
	{
		name:        "List.Sort.Int.Desc",
		input:       NewList[any](5, 1981, 3, 15, 142, 1, 23, 89, 67, 1203, 439, 24),
		expected:    fmt.Sprint([]any{1981, 1203, 439, 142, 89, 67, 24, 23, 15, 5, 3, 1}),
		expectPanic: false,
		runnable: func(t *testing.T, list Collection[any], parameters []any) string {
			list.Sort(func(a, b any) int {
				return b.(int) - a.(int)
			})
			return list.String()
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

func TestSome(t *testing.T) {
	for _, v := range someCases {
		caseRunner[bool](t, v)
	}
}

func TestNone(t *testing.T) {
	for _, v := range noneCases {
		caseRunner[bool](t, v)
	}
}

func TestPop(t *testing.T) {
	for _, v := range popCases {
		caseRunner[bool](t, v)
	}
}

func TestShift(t *testing.T) {
	for _, v := range shiftCases {
		caseRunner[bool](t, v)
	}
}

func TestString(t *testing.T) {
	for _, v := range stringCases {
		caseRunner[string](t, v)
	}
}

func TestSet(t *testing.T) {
	for _, v := range setCases {
		caseRunner[bool](t, v)
	}
}

func TestInterval(t *testing.T) {
	for _, v := range intervalCases {
		caseRunner[bool](t, v)
	}
}

func TestSort(t *testing.T) {
	for _, v := range sortCases {
		caseRunner[string](t, v)
	}
}
