package telegram

import (
	"snapbot/server/handlers/telegram"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	bot := router.Group("/telegram")
	{
		handler := &telegram.TelegramHandler{}
		bot.POST("/", handler.SendTelegramMessage)
	}
}
