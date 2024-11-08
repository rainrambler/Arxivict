package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	//logSettings()
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

	log.Printf("INFO: Program started at %v\n", time.Now())

	//SumKeywordsInTitle(`arxiv-metadata-oai-snapshot.json`)

	CreateWordCloud(`arxiv-metadata-oai-snapshot.json`, "cs.AI")

	log.Printf("INFO: Program ended at %v\n", time.Now())
}
