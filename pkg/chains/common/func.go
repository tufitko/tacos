package common

type FuncUnit[T any] struct {
	Func func(ctx T) (T, error)
}

func NewFuncUnit[T any](f func(ctx T) (T, error)) *FuncUnit[T] {
	return &FuncUnit[T]{Func: f}
}

func (u *FuncUnit[T]) Do(ctx T) (T, error) {
	return u.Func(ctx)
}
