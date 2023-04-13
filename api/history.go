package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iot-master-contribe/history/types"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询历史数量
// @Schemes
// @Description 查询历史
// @Tags history
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回历史数量
// @Router /history/count [post]
func noopHistoryCount() {}

// @Summary 查询历史
// @Schemes
// @Description 查询历史
// @Tags history
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.History] 返回历史信息
// @Router /history/search [post]
func noopHistorySearch() {}

// @Summary 查询历史
// @Schemes
// @Description 查询历史
// @Tags history
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.History] 返回历史信息
// @Router /history/list [get]
func noopHistoryList() {}

// @Summary 删除历史
// @Schemes
// @Description 删除历史
// @Tags history
// @Param id path int true "历史ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.History] 返回历史信息
// @Router /history/{id}/delete [get]
func noopHistoryDelete() {}

func historyRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.History]())

	app.POST("/search", curd.ApiSearch[types.History]())

	app.GET("/list", curd.ApiList[types.History]())

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDelete[types.History](nil, nil))

}
