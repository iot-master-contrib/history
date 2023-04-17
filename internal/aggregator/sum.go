package aggregator

import "context"

type sumAggregator struct {
	baseAggregator
	value float64
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

func (a *sumAggregator) Value() (val float64) {
	val = a.value
	a.value = 0
	a.dirty = false
	return
}
