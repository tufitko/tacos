package local

import (
	"github.com/tufitko/tacos/pkg/chains"
	"github.com/tufitko/tacos/runner"
)

type Runner struct {
	chains map[runner.ActionType]chains.Chain[runner.Ctx]
}

func NewRunner() *Runner {
	return &Runner{
		chains: map[runner.ActionType]chains.Chain[runner.Ctx]{
			runner.ATPlan:   planChain,
			runner.AtApply:  applyChain,
			runner.ATLock:   lockChain,
			runner.ATUnlock: unlockChain,
		},
	}
}

func (r *Runner) Do(ctx runner.Ctx, action runner.Action) error {
	chain, exist := r.chains[action.Type]
	if !exist {
		return runner.ErrActionNotSupported
	}

	return chain.Do(ctx)
}
