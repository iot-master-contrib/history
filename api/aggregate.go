package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
)

func aggregateRouter(app *gin.RouterGroup) {

	app.GET("", aggregateHistory)

	app.GET("/year", createAggregateByDate("%Y"))

	app.GET("/month", createAggregateByDate("%Y-%m"))

	app.GET("/day", createAggregateByDate("%Y-%m-%d"))

	app.GET("/hour", createAggregateByDate("%Y-%m-%d %H"))

	app.GET("/minute", createAggregateByDate("%Y-%m-%d %H:%i"))

	app.GET("/type", aggregateByType)

	app.GET("/area", aggregateByArea)

	app.GET("/group", aggregateByGroup)

}

type Result struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Total float64 `json:"total"`
}

// @Summary 按区域统计
// @Schemes
// @Description 按区域统计
// @Tags aggregate
// @Param point query string true "数据点"
// @Param start query string false "起始时间"
// @Param end query string false "结束时间"
// @Produce json
// @Success 200 {object} ReplyData[Result] 返回统计数据
// @Router /aggregate/area [get]
func aggregateByArea(ctx *gin.Context) {
	point := ctx.Query("point")
	start := ctx.Query("start")
	end := ctx.Query("end")

	var results []Result
	query := db.Engine.Table([]string{"history", "h"}).
		Select("dg.id, dg.name, sum(h.value) as total").
		Join("INNER", []string{"device", "d"}, "d.id = h.device_id").
		Join("INNER", []string{"device_area", "dg"}, "dg.id = d.group_id").
		And("h.point = ?", point)
	if start != "" {
		query.And("h.time >= ?", start)
	}
	if end != "" {
		query.And("h.time <= ?", end)
	}
	err := query.GroupBy("dg.id").Find(&results)
	//And("h.time > ?", time.Now().Format("2006-01-02 15:04:05")).
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, results)
}

// @Summary 按分组统计
// @Schemes
// @Description 按分组统计
// @Tags aggregate
// @Param point query string true "数据点"
// @Param start query string false "起始时间"
// @Param end query string false "结束时间"
// @Produce json
// @Success 200 {object} ReplyData[Result] 返回统计数据
// @Router /aggregate/group [get]
func aggregateByGroup(ctx *gin.Context) {
	point := ctx.Query("point")
	start := ctx.Query("start")
	end := ctx.Query("end")

	var results []Result
	query := db.Engine.Table([]string{"history", "h"}).
		Select("dg.id, dg.name, sum(h.value) as total").
		Join("INNER", []string{"device", "d"}, "d.id = h.device_id").
		Join("INNER", []string{"device_group", "dg"}, "dg.id = d.group_id").
		And("h.point = ?", point)
	if start != "" {
		query.And("h.time >= ?", start)
	}
	if end != "" {
		query.And("h.time <= ?", end)
	}
	err := query.GroupBy("dg.id").Find(&results)
	//And("h.time > ?", time.Now().Format("2006-01-02 15:04:05")).
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, results)
}

// @Summary 按类型统计
// @Schemes
// @Description 按类型统计
// @Tags aggregate
// @Param point query string true "数据点"
// @Param start query string false "起始时间"
// @Param end query string false "结束时间"
// @Produce json
// @Success 200 {object} ReplyData[Result] 返回统计数据
// @Router /aggregate/type [get]
func aggregateByType(ctx *gin.Context) {
	point := ctx.Query("point")
	start := ctx.Query("start")
	end := ctx.Query("end")

	var results []Result
	query := db.Engine.Table([]string{"history", "h"}).
		Select("dg.id, dg.name, sum(h.value) as total").
		Join("INNER", []string{"device", "d"}, "d.id = h.device_id").
		Join("INNER", []string{"device_type", "dg"}, "dg.id = d.type_id").
		And("h.point = ?", point)
	if start != "" {
		query.And("h.time >= ?", start)
	}
	if end != "" {
		query.And("h.time <= ?", end)
	}
	err := query.GroupBy("dg.id").Find(&results)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, results)
}

type ResultTime struct {
	Time  string  `json:"time"`
	Total float64 `json:"total"`
}

type ParamAggregate struct {
	Point string `form:"point" json:"point"` //数据点位
	Start string `form:"start" json:"start"` //起始时间
	End   string `form:"end" json:"end"`     //结束时间
	Type  string `form:"type" json:"type"`   //设备类型
	Area  string `form:"area" json:"area"`   //设备区域
	Group string `form:"group" json:"group"` //设备分组
}

// @Summary 原始数据按时间统计
// @Schemes
// @Description 原始数据按时间统计
// @Tags aggregate
// @Param search query ParamAggregate true "数据点"
// @Produce json
// @Success 200 {object} ReplyData[Result] 返回统计数据
// @Router /aggregate [get]
func aggregateHistory(ctx *gin.Context) {
	var param ParamAggregate
	err := ctx.ShouldBindQuery(&param)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	var results []ResultTime
	query := db.Engine.Table([]string{"history", "h"}).
		Select("h.time, sum(h.value) as total")
	if param.Type != "" || param.Area != "" || param.Group != "" {
		query.Join("INNER", []string{"device", "d"}, "d.id = h.device_id")

		if param.Type != "" {
			query.And("d.type_id = ?", param.Type)
		}
		if param.Area != "" {
			query.And("d.area_id = ?", param.Area)
		}
		if param.Group != "" {
			query.And("d.group_id = ?", param.Group)
		}
	}
	query.And("h.point = ?", param.Point)
	if param.Start != "" {
		query.And("h.time >= ?", param.Start)
	}
	if param.End != "" {
		query.And("h.time <= ?", param.End)
	}
	err = query.GroupBy("time").Find(&results)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, results)
}

func createAggregateByDate(format string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var param ParamAggregate
		err := ctx.ShouldBindQuery(&param)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		var results []ResultTime
		query := db.Engine.Table([]string{"history", "h"}).
			Select("date_format(h.time, '" + format + "') as date, sum(h.value) as total")
		if param.Type != "" || param.Area != "" || param.Group != "" {
			query.Join("INNER", []string{"device", "d"}, "d.id = h.device_id")

			if param.Type != "" {
				query.And("d.type_id = ?", param.Type)
			}
			if param.Area != "" {
				query.And("d.area_id = ?", param.Area)
			}
			if param.Group != "" {
				query.And("d.group_id = ?", param.Group)
			}
		}

		query.And("h.point = ?", param.Point)
		if param.Start != "" {
			query.And("h.time >= ?", param.Start)
		}
		if param.End != "" {
			query.And("h.time <= ?", param.End)
		}
		err = query.GroupBy("date").Find(&results)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		curd.OK(ctx, results)
	}
}

// @Summary 按年统计
// @Schemes
// @Description 按年统计
// @Tags aggregate
// @Param search query ParamAggregate true "参数"
// @Produce json
// @Success 200 {object} ReplyData[Result] 返回统计数据
// @Router /aggregate/year [get]
func noopAggregateByYear() {
}

// @Summary 按月统计
// @Schemes
// @Description 按月统计
// @Tags aggregate
// @Param search query ParamAggregate true "参数"
// @Produce json
// @Success 200 {object} ReplyData[Result] 返回统计数据
// @Router /aggregate/month [get]
func noopAggregateByMonth() {
}

// @Summary 按日统计
// @Schemes
// @Description 按日统计
// @Tags aggregate
// @Param search query ParamAggregate true "参数"
// @Produce json
// @Success 200 {object} ReplyData[Result] 返回统计数据
// @Router /aggregate/day [get]
func noopAggregateByDay() {
}

// @Summary 按小时统计
// @Schemes
// @Description 按小时统计
// @Tags aggregate
// @Param search query ParamAggregate true "参数"
// @Produce json
// @Success 200 {object} ReplyData[Result] 返回统计数据
// @Router /aggregate/hour [get]
func noopAggregateByHour() {
}

// @Summary 按分钟统计
// @Schemes
// @Description 按分钟统计
// @Tags aggregate
// @Param search query ParamAggregate true "参数"
// @Produce json
// @Success 200 {object} ReplyData[Result] 返回统计数据
// @Router /aggregate/minute [get]
func noopAggregateByMinute() {
}
