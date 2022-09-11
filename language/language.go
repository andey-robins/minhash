package language

import "math/rand"

type lang interface {
	int64 | int | rune
}

type Language[T lang] struct {
	terminals []T
	start     T
	end       T
}

func New[T lang]() *Language[T] {
	return &Language[T]{}
}

func (l *Language[T]) AddIntRange(start, end int) {
	l.start = T(start)
	l.end = T(end)
	for i := start; i < end; i++ {
		l.terminals = append(l.terminals, T(i))
	}
}

func (l *Language[T]) AddInt(i int) {
	l.terminals = append(l.terminals, T(i))
}

func (l *Language[T]) SparseList() []T {
	list := make([]T, 0)
	for i := l.start; i < l.end; i++ {
		if rand.Intn(100) == 0 {
			list = append(list, 1)
		} else {
			list = append(list, 0)
		}
	}
	return list
}
