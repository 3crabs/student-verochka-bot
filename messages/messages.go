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
	return "Вот чем я могу вам помочь, отправь:\n\n" +
		"- /ping отобью pong\n" +
		"- /today_lessons покажу расписание на сегодня\n" +
		"- /tomorrow_lessons покажу расписание на завтра\n" +
		"- /weather покажу погоду у универа\n" +
		"- /new_year - посчитаю количество дней до нового года\n" +
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
			"Д — Димитрова 66 (Филологический, социологический)\n" +
			"Л — Ленина 61 (Математический и биологический)\n" +
			"М — Ленина 61а (Исторический и географический)\n" +
			"С — Социалистический 68 (Экономический и юридический)"
	}
	return s
}

func WeatherMessage(w weather.Weather) string {
	return fmt.Sprintf("Сейчас у универа:\n"+
		"%s\n"+
		"Температура: %d°C\n"+
		"Ощущается как: %d°C\n"+
		"Скорость ветра: %dм/с\n"+
		"\nПогода на %s:\n"+
		"%s\n"+
		"Температура: %d°C\n"+
		"Ощущается как: %d°C\n"+
		"Скорость ветра: %dм/с\n",
		w.Fact.GetCondition(),
		w.Fact.Temp,
		w.Fact.FeelsLike,
		w.Fact.WindSpeed,
		w.Forecast.Parts[0].GetPartName(),
		w.Forecast.Parts[0].GetCondition(),
		w.Forecast.Parts[0].TempAvg,
		w.Forecast.Parts[0].FeelsLike,
		w.Forecast.Parts[0].WindSpeed)
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
