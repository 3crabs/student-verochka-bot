package messages

import (
	"fmt"
	weather "github.com/3crabs/go-yandex-weather-api"
	"strings"
	"student_bot/parser"
	"time"
)

func StartMessage() string {
	return "Привет, я Верочка!"
}

func HelpMessage() string {
	return "Вот чем я могу вам помочь, отправь:\n" +
		"- /ping и я отобью pong\n" +
		"- /today_lessons и я покажу расписание на сегодня\n" +
		"- /tomorrow_lessons и я покажу расписание на завтра\n" +
		"- /weather и я покажу погоду у универа\n" +
		"\nНу а больше я пока ничего не умею"
}

func PongMessage() string {
	return "pong"
}

func LessonsMessage(schedule []parser.Lesson, prefix, emptyText string) string {
	if len(schedule) == 0 {
		return emptyText
	}

	s := ""
	if prefix != "" {
		s += prefix + "\n\n"
	}
	for _, l := range schedule {
		if strings.Contains(l.User, "Ярных В.В.") {
			s += "⭐️ "
		}
		s += fmt.Sprintf("%s %s (%s)\n", l.Time, l.Name, l.Place)
	}
	if prefix != "" {
		s += "\n" +
			"Д — Димитрова 66 (Филологический, социологический факультеты, отделе)\n" +
			"Л — Ленина 61 (Математический и биологический факультеты)\n" +
			"М — Ленина 61а (Исторический и географический факультеты)"
	}
	return s
}

func WeatherMessage(w weather.Weather) string {
	s := fmt.Sprintf("Сегодня у универа:\n\n"+
		"Температура %d°C\n"+
		"Ощущается как %d°C\n"+
		"Скорость ветра %d°C\n"+
		"Порывы ветра до %d м/с\n",
		w.Fact.Temp,
		w.Fact.FeelsLike,
		w.Fact.WindSpeed,
		w.Fact.WindGust)
	return s
}

func NewYearMessage() string {
	loc := time.FixedZone("UTC+7", +7*60*60)
	ny := time.Date(2022, 1, 1, 0, 0, 0, 0, loc)
	now := time.Now().In(loc)
	days := ny.Sub(now).Hours() / 24
	return toStrDays(int(days))
}

func toStrDays(days int) string {
	if days == 0 {
		return "Сегодня новый год!!!"
	}
	if days == 1 {
		return "Новый год завтра!"
	}
	s := ""
	a := days % 10
	if days < 10 {
		switch a {
		case 2:
			s = "дня"
		case 3:
			s = "дня"
		case 4:
			s = "дня"
		default:
			s = "дней"
		}
	}
	if days >= 10 && days < 20 {
		s = "дней"
	}
	if days >= 20 {
		switch a {
		case 1:
			s = "день"
		case 2:
			s = "дня"
		case 3:
			s = "дня"
		case 4:
			s = "дня"
		default:
			s = "дней"
		}
	}
	return fmt.Sprintf("До нового года %d %s", days, s)
}
