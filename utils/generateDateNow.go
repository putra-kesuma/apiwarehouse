package utils

import "time"

func  GenDateNow() string  {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 3:4:5")
}
