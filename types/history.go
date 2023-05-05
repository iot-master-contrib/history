package types

import "github.com/zgwit/iot-master/v3/model"

type History struct {
	Id       int64      `json:"id"`
	DeviceId string     `json:"device_id" xorm:"index"`
	Device   string     `json:"device,omitempty" xorm:"-"`
	Point    string     `json:"point" xorm:"index"` //数据点
	Value    float64    `json:"value"`              //值
	Time     model.Time `json:"time"`
}
