package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"report/types"
)

// @Summary 查询计划任务数量
// @Schemes
// @Description 查询计划任务数量
// @Tags job
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回计划任务数量
// @Router /job/count [post]
func noopJobCount() {}

// @Summary 查询计划任务
// @Schemes
// @Description 查询计划任务
// @Tags job
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Job] 返回计划任务信息
// @Router /job/search [post]
func noopJobSearch() {}

// @Summary 查询计划任务
// @Schemes
// @Description 查询计划任务
// @Tags job
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Job] 返回计划任务信息
// @Router /job/list [get]
func noopJobList() {}

// @Summary 创建计划任务
// @Schemes
// @Description 创建计划任务
// @Tags job
// @Param search body types.Job true "计划任务信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Job] 返回计划任务信息
// @Router /job/create [post]
func noopJobCreate() {}

// @Summary 修改计划任务
// @Schemes
// @Description 修改计划任务
// @Tags job
// @Param id path int true "计划任务ID"
// @Param job body types.Job true "计划任务信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Job] 返回计划任务信息
// @Router /job/{id} [post]
func noopJobUpdate() {}

// @Summary 获取计划任务
// @Schemes
// @Description 获取计划任务
// @Tags job
// @Param id path int true "计划任务ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Job] 返回计划任务信息
// @Router /job/{id} [get]
func noopJobGet() {}

// @Summary 删除计划任务
// @Schemes
// @Description 删除计划任务
// @Tags job
// @Param id path int true "计划任务ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Job] 返回计划任务信息
// @Router /job/{id}/delete [get]
func noopJobDelete() {}

// @Summary 启用计划任务
// @Schemes
// @Description 启用计划任务
// @Tags job
// @Param id path int true "计划任务ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Job] 返回计划任务信息
// @Router /job/{id}/enable [get]
func noopJobEnable() {}

// @Summary 禁用计划任务
// @Schemes
// @Description 禁用计划任务
// @Tags job
// @Param id path int true "计划任务ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Job] 返回计划任务信息
// @Router /job/{id}/disable [get]
func noopJobDisable() {}

// @Summary 导出计划任务
// @Schemes
// @Description 导出计划任务
// @Tags product
// @Accept json
// @Produce octet-stream
// @Router /job/export [get]
func noopJobExport() {}

// @Summary 导入计划任务
// @Schemes
// @Description 导入计划任务
// @Tags product
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回计划任务数量
// @Router /job/import [post]
func noopJobImport() {}

func jobRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Job]())
	app.POST("/search", curd.ApiSearch[types.Job]())
	app.GET("/list", curd.ApiList[types.Job]())
	app.POST("/create", curd.ApiCreate[types.Job](curd.GenerateRandomId[types.Job](8), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[types.Job]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[types.Job](nil, nil,
		"id", "name", "desc", "crontab", "product_id", "points", "disabled"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[types.Job](nil, nil))
	app.GET("/export", curd.ApiExport[types.Job]("job"))
	app.POST("/import", curd.ApiImport[types.Job]())

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisable[types.Job](true, nil, nil))
	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisable[types.Job](false, nil, nil))
}
