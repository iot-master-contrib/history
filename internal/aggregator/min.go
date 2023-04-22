package aggregator

import (
	"context"
	"errors"
)

type minAggregator struct {
	baseAggregator
	value float64
	dirty bool
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

func (a *minAggregator) Pop() (val float64, err error) {
	if !a.dirty {
		return 0, errors.New("无数据")
	}
	val = a.value
	a.dirty = false
	return
}
