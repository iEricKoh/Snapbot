package telegram

import (
	"net/http"

	. "github.com/iEricKoh/snapbot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *TelegramHandler) SendTelegramMessage(c *gin.Context) {
	var params BotParams

	bot, err := tgbotapi.NewBotAPI(Config.TELEGRAM_TOKEN)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindWith(&params, binding.FormPost); err == nil {
		message := tgbotapi.NewMessage(362564169, params.Content)

		message.ParseMode = "HTML"
		message.DisableWebPagePreview = false

		result, _ := bot.Send(message)

		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

type TelegramHandler struct{}

type BotParams struct {
	Content string `form:"content" json:"content" binding:"required"`
}
