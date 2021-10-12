package messages

import (
	"fmt"
	weather "github.com/3crabs/go-yandex-weather-api"
	"strings"
	"student_bot/parser"
)

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

func StartMessage() string {
	return "Привет, я Верочка!"
}

func PongMessage() string {
	return "pong"
}

// HelpMessage
//today_lessons - пары сегодня
//tomorrow_lessons - пары завтра
//weather - погода у универа
//ping - проверка связи
//help - помощь
func HelpMessage() string {
	return "Вот чем я могу вам помочь, отправь:\n" +
		"- /ping и я отобью pong\n" +
		"- /today_lessons и я покажу расписание на сегодня\n" +
		"- /tomorrow_lessons и я покажу расписание на завтра\n" +
		"- /weather и я покажу погоду у универа\n" +
		"\nНу а больше я пока ничего не умею"
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
