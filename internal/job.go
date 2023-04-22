package internal

import (
	"github.com/robfig/cron/v3"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"history/internal/aggregator"
	"history/types"
	"time"
)

var _cron *cron.Cron

type Device struct {
	aggregators []aggregator.Aggregator
}

func (d *Device) Push(ctx map[string]interface{}) {
	for _, a := range d.aggregators {
		err := a.Push(ctx)
		if err != nil {
			log.Error(err)
		}
	}
}

type Job struct {
	model   types.Job
	devices lib.Map[Device]
}

func (j *Job) Push(id string, ctx map[string]interface{}) {
	dev := j.devices.Load(id)
	if dev == nil {
		dev = &Device{}
		j.devices.Store(id, dev)

		for _, a := range j.model.Aggregators {
			aa, ee := aggregator.New(&a)
			if ee != nil {
				log.Error(ee)
				continue
			}
			dev.aggregators = append(dev.aggregators, aa)
		}
	}
	dev.Push(ctx)
}

func Push(pid, id string, ctx map[string]interface{}) {
	jobs.Range(func(name string, job *Job) bool {
		if job.model.ProductId != pid {
			return true
		}

		job.Push(id, ctx)
		return true
	})
}

var jobs lib.Map[Job]

func NewJob(job *types.Job) error {
	//job.ProductId
	j := &Job{
		model: *job,
	}
	jobs.Store(job.Id, j)

	_, err := _cron.AddFunc(job.Crontab, func() {
		now := time.Now()
		tm := model.Time(now)
		//ts := now.Unix()

		log.Info("[Job] active ", now, job.Id, job.Name)

		var stores []types.History
		j.devices.Range(func(name string, dev *Device) bool {
			for k, m := range dev.aggregators {
				stores = append(stores, types.History{
					DeviceId: name,
					Point:    job.Aggregators[k].Assign,
					Value:    m.Value(),
					Time:     tm,
				})
			}
			return true
		})

		//入库
		if len(stores) > 0 {
			n, err := db.Engine.Insert(stores)
			if err != nil {
				log.Error(err)
			} else {
				log.Info("[Job] store  ", n)
			}
		}

	})

	return err
}

func StartJobs() error {
	_cron = cron.New()

	var js []*types.Job
	err := db.Engine.Find(&js)
	if err != nil {
		return err
	}

	for _, j := range js {
		err = NewJob(j)
		if err != nil {
			log.Error(err)
		}
	}

	//_, err := _cron.AddFunc("* * * * *", storeJob) //测试，每分钟执行一次

	_cron.Start()
	return nil
}

func StopJobs() {
	_cron.Stop()
}
