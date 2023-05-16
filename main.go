package main

import (
	"embed"
	"encoding/json"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/banner"
	"github.com/zgwit/iot-master/v3/pkg/build"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"history/api"
	"history/config"
	_ "history/docs"
	"history/internal"
	"history/types"
	"net/http"
)

//go:embed all:app/history
var wwwFiles embed.FS

// @title 历史数据接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /app/history/api/
// @query.collection.format multi
func main() {
	banner.Print("iot-master-plugin:history")
	build.Print()

	config.Load()

	err := log.Open()
	if err != nil {
		log.Fatal(err)
	}

	//加载数据库
	err = db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//同步表结构
	err = db.Engine.Sync2(
		new(types.History), new(types.Job),
	)
	if err != nil {
		log.Fatal(err)
	}

	//MQTT总线
	err = mqtt.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer mqtt.Close()

	//注册应用
	//for _, v := range config.Config.Apps {
	payload, _ := json.Marshal(model.App{
		Id:   "history",
		Name: "历史数据",
		Entries: []model.AppEntry{{
			Path: "app/history/history",
			Name: "历史数据",
		}, {
			Path: "app/history/job",
			Name: "计划任务",
		}},
		Type:    "tcp",
		Address: "http://localhost" + web.GetOptions().Addr,
		Icon: "/app/history/assets/history.svg",
	})
	_ = mqtt.Publish("master/register", payload, false, 0)
	//}

	internal.SubscribeProperty(mqtt.Client)

	err = internal.StartJobs()
	if err != nil {
		log.Fatal(err)
	}
	defer internal.StopJobs()

	app := web.CreateEngine()

	//注册前端接口
	api.RegisterRoutes(app.Group("/app/history/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(app.Group("/app/history"))

	//前端静态文件
	app.RegisterFS(http.FS(wwwFiles), "", "app/history/index.html")

	//监听HTTP
	app.Serve()
}
