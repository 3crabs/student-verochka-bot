package commands

type Command string

//today_lessons - пары сегодня
//tomorrow_lessons - пары завтра
//weather - погода у универа
//new_year - дней до нового года
//ping - проверка связи
//help - помощь
const (
	Start           = Command("/start")
	Help            = Command("/help")
	Ping            = Command("/ping")
	TodayLessons    = Command("/today_lessons")
	TomorrowLessons = Command("/tomorrow_lessons")
	Weather         = Command("/weather")
	NewYear         = Command("/new_year")
)
