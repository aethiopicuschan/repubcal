package repubcal

import (
	"time"

	"github.com/aethiopicuschan/romannum"
)

type Date struct {
	Year     string
	YearInt  int
	Month    string
	MonthInt int
	Day      string
	DayInt   int
	Weekday  string
}

func NewDate(src time.Time) (date Date, err error) {
	zero := time.Date(1792, 9, 22, 0, 0, 0, 0, time.UTC)
	diff := int(src.Sub(zero).Hours() / 24)
	year := 1
	days := daysOf(year)
	for diff >= days {
		diff -= days
		year += 1
		days = daysOf(year)
	}
	date.YearInt = year
	date.Year, err = romannum.ToRoman(year)
	if err != nil {
		return
	}
	date.MonthInt = int(1 + (diff / 30))
	date.Month = monthList()[date.MonthInt-1]
	date.DayInt = int(1 + (diff % 30))
	date.Day = dayList()[30*(date.MonthInt-1)+(date.DayInt-1)]
	weekday := date.DayInt%10 - 1
	if weekday == -1 {
		weekday = 9
	}
	date.Weekday = weekdayList()[weekday]
	return
}

func isLeapYear(year int) bool {
	list := []int{3, 7, 11, 15, 20}
	for _, l := range list {
		if year == l {
			return true
		}
	}
	return year > 20 && (year%4 == 0 && (year%100 > 0) || year%400 == 0)
}

func daysOf(year int) int {
	if isLeapYear(year) {
		return 366
	} else {
		return 365
	}
}
