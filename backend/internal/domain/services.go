package domain

import "time"

func GetDayRangeFromDate(date time.Time) (startOfTheDay time.Time, endOfTheDay time.Time) {
	startOfTheDay = date.Truncate(24 * time.Hour)
	endOfTheDay = startOfTheDay.Add(24 * time.Hour)

	return startOfTheDay, endOfTheDay
}
