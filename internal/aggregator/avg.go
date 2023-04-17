package aggregator

import "context"

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
	a.dirty = true
	return nil
}

func (a *avgAggregator) Value() (val float64) {
	if a.count == 0 {
		return 0
	}
	val = a.value / a.count
	a.value = 0
	a.count = 0
	a.dirty = false
	return
}
