package aggregator

import "context"

type firstAggregator struct {
	baseAggregator
	value float64
}

func (a *firstAggregator) Push(ctx map[string]interface{}) error {
	if !a.dirty {
		return nil
	}
	res, err := a.expression.EvalFloat64(context.Background(), ctx)
	if err != nil {
		return err
	}

	a.value = res
	a.dirty = true
	return nil
}

func (a *firstAggregator) Value() (val float64) {
	val = a.value
	a.dirty = false
	return
}
