package delivery

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
)

type Deliverer struct {
	Bot     *bot.Bot
	InfoLog *log.Logger
}

func (d *Deliverer) SendMessage(ctx context.Context, chat_id int64, message string) {
	_, err := d.Bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chat_id,
		Text:   message,
	})
	if err != nil {
		d.InfoLog.Println(chat_id)
		d.InfoLog.Println(err)
	}
}
