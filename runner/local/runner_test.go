package local_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/tufitko/tacos/runner"
	"github.com/tufitko/tacos/runner/local"
)

func TestRunner(t *testing.T) {
	rnr := local.NewRunner()
	rnr.Do(runner.Ctx{context.Background()}, runner.Action{Type: runner.ATPlan})

	fmt.Print("\n\n\n\n\n")
	rnr.Do(runner.Ctx{context.Background()}, runner.Action{Type: runner.AtApply})
}
