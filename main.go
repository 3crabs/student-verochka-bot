package main

import (
	"errors"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/umputun/go-flags"
	"log"
	"os"
	"student_bot/date"
	"student_bot/messages"
	"student_bot/parser"
)

type Opts struct {
	Token string `short:"t" long:"token" description:"Telegram api token"`
}

var opts Opts

func main() {
	p := flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash|flags.HelpFlag)
	p.SubcommandsOptional = true
	if _, err := p.Parse(); err != nil {
		if err.(*flags.Error).Type != flags.ErrHelp {
			log.Println(errors.New("[ERROR] cli error: " + err.Error()))
		}
		os.Exit(2)
	}

	bot, err := tgbot.NewBotAPI(opts.Token)
	if err != nil {
		log.Println(err)
		return
	}

	u := tgbot.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		// empty message
		if update.Message == nil {
			continue
		}

		// command /start
		if update.Message.Text == "/start" {
			_, _ = bot.Send(tgbot.NewMessage(update.Message.Chat.ID, messages.StartMessage()))
			continue
		}

		// command /help
		if update.Message.Text == "/help" {
			_, _ = bot.Send(tgbot.NewMessage(update.Message.Chat.ID, messages.HelpMessage()))
			continue
		}

		// command /ping
		if update.Message.Text == "/ping" {
			_, _ = bot.Send(tgbot.NewMessage(update.Message.Chat.ID, messages.PongMessage()))
			continue
		}

		// command /today_lessons
		if update.Message.Text == "/today_lessons" {
			_, _ = bot.Send(tgbot.NewMessage(update.Message.Chat.ID, messages.LessonsMessage(
				parser.ParseByDay(date.Today()),
				"Сегодня пар нет",
			)))
			continue
		}

		// command /tomorrow_lessons
		if update.Message.Text == "/tomorrow_lessons" {
			_, _ = bot.Send(tgbot.NewMessage(update.Message.Chat.ID, messages.LessonsMessage(
				parser.ParseByDay(date.Today()+1),
				"Завтра пар нет",
			)))
			continue
		}
	}
}
