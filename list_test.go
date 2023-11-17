package collections

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

type listTestCase[T comparable] struct {
	name              string
	input             List[any]
	parameters        []any
	expectPanic       bool
	expected          T
	runnable          func(*testing.T, List[any], []any) T
	nilTypeComparison bool
}

var empty = []any{}
var oneTwoThree = []any{1, 2, 3}

var oneTwoThreeList = NewList[any](1, 2, 3)

func caseRunner[T comparable](t *testing.T, c listTestCase[T]) {
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

var lengthCases = []listTestCase[int]{
	{
		name:        "SimpleList.Length.Empty",
		input:       NewList[any](empty...),
		parameters:  nil,
		expectPanic: false,
		expected:    0,
		runnable: func(t *testing.T, list List[any], parameters []any) int {
			return list.Length()
		},
	},
	{
		name:        "SimpleList.Clone.Length.Empty",
		input:       NewList[any](empty...),
		parameters:  nil,
		expectPanic: false,
		expected:    0,
		runnable: func(t *testing.T, list List[any], parameters []any) int {
			return list.Clone().Length()
		},
	},
	{
		name:        "SimpleList.Length.Value",
		input:       NewListFrom[any](oneTwoThree),
		parameters:  nil,
		expectPanic: false,
		expected:    len(oneTwoThreeList.Elements()),
		runnable: func(t *testing.T, list List[any], parameters []any) int {
			return list.Length()
		},
	},
	{
		name:        "SimpleList.Clone.Length.Value",
		input:       NewListFrom[any](oneTwoThree),
		parameters:  nil,
		expectPanic: false,
		expected:    len(oneTwoThreeList.Elements()),
		runnable: func(t *testing.T, list List[any], parameters []any) int {
			return list.Clone().Length()
		},
	},
}

var emptyCases = []listTestCase[bool]{
	{
		name:        "SimpleList.IsEmpty.Empty",
		input:       NewList[any](empty...),
		parameters:  nil,
		expectPanic: false,
		expected:    true,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.IsEmpty()
		},
	},
	{
		name:        "SimpleList.IsNotEmpty.Empty",
		input:       NewList[any](empty...),
		parameters:  nil,
		expectPanic: false,
		expected:    false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.IsNotEmpty()
		},
	},
	{
		name:        "SimpleList.IsNotEmpty.Filled",
		input:       NewListFrom[any](oneTwoThree),
		parameters:  nil,
		expectPanic: false,
		expected:    true,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.IsNotEmpty()
		},
	},
	{
		name:        "SimpleList.IsEmpty.Filled",
		input:       NewListFrom[any](oneTwoThree),
		parameters:  nil,
		expectPanic: false,
		expected:    false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.IsEmpty()
		},
	},
}

var atCases = []listTestCase[any]{
	{
		name:              "SimpleList.At.Empty",
		input:             NewList[any](empty...),
		parameters:        nil,
		expectPanic:       false,
		expected:          nil,
		nilTypeComparison: false,
		runnable: func(t *testing.T, list List[any], parameters []any) any {
			return list.At(0)
		},
	},
	{
		name:              "SimpleList.At.MutableValue",
		input:             NewListFrom[any](oneTwoThree),
		parameters:        nil,
		expectPanic:       false,
		expected:          true,
		nilTypeComparison: false,
		runnable: func(t *testing.T, list List[any], parameters []any) any {
			n := list.At(0)
			*n = 4
			return list.ElementAt(0).(int) == 4
		},
	},
	{
		name:              "SimpleList.ElementAt.Empty",
		input:             NewList[any](empty...),
		parameters:        nil,
		expectPanic:       true,
		expected:          false,
		nilTypeComparison: false,
		runnable: func(t *testing.T, list List[any], parameters []any) any {
			return list.ElementAt(0).(int) == 1
		},
	},
	{
		name:              "SimpleList.ElementAt.NonMutableValue",
		input:             NewListFrom[any](oneTwoThree),
		parameters:        nil,
		expectPanic:       false,
		expected:          true,
		nilTypeComparison: false,
		runnable: func(t *testing.T, list List[any], parameters []any) any {
			n := list.ElementAt(0)
			if n != 1 {
				t.Error("SimpleList.ElementAt is different from expected")
			}
			n = 4
			a := list.ElementAt(0)
			if a != 1 {
				t.Error("SimpleList.ElementAt is different from expected")
			}
			return a != n
		},
	},
}

var indexCases = []listTestCase[bool]{
	{
		name:        "SimpleList.FirstElement.Empty",
		input:       NewList[any](empty...),
		expected:    false,
		expectPanic: true,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.FirstElement() == 1
		},
	},
	{
		name:        "SimpleList.First.Empty",
		input:       NewList[any](empty...),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.First() == nil
		},
	},
	{
		name:        "SimpleList.LastElement.Empty",
		input:       NewList[any](empty...),
		expected:    false,
		expectPanic: true,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.LastElement() == 1
		},
	},
	{
		name:        "SimpleList.Last.Empty",
		input:       NewList[any](empty...),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.Last() == nil
		},
	},
	{
		name:        "SimpleList.FirstIndexWhere.NotSatisfied",
		input:       NewList[any](empty...),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.FirstIndexWhere(func(a any) bool {
				return true
			}) == -1
		},
	},
	{
		name:        "SimpleList.FirstWhere.NotSatisfied",
		input:       NewList[any](empty...),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.FirstWhere(func(a any) bool {
				return true
			}) == nil
		},
	},
	{
		name:        "SimpleList.FirstWhere.Satisfied",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.FirstWhere(func(a any) bool {
				return true
			}) == list.First()
		},
	},
	{
		name:        "SimpleList.FirstElementWhere.NotSatisfied",
		input:       NewList[any](empty...),
		expected:    true,
		expectPanic: true,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.FirstElementWhere(func(a any) bool {
				return true
			}) == nil
		},
	},
	{
		name:        "SimpleList.FirstElementWhere.Satisfied",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.FirstElementWhere(func(a any) bool {
				return true
			}) == 1
		},
	},
	{
		name:        "SimpleList.LastWhere.NotSatisfied",
		input:       NewList[any](empty...),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.LastWhere(func(a any) bool {
				return true
			}) == nil
		},
	},
	{
		name:        "SimpleList.LastWhere.Satisfied",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.LastWhere(func(a any) bool {
				return true
			}) == list.Last()
		},
	},
	{
		name:        "SimpleList.FirstIndexWhere.Even",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.FirstIndexWhere(func(a any) bool {
				return a.(int)%2 == 0
			}) == 1
		},
	},
	{
		name:        "SimpleList.FirstIndexWhere.Odd",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.FirstIndexWhere(func(a any) bool {
				return a.(int)%2 == 1
			}) == 0
		},
	},
	{
		name:        "SimpleList.LastIndexWhere.Odd",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.LastIndexWhere(func(a any) bool {
				return a.(int)%2 == 1
			}) == 2
		},
	},

	{
		name:        "SimpleList.LastElementWhere.NotSatisfied",
		input:       NewList[any](empty...),
		expected:    true,
		expectPanic: true,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.LastElementWhere(func(a any) bool {
				return true
			}) == nil
		},
	},
	{
		name:        "SimpleList.LastElementWhere.Satisfied",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.LastElementWhere(func(a any) bool {
				return true
			}) == 3
		},
	},
	{
		name:        "SimpleList.IndexWhere.Odd",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			odds := list.IndexWhere(func(a any) bool {
				return a.(int)%2 == 1
			})
			return odds.ElementAt(0) == 0 && odds.ElementAt(1) == 2
		},
	},
}
var whereCases = []listTestCase[bool]{
	{
		name:        "SimpleList.Where.Odd",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			odds := list.Where(func(candidate any) bool {
				return candidate.(int)%2 == 1
			})
			return odds.ElementAt(0).(int) == 1 && odds.ElementAt(1).(int) == 3 && odds.Length() == 2
		},
	},
	{
		name:        "SimpleList.Where.Even",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			odds := list.Where(func(candidate any) bool {
				return candidate.(int)%2 == 0
			})
			return odds.ElementAt(0).(int) == 2 && odds.Length() == 1
		},
	},
}
var mapCases = []listTestCase[bool]{
	{
		name:        "SimpleList.HashMap.String",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			strs := list.Map(func(candidate any) any {
				return fmt.Sprint(candidate)
			})
			return strs.ElementAt(0).(string) == "1" && strs.ElementAt(1).(string) == "2" && strs.ElementAt(2).(string) == "3" && strs.Length() == 3
		},
	},
}

var reduceCases = []listTestCase[bool]{
	{
		name:        "SimpleList.Reduce.Sum",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.Reduce(func(acc any, candidate any, idx int) any {
				return acc.(int) + candidate.(int)
			}, 0) == 6
		},
	},
	{
		name:        "SimpleList.Reduce.HashMap",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			strs := list.Reduce(func(acc any, candidate any, idx int) any {
				return acc.(List[string]).Push(fmt.Sprint(candidate))
			}, NewList[string]()).(List[string])
			return strs.ElementAt(0) == "1" && strs.ElementAt(1) == "2" && strs.ElementAt(2) == "3" && strs.Length() == 3
		},
	},
}

var everyCases = []listTestCase[bool]{
	{
		name:        "SimpleList.Every.lt",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.Every(func(candidate any) bool {
				return candidate.(int) < 4
			})
		},
	},
	{
		name:        "SimpleList.Every.gt",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.Every(func(candidate any) bool {
				return candidate.(int) > 0
			})
		},
	},
	{
		name:        "SimpleList.Every.gt.false",
		input:       NewListFrom[any](oneTwoThree),
		expected:    false,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.Every(func(candidate any) bool {
				return candidate.(int) > 1
			})
		},
	},
}

var someCases = []listTestCase[bool]{
	{
		name:        "SimpleList.Some.lt",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.Some(func(candidate any) bool {
				return candidate.(int) < 4
			})
		},
	},
	{
		name:        "SimpleList.Some.gt",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.Some(func(candidate any) bool {
				return candidate.(int) > 0
			})
		},
	},
	{
		name:        "SimpleList.Every.eq",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.Some(func(candidate any) bool {
				return candidate.(int) == 1
			})
		},
	},

	{
		name:        "SimpleList.Every.eq.false",
		input:       NewListFrom[any](oneTwoThree),
		expected:    false,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.Some(func(candidate any) bool {
				return candidate.(int) == 10
			})
		},
	},
}

var noneCases = []listTestCase[bool]{
	{
		name:        "SimpleList.None.lt",
		input:       NewListFrom[any](oneTwoThree),
		expected:    false,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.None(func(candidate any) bool {
				return candidate.(int) < 4
			})
		},
	},
	{
		name:        "SimpleList.None.gt",
		input:       NewListFrom[any](oneTwoThree),
		expected:    false,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.None(func(candidate any) bool {
				return candidate.(int) > 0
			})
		},
	},
	{
		name:        "SimpleList.None.eq",
		input:       NewListFrom[any](oneTwoThree),
		expected:    false,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.None(func(candidate any) bool {
				return candidate.(int) == 1
			})
		},
	},

	{
		name:        "SimpleList.None.eq.false",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.None(func(candidate any) bool {
				return candidate.(int) == 10
			})
		},
	},
}

var popCases = []listTestCase[bool]{
	{
		name:        "SimpleList.Pop",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			popped := list.Pop()
			return popped.ElementAt(0) == 1 && popped.ElementAt(1) == 2 && popped.Length() == 2

		},
	},
	{
		name:        "SimpleList.Push",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			popped := list.Push(4, 5)
			return popped.ElementAt(0) == 1 && popped.ElementAt(1) == 2 && popped.ElementAt(2) == 3 && popped.ElementAt(3) == 4 && popped.ElementAt(4) == 5 && popped.Length() == 5
		},
	},
	{
		name:        "SimpleList.Pop.Double",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			popped := list.Pop().Pop()
			return popped.ElementAt(0) == 1 && popped.Length() == 1
		},
	},
}

var shiftCases = []listTestCase[bool]{
	{
		name:        "SimpleList.Shift",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			popped := list.Shift()
			return popped.ElementAt(0) == 2 && popped.ElementAt(1) == 3 && popped.Length() == 2

		},
	},
	{
		name:        "SimpleList.Shift.Double",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			popped := list.Shift().Shift()
			return popped.ElementAt(0) == 3 && popped.Length() == 1

		},
	},
}

var setCases = []listTestCase[bool]{
	{
		name:        "SimpleList.Set",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			set := list.Set(1, 4)
			return set.Length() == list.Length() && set.ElementAt(0) == list.ElementAt(0) && set.ElementAt(1) == 4 && set.ElementAt(2) == list.ElementAt(2)

		},
	},
}

var stringCases = []listTestCase[string]{
	{
		name:        "SimpleList.String",
		input:       NewListFrom[any](oneTwoThree),
		expected:    fmt.Sprint(oneTwoThreeList.Elements()),
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) string {
			return list.String()
		},
	},
}

var intervalCases = []listTestCase[bool]{
	{
		name:        "SimpleList.Interval",
		input:       NewListFrom[any](oneTwoThree),
		expected:    true,
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			interval := list.Interval(0, 1)
			return interval.ElementAt(0).(int) == 1 && interval.ElementAt(1).(int) == 2
		},
	},
}

var sortCases = []listTestCase[string]{
	{
		name:        "SimpleList.Sort.Int.Esc",
		input:       NewList[any](5, 1981, 3, 15, 142, 1, 23, 89, 67, 1203, 439, 24),
		expected:    fmt.Sprint([]any{1, 3, 5, 15, 23, 24, 67, 89, 142, 439, 1203, 1981}),
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) string {
			list.Sort(func(a, b any) int {
				return a.(int) - b.(int)
			})
			return list.String()
		},
	},
	{
		name:        "SimpleList.Sort.Int.Desc",
		input:       NewList[any](5, 1981, 3, 15, 142, 1, 23, 89, 67, 1203, 439, 24),
		expected:    fmt.Sprint([]any{1981, 1203, 439, 142, 89, 67, 24, 23, 15, 5, 3, 1}),
		expectPanic: false,
		runnable: func(t *testing.T, list List[any], parameters []any) string {
			list.Sort(func(a, b any) int {
				return b.(int) - a.(int)
			})
			return list.String()
		},
	},
}

var clearCases = []listTestCase[bool]{
	{
		name:     "SimpleList.Clear.IsEmpty",
		input:    NewListFrom[any](oneTwoThree),
		expected: true,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.Clear().IsEmpty()
		},
	},
}

var specificCases = []listTestCase[bool]{
	{
		name:     "SimpleList.ThreadSafety",
		input:    NewListFrom[any](oneTwoThree).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...),
		expected: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			var wg sync.WaitGroup
			dest := list.Clone().Clear()
			for _, element := range list.Elements() {
				wg.Add(1)
				go func(el any) {
					dest.Push(el)
					wg.Done()
				}(element)
			}
			wg.Wait()
			return list.Length() == dest.Length()
		},
	},
	{
		name:     "SimpleList.IsThreadSafe",
		input:    NewListFrom[any](oneTwoThree),
		expected: false,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.IsThreadSafe()
		},
	},
	{
		name:     "SimpleList.IsDynamicallySized",
		input:    NewListFrom[any](oneTwoThree),
		expected: true,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.IsDynamicallySized()
		},
	},
	{
		name:     "SafeList.ThreadSafety",
		input:    NewSafeListFrom[any](oneTwoThree).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...).Push(oneTwoThree...),
		expected: true,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			var wg sync.WaitGroup
			dest := list.Clone().Clear()
			for _, element := range list.Elements() {
				wg.Add(1)
				go func(el any) {
					dest.Push(el)
					wg.Done()
				}(element)
			}
			wg.Wait()
			return list.Length() == dest.Length()
		},
	},
	{
		name:     "SafeList.IsThreadSafe",
		input:    NewSafeListFrom[any](oneTwoThree),
		expected: true,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.IsThreadSafe()
		},
	},
	{
		name:     "SafeList.IsDynamicallySized",
		input:    NewSafeListFrom[any](oneTwoThree),
		expected: true,
		runnable: func(t *testing.T, list List[any], parameters []any) bool {
			return list.IsDynamicallySized()
		},
	},
}

func TestLength(t *testing.T) {
	for _, v := range lengthCases {
		safe := cloneSafe(v)
		caseRunner[int](t, v)
		caseRunner[int](t, safe)
	}
}

func TestEmpty(t *testing.T) {
	for _, v := range emptyCases {
		safe := cloneSafe(v)
		caseRunner[bool](t, v)
		caseRunner[bool](t, safe)
	}
}

func TestAt(t *testing.T) {
	for _, v := range atCases {
		safe := cloneSafe(v)
		caseRunner[any](t, v)
		caseRunner[any](t, safe)
	}
}

func TestIndexes(t *testing.T) {
	for _, v := range indexCases {
		safe := cloneSafe(v)
		caseRunner[bool](t, v)
		caseRunner[bool](t, safe)
	}
}

func TestWhere(t *testing.T) {
	for _, v := range whereCases {
		safe := cloneSafe(v)
		caseRunner[bool](t, v)
		caseRunner[bool](t, safe)
	}
}

func TestMap(t *testing.T) {
	for _, v := range mapCases {
		safe := cloneSafe(v)
		caseRunner[bool](t, v)
		caseRunner[bool](t, safe)
	}
}

func TestReduce(t *testing.T) {
	for _, v := range reduceCases {
		safe := cloneSafe(v)
		caseRunner[bool](t, v)
		caseRunner[bool](t, safe)
	}
}

func TestEvery(t *testing.T) {
	for _, v := range everyCases {
		safe := cloneSafe(v)
		caseRunner[bool](t, v)
		caseRunner[bool](t, safe)
	}
}

func TestSome(t *testing.T) {
	for _, v := range someCases {
		safe := cloneSafe(v)
		caseRunner[bool](t, v)
		caseRunner[bool](t, safe)
	}
}

func TestNone(t *testing.T) {
	for _, v := range noneCases {
		safe := cloneSafe(v)
		caseRunner[bool](t, v)
		caseRunner[bool](t, safe)
	}
}

func TestPop(t *testing.T) {
	for _, v := range popCases {
		safe := cloneSafe(v)
		caseRunner[bool](t, v)
		caseRunner[bool](t, safe)
	}
}

func TestShift(t *testing.T) {
	for _, v := range shiftCases {
		safe := cloneSafe(v)
		caseRunner[bool](t, v)
		caseRunner[bool](t, safe)
	}
}

func TestString(t *testing.T) {
	for _, v := range stringCases {
		safe := cloneSafe(v)
		caseRunner[string](t, v)
		caseRunner[string](t, safe)
	}
}

func TestSet(t *testing.T) {
	for _, v := range setCases {
		safe := cloneSafe(v)
		caseRunner[bool](t, v)
		caseRunner[bool](t, safe)
	}
}

func TestInterval(t *testing.T) {
	for _, v := range intervalCases {
		safe := cloneSafe(v)
		caseRunner[bool](t, v)
		caseRunner[bool](t, safe)
	}
}

func TestSort(t *testing.T) {
	for _, v := range sortCases {
		safe := cloneSafe(v)
		caseRunner[string](t, v)
		caseRunner[string](t, safe)
	}
}

func TestClear(t *testing.T) {
	for _, v := range clearCases {
		safe := cloneSafe(v)
		caseRunner[bool](t, v)
		caseRunner[bool](t, safe)
	}
}

func TestSafety(t *testing.T) {
	for _, v := range specificCases {
		caseRunner[bool](t, v)
	}
}

func cloneSafe[T comparable](t listTestCase[T]) listTestCase[T] {
	return listTestCase[T]{
		name:              strings.Replace(t.name, "SimpleList", "SafeList", -1),
		input:             NewSafeList[any](t.input.Elements()...),
		parameters:        t.parameters,
		expectPanic:       t.expectPanic,
		expected:          t.expected,
		runnable:          t.runnable,
		nilTypeComparison: t.nilTypeComparison,
	}
}
