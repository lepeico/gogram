package main

import "github.com/lepeico/gogram"

func main() {
	bot, _ := gogram.New("1234567890:LOVESEX-MAKINGBOTSWITHGOGRAM")

	bot.On("text", func(ctx *gogram.Context) {
		ctx.Reply(ctx.Message.Text)
	})

	bot.Launch()
}
