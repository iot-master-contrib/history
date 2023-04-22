package aggregator

import (
	"context"
	"errors"
)

type incAggregator struct {
	baseAggregator
	first float64
	last  float64

	dirty   bool
	current float64 //TODO 数据漏传，使用均值？
}

func (a *incAggregator) Push(ctx map[string]interface{}) error {
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

func (a *incAggregator) Pop() (val float64, err error) {
	if !a.dirty {
		return 0, errors.New("无数据")
	}

	val = a.last - a.first
	a.dirty = false
	return
}
