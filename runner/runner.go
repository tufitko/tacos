package runner

type Runner interface {
	Do(Ctx, Action) error
}
