package telegram

import (
	"github.com/gin-gonic/gin"
	"github.com/iEricKoh/snapbot/server/handlers/telegram"
)

func Routes(router *gin.RouterGroup) {
	bot := router.Group("/telegram")
	{
		handler := &telegram.TelegramHandler{}
		bot.POST("/", handler.SendTelegramMessage)
	}
}
