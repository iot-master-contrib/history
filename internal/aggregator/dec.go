package aggregator

import "context"

type decAggregator struct {
	baseAggregator
	first float64
	last  float64
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

func (a *decAggregator) Value() (val float64) {
	val = a.first - a.last
	a.dirty = false
	return
}
