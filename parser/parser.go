package parser

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type Lesson struct {
	Day    int
	Number string
	Time   string
	Name   string
	User   string
	Place  string
}

func ParseByDay(day int) []Lesson {
	return selectLessonsByDay(parse(), day)
}

func parse() []Lesson {
	url := "https://www.asu.ru/timetable/students/12/2129440415/"

	// Get html
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	k := 0
	var lessons []Lesson
	var l Lesson
	doc.Find(".schedule tr>td, .schedule tr>td[width=\"100%\"]").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if strings.Contains(text, "Понедельник") ||
			strings.Contains(text, "Вторник") ||
			strings.Contains(text, "Среда") ||
			strings.Contains(text, "Четверг") ||
			strings.Contains(text, "Пятница") ||
			strings.Contains(text, "Суббота") {
			k = 0
			l = Lesson{}
			if strings.Contains(text, "Понедельник") {
				l.Day = 1
			}
			if strings.Contains(text, "Вторник") {
				l.Day = 2
			}
			if strings.Contains(text, "Среда") {
				l.Day = 3
			}
			if strings.Contains(text, "Четверг") {
				l.Day = 4
			}
			if strings.Contains(text, "Пятница") {
				l.Day = 5
			}
			if strings.Contains(text, "Суббота") {
				l.Day = 6
			}
		}
		n := (k - 1) % 6
		if n == 0 {
			l.Number = trim(text)
		}
		if n == 1 {
			l.Time = trim(text)
		}
		if n == 2 {
			l.Name = trim(text)
		}
		if n == 3 {
			l.User = trim(text)
		}
		if n == 4 {
			l.Place = trim(text)
			if l.Place == "" {
				l.Place = "не приходи"
			}
			lessons = append(lessons, l)
		}
		k++
	})
	return lessons
}

func trim(s string) string {
	s = strings.Replace(s, "пр.з.", "пр.", -1)
	s = strings.Replace(s, "  ", "", -1)
	s = strings.Replace(s, "\n", " ", -1)
	return s
}

func selectLessonsByDay(schedule []Lesson, day int) []Lesson {
	var todayLessons []Lesson
	for _, l := range schedule {
		if l.Day == day {
			todayLessons = append(todayLessons, l)
		}
	}
	return todayLessons
}
