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
	case "INC":
		agg = &incAggregator{}
	case "DEC":
		agg = &decAggregator{}
	case "SUM":
		agg = &sumAggregator{}
	case "AVG":
		agg = &avgAggregator{}
	case "COUNT":
		agg = &countAggregator{}
	case "MIN":
		agg = &minAggregator{}
	case "MAX":
		agg = &maxAggregator{}
	case "FIRST":
		agg = &firstAggregator{}
	case "LAST":
		agg = &lastAggregator{}
	default:
		err = fmt.Errorf("Unknown type %s ", m.Type)
		return
	}
	//agg.expression, err = calc.New(a.Expression)
	err = agg.Compile(m.Expression)
	return
}
