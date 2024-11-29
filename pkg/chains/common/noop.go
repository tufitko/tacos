package common

type NoopUnit[T any] struct{}

func NewNoopUnit[T any]() *NoopUnit[T] {
	return &NoopUnit[T]{}
}

func (u *NoopUnit[T]) Do(ctx T) (T, error) {
	return ctx, nil
}
