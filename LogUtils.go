package main

import (
	"time"
)

func generateLogFileName() string {
	t := time.Now()

	datepart := t.Format("2006-01-02-15-04-05")
	return "log" + datepart + ".txt"
}
