package main

import (
	"github.com/iot-master-contrib/history"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/web"
)

func main() {
	app := web.CreateEngine()

	//调用启动
	err := history.Startup(app)
	if err != nil {
		log.Fatal(err)
	}

	err = history.Register()
	if err != nil {
		log.Fatal(err)
	}

	//注册静态页面
	fs := app.FileSystem()
	history.Static(fs)

	//监听HTTP
	app.Serve()
}
