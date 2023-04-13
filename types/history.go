package types

import "github.com/zgwit/iot-master/v3/model"

type History struct {
	Id       int64      `json:"id"`
	DeviceId string     `json:"device_id" xorm:"index"`
	Point    string     `json:"point" xorm:"index"`
	Value    float64    `json:"value"`
	Time     model.Time `json:"time"`
}
