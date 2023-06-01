package types

import "time"

type Job struct {
	Id          string       `json:"id" xorm:"pk"`
	ProductId   string       `json:"product_id" xorm:"index"`
	Name        string       `json:"name"`        //名称
	Desc        string       `json:"desc"`        //说明
	Crontab     string       `json:"crontab"`     //定时计划
	Aggregators []Aggregator `json:"aggregators"` //聚合器
	Disabled    bool         `json:"disabled"`    //禁用
	Created     time.Time    `json:"created" xorm:"created"`
}

type Aggregator struct {
	Expression string `json:"expression"` //表达式
	Type       string `json:"type"`       //聚合算法 inc dec avg count min max sum last first
	Assign     string `json:"assign"`     //赋值
}
