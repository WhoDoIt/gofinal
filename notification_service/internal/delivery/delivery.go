package delivery

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
)

type Deliverer struct {
	Bot *bot.Bot
}

func (d *Deliverer) SendMessage(ctx context.Context, chat_id int64, message string) {
	_, err := d.Bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chat_id,
		Text:   message,
	})
	if err != nil {
		fmt.Println(chat_id)
		fmt.Println(err)
	}
}
