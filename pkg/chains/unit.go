package chains

type Unit[T any] interface {
	Do(ctx T) (T, error)
}

type ErrorCallback[T any] func(ctx T, err error)

type RevertableUnit[T any] interface {
	OnErrorCallback() ErrorCallback[T]
}
