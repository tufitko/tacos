package local

import (
	"fmt"

	"github.com/tufitko/tacos/pkg/chains"
	"github.com/tufitko/tacos/pkg/chains/common"
	"github.com/tufitko/tacos/runner"
)

var planChain = chains.Chain[runner.Ctx]{
	Units: []chains.Unit[runner.Ctx]{
		// todo: write normal units
		common.NewLogUnit[runner.Ctx]("generate tree of modules"),
		common.NewLogUnit[runner.Ctx]("mark modules, which was changed (don't mark dependent modules) (changed mark)"),
		common.NewLogUnit[runner.Ctx]("mark all touched modules (just print it in output later) (touched mark)"),
		common.WithRevert[runner.Ctx](
			common.NewNoopUnit[runner.Ctx](),
			func(ctx runner.Ctx, err error) {
				fmt.Println("set error check on Github")
				fmt.Println("comment on Github with error")
			},
		),
		common.NewLogUnit[runner.Ctx]("run plan for changed modules"),
		common.NewLogUnit[runner.Ctx]("set successful check on Github"),
		common.NewLogUnit[runner.Ctx]("comment on Github with success"),
	},
}

var applyChain = chains.Chain[runner.Ctx]{
	Units: []chains.Unit[runner.Ctx]{
		common.NewLogUnit[runner.Ctx]("check requirements"),
		common.NewLogUnit[runner.Ctx]("generate tree of modules"),
		common.NewLogUnit[runner.Ctx]("mark modules, which was changed (don't mark dependent modules) (changed mark)"),
		common.NewLogUnit[runner.Ctx]("mark all touched modules (touched mark)"),
		common.WithRevert[runner.Ctx](
			common.NewNoopUnit[runner.Ctx](),
			func(ctx runner.Ctx, err error) {
				fmt.Println("set error check on Github")
				fmt.Println("comment on Github with error")
			},
		),
		common.NewLogUnit[runner.Ctx]("lock all touched modules"),
		common.NewLogUnit[runner.Ctx]("run apply for changed modules and dependent modules if output was changed (recursive)"),
		common.NewLogUnit[runner.Ctx]("set successful check on Github"),
		common.NewLogUnit[runner.Ctx]("comment on Github with success"),
	},
}

var lockChain = chains.Chain[runner.Ctx]{
	Units: []chains.Unit[runner.Ctx]{
		common.NewLogUnit[runner.Ctx]("check requirements ???"),
		common.NewLogUnit[runner.Ctx]("generate tree of modules"),
		common.NewLogUnit[runner.Ctx]("mark modules, which was changed (don't mark dependent modules) (changed mark)"),
		common.NewLogUnit[runner.Ctx]("mark all touched modules (touched mark)"),
		common.WithRevert[runner.Ctx](
			common.NewNoopUnit[runner.Ctx](),
			func(ctx runner.Ctx, err error) {
				fmt.Println("comment on Github with error")
			},
		),
		common.NewLogUnit[runner.Ctx]("lock all touched modules"),
		common.NewLogUnit[runner.Ctx]("comment on Github with success"),
	},
}

var unlockChain = chains.Chain[runner.Ctx]{
	Units: []chains.Unit[runner.Ctx]{
		common.WithRevert[runner.Ctx](
			common.NewNoopUnit[runner.Ctx](),
			func(ctx runner.Ctx, err error) {
				fmt.Println("comment on Github with error")
			},
		),
		common.NewLogUnit[runner.Ctx]("unlock all locks"),
		common.NewLogUnit[runner.Ctx]("comment on Github with success"),
	},
}
