package main

import (
	"errors"
	"time"
)

const (
	FixedDay = iota
	NthWeekday
)

type Holiday struct {
	Type                     int
	Month                    int
	DayOfMonth               int
	DayOfWeek                int
	NthOfMonth               int
	ObservedOnClosestWeekday bool
}

// GetHolidays calculates and returns the holidays in the time range
func GetHolidays(checkoutDate time.Time, rentalDays int) (map[time.Time]bool, error) {
	holidayDefinitions := []Holiday{
		{FixedDay, 7, 4, 0, 0, true},
		{NthWeekday, 9, 0, 2, 1, false},
	}

	// map that has key=Time for holiday existence
	holidays := make(map[time.Time]bool)

	// calculate the year of the return date
	endYear := checkoutDate.AddDate(0, 0, rentalDays).Year()

	// loop through the holidays and years and add to the map
	for _, holiday := range holidayDefinitions {
		for year := checkoutDate.Year(); year <= endYear; year++ {
			switch holiday.Type {
			case FixedDay:
				{
					date := time.Date(year, time.Month(holiday.Month), holiday.DayOfMonth, 0, 0, 0, 0, time.Local)
					holidays[adjustForWeekendObservance(date, holiday.ObservedOnClosestWeekday)] = true
				}
			case NthWeekday:
				{
					date := time.Date(year, time.Month(holiday.Month), 1, 0, 0, 0, 0, time.Local)
					weekdayOffset := holiday.DayOfWeek - int(date.Weekday())

					// adjust date based on day-of-week difference from the desired weekday
					date = date.AddDate(0, 0, (weekdayOffset+7)%7)

					// adjust date based on which day-of-week is desired (eg the 2nd monday)
					date = date.AddDate(0, 0, 7*(holiday.NthOfMonth-1))

					if int(date.Month()) == holiday.Month {
						holidays[adjustForWeekendObservance(date, holiday.ObservedOnClosestWeekday)] = true
					} else {
						// TODO: how should an illegal holiday configuration be handled?
						//       eg "the 10th Sunday of the month"
						//       currently do nothing with the assumption that such a configuration
						//       doesn't exist and that any newly added holidays will be tested
					}
				}
			default:
				{
					return nil, errors.New("unsupported holiday type configuration")
				}
			}
		}
	}

	return holidays, nil
}

// adjustForWeekendObservance - adjusts one day back/ahead if the given
// day is a weekend day
func adjustForWeekendObservance(date time.Time, observedOnClosestWeekday bool) time.Time {
	if observedOnClosestWeekday {
		if date.Weekday() == time.Saturday {
			return date.AddDate(0, 0, -1)
		} else if date.Weekday() == time.Sunday {
			return date.AddDate(0, 0, 1)
		} else {
			return date
		}
	} else {
		return date
	}
}
