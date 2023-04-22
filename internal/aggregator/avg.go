package aggregator

import (
	"context"
	"errors"
)

type avgAggregator struct {
	baseAggregator
	value float64
	count float64
}

func (a *avgAggregator) Push(ctx map[string]interface{}) error {
	res, err := a.expression.EvalFloat64(context.Background(), ctx)
	if err != nil {
		return err
	}
	a.value += res
	a.count++
	return nil
}

func (a *avgAggregator) Pop() (val float64, err error) {
	if a.count == 0 {
		return 0, errors.New("无")
	}
	val = a.value / a.count
	a.value = 0
	a.count = 0
	return
}
