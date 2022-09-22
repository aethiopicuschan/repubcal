package repubcal_test

import (
	"testing"
	"time"

	"github.com/aethiopicuschan/repubcal"
)

func TestNewDate(t *testing.T) {
	// Coup d'état of 9 Thermidor
	date, err := repubcal.NewDate(time.Date(1794, 7, 27, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("%s", err)
	}
	if date.Year != "II" {
		t.Errorf("want \"II\", got %s", date.Year)
	}
	if date.YearInt != 2 {
		t.Errorf("want 2, got %d", date.YearInt)
	}
	if date.Month != "Thermidor" {
		t.Errorf("want \"Thermidor\", got %s", date.Month)
	}
	if date.MonthInt != 11 {
		t.Errorf("want 11, got %d", date.MonthInt)
	}
	if date.Day != "Mûre" {
		t.Errorf("want \"Mûre\", got %s", date.Day)
	}
	if date.DayInt != 9 {
		t.Errorf("want 9, got %d", date.DayInt)
	}
	if date.Weekday != "Nonidi" {
		t.Errorf("want \"Nonidi\", got %s", date.Weekday)
	}
}
