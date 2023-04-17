package types

import "github.com/zgwit/iot-master/v3/model"

type Job struct {
	Id          string       `json:"id" xorm:"pk"`
	ProductId   string       `json:"product_id" xorm:"index"`
	Name        string       `json:"name"`
	Desc        string       `json:"desc"`
	Crontab     string       `json:"crontab"`
	Aggregators []Aggregator `json:"aggregators"`
	Disabled    bool         `json:"disabled"`
	Created     model.Time   `json:"created" xorm:"created"`
}

type Aggregator struct {
	Expression string `json:"expression"`
	Type       string `json:"type"` //inc avg count min max sum last first
	Assign     string `json:"assign"`
}
