package main

import (
	"errors"
	"log"
	"math/rand"
	"os"
	"strings"
	"student_bot/commands"
	"student_bot/date"
	"student_bot/messages"
	"student_bot/new_year_service"
	"student_bot/parser"
	"time"

	weather "github.com/3crabs/go-yandex-weather-api"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/umputun/go-flags"
)

type Opts struct {
	Token string `short:"t" long:"token" description:"Telegram api token"`
	Key   string `short:"k" long:"key" description:"Yandex weather API key"`
	Name  string `short:"n" long:"name" description:"Telegram bot name" default:"@student_verochka_bot"`
}

var sendMessage = false

var opts Opts

func main() {
	run()
}

func run() {
	rand.Seed(time.Now().UnixNano())

	p := flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash|flags.HelpFlag)
	p.SubcommandsOptional = true
	if _, err := p.Parse(); err != nil {
		if err.(*flags.Error).Type != flags.ErrHelp {
			log.Println(errors.New("[ERROR] cli error: " + err.Error()))
		}
		os.Exit(2)
	}

	imageService, err := new_year_service.NewNewYearService()
	if err != nil {
		log.Panicln(err)
	}
	imageService = new_year_service.NewNewYearServiceLogWrapper(imageService)

	bot, err := tgbot.NewBotAPI(opts.Token)
	if err != nil {
		log.Panicln(err)
	}

	u := tgbot.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Bot is start up!")

	for update := range updates {

		// empty message
		if update.Message == nil {
			continue
		}

		text := update.Message.Text
		chatId := update.Message.Chat.ID

		if strings.Contains(strings.ToLower(text), "спасибо") {
			if sendMessage {
				_, _ = bot.Send(tgbot.NewMessage(chatId, messages.ThanksMessage()))
				sendMessage = false
			}
			continue
		}
		sendMessage = false

		command := commands.Command(strings.Replace(text, opts.Name, "", 1))
		switch command {

		case commands.Start:
			send(bot, tgbot.NewMessage(chatId, messages.StartMessage()))

		case commands.Help:
			send(bot, tgbot.NewMessage(chatId, messages.HelpMessage()))

		case commands.Ping:
			send(bot, tgbot.NewMessage(chatId, messages.PongMessage()))

		case commands.TodayLessons:
			send(bot, tgbot.NewMessage(chatId, messages.LessonsMessage(
				parser.ParseByDay(date.Today()),
				"Сегодня, "+update.Message.From.FirstName+", эти пары:",
				"Сегодня пар нет",
			)))

		case commands.TomorrowLessons:
			send(bot, tgbot.NewMessage(chatId, messages.LessonsMessage(
				parser.ParseByDay(date.Today()+1),
				"Завтра, "+update.Message.From.FirstName+", эти пары:",
				"Завтра пар нет",
			)))

		case commands.Weather:
			w, err := weather.GetWeatherWithCache(opts.Key, 53.346853, 83.777012, time.Hour)
			if err != nil {
				continue
			}
			send(bot, tgbot.NewMessage(chatId, messages.WeatherMessage(w)))

		case commands.NewYear:
			url := imageService.GetRandomImageURL()
			message := imageService.GetRandomMessage()
			msg := tgbot.NewPhotoShare(chatId, url)
			msg.Caption = messages.NewYearMessage(message)
			send(bot, msg)

		default:
		}
	}
}

func send(bot *tgbot.BotAPI, msg tgbot.Chattable) {
	_, _ = bot.Send(msg)
	sendMessage = true
}
