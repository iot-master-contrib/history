package aggregator

import (
	"context"
	"errors"
)

type decAggregator struct {
	baseAggregator
	first float64
	last  float64
	dirty bool
}

func (a *decAggregator) Push(ctx map[string]interface{}) error {
	res, err := a.expression.EvalFloat64(context.Background(), ctx)
	if err != nil {
		return err
	}
	a.last = res

	if !a.dirty {
		a.first = res
		a.dirty = true
	}

	return nil
}

func (a *decAggregator) Pop() (val float64, err error) {
	if !a.dirty {
		return 0, errors.New("无数据")
	}

	val = a.first - a.last
	a.dirty = false
	return
}
