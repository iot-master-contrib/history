package aggregator

import "context"

type lastAggregator struct {
	baseAggregator
	value float64
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

func (a *lastAggregator) Value() (val float64) {
	val = a.value
	a.dirty = false
	return
}
