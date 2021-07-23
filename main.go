package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"goBotBoilerplate/modules"
)

// Because why not?
func errPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	var TOKEN string

	// Use TOKEN passed via args else fallback to ENV
	if len(os.Args) == 2 {
		TOKEN = os.Args[1]
	} else {
		TOKEN = os.Getenv("TOKEN")
	}

	bot, err := gotgbot.NewBot(
		TOKEN,
		&gotgbot.BotOpts{
			Client:      http.Client{},
			GetTimeout:  gotgbot.DefaultGetTimeout,
			PostTimeout: gotgbot.DefaultPostTimeout,
		})
	errPanic(err) // Panic if there's an error

	updater := ext.NewUpdater(nil)

	dispatcher := updater.Dispatcher
	dispatcher.AddHandler(handlers.NewCommand("start", start.Start))

	err = updater.StartPolling(bot, &ext.PollingOpts{})
	errPanic(err)

	startMsg :=
		`
----------------
| Bot Started! |
----------------`
	log.Println(startMsg)

	updater.Idle()
}
