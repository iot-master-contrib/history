package aggregator

import (
	"github.com/PaesslerAG/gval"
	"github.com/zgwit/iot-master/v3/pkg/calc"
)

//baseAggregator 聚合器
type baseAggregator struct {
	expression gval.Evaluable
	dirty      bool
}

//Compile 初始化
func (a *baseAggregator) Compile(expression string) (err error) {
	a.expression, err = calc.New(expression)
	return
}

func (a *baseAggregator) Dirty() bool {
	return a.dirty
}
