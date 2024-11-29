package common

import "github.com/tufitko/tacos/pkg/chains"

type RevertableUnit[T any] struct {
	Unit   chains.Unit[T]
	Revert chains.ErrorCallback[T]
}

func WithRevert[T any](unit chains.Unit[T], revert chains.ErrorCallback[T]) *RevertableUnit[T] {
	return &RevertableUnit[T]{Unit: unit, Revert: revert}
}

func NewRevertableUnit[T any](unit chains.Unit[T], revert chains.ErrorCallback[T]) *RevertableUnit[T] {
	return &RevertableUnit[T]{Unit: unit, Revert: revert}
}

func (u *RevertableUnit[T]) Do(ctx T) (T, error) {
	return u.Unit.Do(ctx)
}

func (u *RevertableUnit[T]) OnErrorCallback() chains.ErrorCallback[T] {
	return u.Revert
}
