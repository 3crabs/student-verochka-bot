package main

import (
	"errors"
	weather "github.com/3crabs/go-yandex-weather-api"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/umputun/go-flags"
	"log"
	"os"
	"strings"
	"student_bot/commands"
	"student_bot/date"
	"student_bot/messages"
	"student_bot/parser"
	"time"
)

type Opts struct {
	Token string `short:"t" long:"token" description:"Telegram api token"`
	Key   string `short:"k" long:"key" description:"Yandex weather API key"`
}

var opts Opts

func main() {
	run()
}

func run() {
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

		text := update.Message.Text
		chatId := update.Message.Chat.ID

		switch commands.Command(strings.Replace(text, "@student_verochka_bot", "", 1)) {

		case commands.Start:
			_, _ = bot.Send(tgbot.NewMessage(chatId, messages.StartMessage()))

		case commands.Help:
			_, _ = bot.Send(tgbot.NewMessage(chatId, messages.HelpMessage()))

		case commands.Ping:
			_, _ = bot.Send(tgbot.NewMessage(chatId, messages.PongMessage()))

		case commands.TodayLessons:
			_, _ = bot.Send(tgbot.NewMessage(chatId, messages.LessonsMessage(
				parser.ParseByDay(date.Today()),
				"Сегодня, "+update.Message.From.FirstName+", эти пары:",
				"Сегодня пар нет",
			)))

		case commands.TomorrowLessons:
			_, _ = bot.Send(tgbot.NewMessage(chatId, messages.LessonsMessage(
				parser.ParseByDay(date.Today()+1),
				"Завтра, "+update.Message.From.FirstName+", эти пары:",
				"Завтра пар нет",
			)))

		case commands.Weather:
			w, err := weather.GetWeatherWithCache(opts.Key, 53.346853, 83.777012, time.Hour)
			if err != nil {
				continue
			}
			_, _ = bot.Send(tgbot.NewMessage(chatId, messages.WeatherMessage(w)))

		case commands.NewYear:
			_, _ = bot.Send(tgbot.NewMessage(chatId, messages.NewYearMessage()))

		default:
		}
	}
}
