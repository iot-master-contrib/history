package aggregator

import (
	"context"
)

type maxAggregator struct {
	baseAggregator
	value float64
}

func (a *maxAggregator) Push(ctx map[string]interface{}) error {
	res, err := a.expression.EvalFloat64(context.Background(), ctx)
	if err != nil {
		return err
	}

	if !a.dirty {
		a.value = res
	} else if res > a.value {
		a.value = res
	}
	a.dirty = true
	return nil
}

func (a *maxAggregator) Value() (val float64) {
	val = a.value
	a.dirty = false
	return
}