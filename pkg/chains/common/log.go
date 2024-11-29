package common

import "fmt"

type LogUnit[T any] struct {
	log string
}

func NewLogUnit[T any](log string) *LogUnit[T] {
	return &LogUnit[T]{log: log}
}

func (u *LogUnit[T]) Do(ctx T) (T, error) {
	fmt.Println(u.log)
	return ctx, nil
}
