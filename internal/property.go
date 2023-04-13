package internal

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iot-master-contribe/history/config"
	"github.com/zgwit/iot-master/v3/pkg/convert"
	"strings"
)

var last = map[string]map[string]float64{}
var current = map[string]map[string]float64{}

func SubscribeProperty(client mqtt.Client) {
	//订阅消息
	client.Subscribe("up/property/+/+", 0, func(client mqtt.Client, message mqtt.Message) {
		topics := strings.Split(message.Topic(), "/")
		pid := topics[2]
		id := topics[3]

		//不关心的数据，不处理
		store, ok := config.Config.Store[pid]
		if !ok {
			return
		}

		//解析数据
		var properties map[string]interface{}
		err := json.Unmarshal(message.Payload(), &properties)
		if err != nil {
			//log
			return
		}

		dev := EnsureDevice(id)
		dev.Store = store

		//缓存当前值
		for k, _ := range store {
			if v, ok := properties[k]; ok {
				dev.Values[k] = convert.ToFloat64(v)
			}
		}

	})
}
