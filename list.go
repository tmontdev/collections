package iterable

func NewList[T any](elements ...T) *List[T] {
	l := &List[T]{}
	return l.Push(elements...).(*List[T])
}

type List[T any] []T

func (l *List[T]) Length() int {
	return len(l.Elements())
}

func (l *List[T]) IsEmpty() bool {
	return l.Length() == 0
}

func (l *List[T]) IsNotEmpty() bool {
	return !l.IsEmpty()
}

func (l *List[T]) At(i int) (at *T) {
	if l.Length()-1 <= i {
		return
	}
	at = &l.Elements()[i]
	return
}

func (l *List[T]) ElementAt(i int) T {
	return l.Elements()[i]
}

func (l *List[T]) Elements() []T {
	return *l
}

func (l *List[T]) Push(elements ...T) Iterable[T] {
	*l = append(l.Elements(), elements...)
	return l
}

func (l *List[T]) Clone() Iterable[T] {
	return NewList[T](l.Elements()...)
}

func (l *List[T]) FirstElement() T {
	return l.Elements()[0]
}

func (l *List[T]) First() *T {
	return &l.Elements()[0]
}

func (l *List[T]) LastElement() T {
	return l.Elements()[l.Length()-1]
}

func (l *List[T]) Last() *T {
	return &l.Elements()[l.Length()-1]
}

func (l *List[T]) FirstIndexWhere(handler PredicateHandler[T]) int {
	for i, v := range l.Elements() {
		if handler(v) {
			return i
		}
	}
	return -1
}

func (l *List[T]) LastIndexWhere(handler PredicateHandler[T]) int {
	idx := -1
	for i, v := range l.Elements() {
		if handler(v) {
			idx = i
		}
	}
	return idx
}

func (l *List[T]) IndexWhere(handler PredicateHandler[T]) Iterable[int] {
	list := NewList[int]()
	for i, v := range l.Elements() {
		if handler(v) {
			list.Push(i)
		}
	}
	return list
}
