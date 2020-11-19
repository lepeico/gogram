package main

import (
	"fmt"

	"github.com/lepeico/gogram/pkg/telegram"
)

func main() {
	tg := telegram.New("1358612875:AAEe5git918PHoNymHsj7h7kosGJxY9PyiU")

	res, _ := tg.GetMe()
	fmt.Println(res)
}
