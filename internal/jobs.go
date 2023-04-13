package internal

import (
	"github.com/iot-master-contribe/history/config"
	"github.com/iot-master-contribe/history/types"
	"github.com/robfig/cron/v3"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/convert"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"time"
)

var _cron *cron.Cron

func StartJobs() error {
	_cron = cron.New()
	//_, err := _cron.AddFunc("* * * * *", storeJob) //测试，每分钟执行一次
	//_, err := _cron.AddFunc("@hourly", storeJob)
	_, err := _cron.AddFunc(config.Config.Crontab, storeJob)
	if err != nil {
		return err
	}

	_cron.Start()
	return nil
}

func StopJobs() {
	_cron.Stop()
}

var last int64

func storeJob() {
	now := time.Now()
	tm := model.Time(now)
	//ts := now.Unix()

	log.Info("[Job] store start ", now)

	var stores []types.History
	devices.Range(func(id string, dev *Device) bool {
		//如果没有新数据，就不保存
		if dev.Update < last {
			//log.Info("skip", id, "offline")
			//清空历史，避免计算错误
			dev.Values = map[string]float64{}
			dev.Last = map[string]float64{}
			return true
		}

		//检查需要保存的变量
		for k, s := range dev.Store {
			if v, ok := dev.Values[k]; ok {
				value := convert.ToFloat64(v)
				if s == "save" {
					stores = append(stores, types.History{
						DeviceId: id,
						Point:    k,
						Value:    value,
						Time:     tm,
					})
				} else if s == "increase" {
					if l, ok := dev.Last[k]; ok {
						stores = append(stores, types.History{
							DeviceId: id,
							Point:    k,
							Value:    value - l,
							Time:     tm,
						})
					}
					dev.Last[k] = value
				}
			}

		}
		return true
	})

	//更新上一次时间
	last = now.Unix()

	//入库
	if len(stores) > 0 {
		n, err := db.Engine.Insert(stores)
		if err != nil {
			log.Error(err)
		} else {
			log.Info("[Job] store  ", n)
		}
	}
}
