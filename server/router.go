package server

import (
	"snapbot/server/routes/telegram"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	telegram.Routes(api)

	return router
}
