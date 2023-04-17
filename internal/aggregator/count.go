package aggregator

import "context"

type countAggregator struct {
	baseAggregator
	count float64
}

func (a *countAggregator) Push(ctx map[string]interface{}) error {
	res, err := a.expression.EvalBool(context.Background(), ctx)
	if err != nil {
		return err
	}
	if res {
		a.count++
	}
	a.dirty = true
	return nil
}

func (a *countAggregator) Value() (val float64) {
	val = a.count
	a.count = 0
	a.dirty = false
	return
}
