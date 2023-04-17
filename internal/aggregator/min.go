package aggregator

import (
	"context"
)

type minAggregator struct {
	baseAggregator
	value float64
}

func (a *minAggregator) Push(ctx map[string]interface{}) error {
	res, err := a.expression.EvalFloat64(context.Background(), ctx)
	if err != nil {
		return err
	}

	if !a.dirty {
		a.value = res
	} else if res < a.value {
		a.value = res
	}
	a.dirty = false

	return nil
}

func (a *minAggregator) Value() (val float64) {
	val = a.value
	a.dirty = false
	return
}
