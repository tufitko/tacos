package chains

import "context"

type Chain[T context.Context] struct {
	Units []Unit[T]
}

func (c *Chain[T]) Do(ctx T) error {
	onErrorCallbacks := make([]ErrorCallback[T], 0)
	for _, unit := range c.Units {
		var err error
		ctx, err = unit.Do(ctx)
		if err != nil {
			return c.onError(ctx, onErrorCallbacks, err)
		}
		if err = ctx.Err(); err != nil {
			return c.onError(ctx, onErrorCallbacks, err)
		}
		if revertable, ok := unit.(RevertableUnit[T]); ok {
			onErrorCallbacks = append(onErrorCallbacks, revertable.OnErrorCallback())
		}
	}
	return nil
}

func (c *Chain[T]) onError(ctx T, callbacks []ErrorCallback[T], err error) error {
	for i := len(callbacks) - 1; i >= 0; i-- {
		callbacks[i](ctx, err)
	}
	return err
}
