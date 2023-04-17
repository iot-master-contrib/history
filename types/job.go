package types

import "github.com/zgwit/iot-master/v3/model"

type Job struct {
	Id        string     `json:"id" xorm:"pk"`
	ProductId string     `json:"product_id" xorm:"index"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Crontab   string     `json:"crontab"`
	Points    []Point    `json:"points"`
	Disabled  bool       `json:"disabled"`
	Created   model.Time `json:"created" xorm:"created"`
}

type Point struct {
	Name string `json:"name"`
	Type string `json:"type"` //store increase average count
}
