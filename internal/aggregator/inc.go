package aggregator

import "context"

type incAggregator struct {
	baseAggregator
	first float64
	last  float64
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

func (a *incAggregator) Value() (val float64) {
	val = a.last - a.first
	a.dirty = false
	return
}
