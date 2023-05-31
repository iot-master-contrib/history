package history

import (
	"embed"
	"encoding/json"
	"github.com/iot-master-contrib/history/api"
	"github.com/iot-master-contrib/history/config"
	_ "github.com/iot-master-contrib/history/docs"
	"github.com/iot-master-contrib/history/internal"
	"github.com/iot-master-contrib/history/types"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/banner"
	"github.com/zgwit/iot-master/v3/pkg/build"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
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
}

func Startup(app *web.Engine) error {
	banner.Print("iot-master-plugin:history")
	build.Print()

	config.Load()

	err := log.Open()
	if err != nil {
		return err
	}

	//加载数据库
	err = db.Open()
	if err != nil {
		return err
	}
	//defer db.Close()

	//同步表结构
	err = db.Engine.Sync2(
		new(types.History), new(types.Job),
	)
	if err != nil {
		return err
	}

	//MQTT总线
	err = mqtt.Open()
	if err != nil {
		return err
	}
	//defer mqtt.Close()


	internal.SubscribeProperty(mqtt.Client)

	err = internal.StartJobs()
	if err != nil {
		return err
	}
	//defer internal.StopJobs()

	//注册前端接口
	api.RegisterRoutes(app.Group("/app/history/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(app.Group("/app/history"), "history")

	return nil
}

func Register() error {
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
		Icon:    "/app/history/assets/history.svg",
	})
	return mqtt.Publish("master/register", payload, false, 0)
}

func Static(fs *web.FileSystem) {
	//前端静态文件
	fs.Put("/app/history", http.FS(wwwFiles), "", "app/history/index.html")
}

func Shutdown() error {

	//只关闭Web就行了，其他通过defer关闭

	return nil
}
