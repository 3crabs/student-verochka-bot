package date

import "time"

func Today() int {
	return int(time.Now().Weekday())
}
