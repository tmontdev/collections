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
