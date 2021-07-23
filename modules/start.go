package start

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"strconv"
	"time"
)

var botStartTime = time.Now().Unix()

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	chatId := ctx.EffectiveChat.Id

	// Uptime, In seconds: currentTime - startTime
	aliveSeconds := int(time.Now().Unix()) - int(botStartTime)

	msg := "Already alive since " + strconv.Itoa(aliveSeconds) +
		" seconds"
	bot.SendMessage(chatId, msg, nil)
	return nil
}
