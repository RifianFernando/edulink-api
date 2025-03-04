package lib

import "time"

func AddDays(days int) time.Time {
	return time.Now().Local().Add(time.Hour * time.Duration(24*days))
}

func GetTimeNow() time.Time {
	return time.Now().Local()
}

func CustomTimeDay(days int) time.Time {
	return GetTimeNow().Add(time.Hour * time.Duration(24*days))
}
