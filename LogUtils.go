package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func generateLogFileName() string {
	t := time.Now()

	datepart := t.Format("2006-01-02-15-04-05")
	return "log" + datepart + ".txt"
}

func logSettings() {
	logfilename := filepath.Join("log", generateLogFileName())
	f, err := os.OpenFile(logfilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	} else {
		log.Printf("Log file %s created and opened!\n", logfilename)
	}
	defer f.Close()

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
}
