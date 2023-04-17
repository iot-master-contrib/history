package internal

import (
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"history/types"
)

type Project struct {
	Points map[string][]types.Aggregator
}

var projects lib.Map[Project]

func LoadPoints(job *types.Job) *Project {
	prj := projects.Load(job.ProductId)
	if prj == nil {
		prj = &Project{Points: map[string][]types.Aggregator{}}
		projects.Store(job.ProductId, prj)
	}

	return prj
}
