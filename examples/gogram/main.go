package main

import (
	"log"

	g "github.com/lepeico/gogram"
	w "github.com/lepeico/withstyle"
)

var (
	debug   = w.WithStyles(w.BgWhite, w.Black)
	error   = w.WithStyles(w.BgRed, w.Black, w.Bold)
	warning = w.WithStyles(w.BgOrange, w.White, w.Underline)
	info    = w.WithStyles(w.BgBlue, w.Black, w.Italic)
)

func main() {
	bot, _ := g.New("1358612875:AAGX4URcBIUz-P4TOF4o491E5TYRLtzYey0")

	bot.Use(func(ctx *g.Context) *g.Context {
		log.Printf("%s from %s", debug("Update"), ctx.From().UserName)
		return ctx
	})

	bot.On("Message", func(ctx *g.Context) { ctx.Reply(ctx.Message.Text) })
	bot.Command("ping", func(ctx *g.Context) { ctx.Reply("pong") })
	bot.Start(func(ctx *g.Context) { ctx.Reply("welcome message") })
	bot.Help(func(ctx *g.Context) { ctx.Reply("help message") })

	bot.Launch()
}
