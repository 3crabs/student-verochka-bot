package messages

import (
	"fmt"
	"strings"
	"student_bot/parser"
)

func LessonsMessage(schedule []parser.Lesson, prefix, emptyText string) string {
	if len(schedule) == 0 {
		return emptyText
	}
	s := prefix + "\n\n"
	for _, l := range schedule {
		if strings.Contains(l.User, "Ярных В.В.") {
			s += "⭐️ "
		}
		s += fmt.Sprintf("%s %s (%s)\n", l.Time, l.Name, l.Place)
	}
	return s + "\n" +
		"Д — Димитрова 66 (Филологический, социологический факультеты, отделе)\n" +
		"Л — Ленина 61 (Математический и биологический факультеты)\n" +
		"М — Ленина 61а (Исторический и географический факультеты)"
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
//ping - проверка связи
//help - помощь
func HelpMessage() string {
	return "Вот чем я могу вам помочь, отправь:\n" +
		"- /ping и я отобью pong\n" +
		"- /today_lessons и я покажу расписание на сегодня\n" +
		"- /tomorrow_lessons и я покажу расписание на завтра\n" +
		"\nНу а больше я пока ничего не умею"
}
