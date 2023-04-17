package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.RouterGroup) {

	jobRouter(app.Group("/job"))

	historyRouter(app.Group("/history"))

	aggregateRouter(app.Group("/aggregate"))

}
