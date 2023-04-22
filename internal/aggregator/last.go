package aggregator

import (
	"context"
	"errors"
)

type lastAggregator struct {
	baseAggregator
	value float64
	dirty bool
}

func (a *lastAggregator) Push(ctx map[string]interface{}) error {
	res, err := a.expression.EvalFloat64(context.Background(), ctx)
	if err != nil {
		return err
	}

	a.value = res
	a.dirty = true
	return nil
}

func (a *lastAggregator) Pop() (val float64, err error) {
	if !a.dirty {
		return 0, errors.New("无数据")
	}
	val = a.value
	a.dirty = false
	return
}
