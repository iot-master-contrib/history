package aggregator

import (
	"fmt"
	"history/types"
)

type Aggregator interface {
	Compile(expression string) error
	Push(ctx map[string]interface{}) error
	Pop() (float64, error)
}

// New 新建
func New(m *types.Aggregator) (agg Aggregator, err error) {
	switch m.Type {
	case "inc", "increase":
		agg = &incAggregator{}
	case "dec", "decrease":
		agg = &decAggregator{}
	case "sum":
		agg = &sumAggregator{}
	case "avg", "average":
		agg = &avgAggregator{}
	case "count":
		agg = &countAggregator{}
	case "min", "minimum":
		agg = &minAggregator{}
	case "max", "maximum":
		agg = &maxAggregator{}
	case "first":
		agg = &firstAggregator{}
	case "last":
		agg = &lastAggregator{}
	default:
		err = fmt.Errorf("Unknown type %s ", m.Type)
		return
	}
	//agg.expression, err = calc.New(a.Expression)
	err = agg.Compile(m.Expression)
	return
}
