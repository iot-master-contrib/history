package aggregator

import (
	"context"
	"errors"
)

type sumAggregator struct {
	baseAggregator
	value float64
	dirty bool
}

func (a *sumAggregator) Push(ctx map[string]interface{}) error {
	res, err := a.expression.EvalFloat64(context.Background(), ctx)
	if err != nil {
		return err
	}
	a.value += res
	a.dirty = true
	return nil
}

func (a *sumAggregator) Pop() (val float64, err error) {
	if !a.dirty {
		return 0, errors.New("无数据")
	}
	val = a.value
	a.value = 0
	a.dirty = false
	return
}
