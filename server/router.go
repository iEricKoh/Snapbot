package server

import (
	"github.com/gin-gonic/gin"
	"github.com/iEricKoh/snapbot/server/routes/telegram"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	telegram.Routes(api)

	return router
}
