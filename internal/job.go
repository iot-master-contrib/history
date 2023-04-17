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

type Job struct {
	projectId   string
	aggregators []aggregator.Aggregator
	devices     lib.Map[Device]
}

var jobs lib.Map[Job]

func NewJob(job *types.Job) error {
	//job.ProductId
	j := &Job{}

	for _, m := range job.Aggregators {
		//此处会重复编译
		a, err := aggregator.New(&m)
		if err != nil {
			return err
		}
		j.aggregators = append(j.aggregators, a)
	}

	_, err := _cron.AddFunc(job.Crontab, func() {
		now := time.Now()
		tm := model.Time(now)
		//ts := now.Unix()

		log.Info("[Job] active ", now, job.Id, job.Name)

		var stores []types.History
		j.devices.Range(func(name string, dev *Device) bool {
			for k, m := range j.aggregators {
				if !m.Dirty() {
					continue
				}

				value := m.Value()
				stores = append(stores, types.History{
					DeviceId: "",
					Point:    job.Aggregators[k].Assign,
					Value:    value,
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

	return nil
}

func StartJobs() error {
	_cron = cron.New()

	//_, err := _cron.AddFunc("* * * * *", storeJob) //测试，每分钟执行一次
	//_, err := _cron.AddFunc("@hourly", storeJob)
	//_, err := _cron.AddFunc(config.Config.Crontab, storeJob)
	//if err != nil {
	//	return err
	//}

	_cron.Start()
	return nil
}

func StopJobs() {
	_cron.Stop()
}

var last int64
