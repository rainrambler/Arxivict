package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Version struct {
	version string `json:"version"`
	created string `json:"created"`
}

type Author struct {
	names []string
	//First  string
	//Middle string
	//Last   string
}

type ArxivPaper struct {
	id             string    `json:"id"`
	submitter      string    `json:"submitter"`
	authors        string    `json:"authors"`
	title          string    `json:"title"`
	comments       string    `json:"comments,omitempty"`
	journalref     string    `json:"journal-ref,omitempty"`
	doi            string    `json:"doi,omitempty"`
	reportno       string    `json:"report-no,omitempty"`
	categories     string    `json:"categories"`
	license        string    `json:"license,omitempty"`
	abstract       string    `json:"abstract"`
	versions       []Version `json:"versions"`
	updatedate     string    `json:"update_date"`
	authors_parsed []Author  `json:"authors_parsed"`
}

type ArxivPapers struct {
	items []ArxivPaper
}

// https://www.golinuxcloud.com/golang-json-unmarshal/
func (p *ArxivPapers) ReadFile(filename string) {
	p.items = []ArxivPaper{}

	jsonfile, err := os.Open(filename)
	if err != nil {
		log.Printf("[WARN]Cannot open file: %s: %v\n", filename, err)
		return
	}

	defer jsonfile.Close()
	bytes, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		log.Printf("[WARN]Cannot read file: %s: %v\n", filename, err)
		return
	}
	arr := []ArxivPaper{}

	err = json.Unmarshal(bytes, &arr)
	if err != nil {
		log.Printf("[WARN]Cannot unmarshal file: %s: %v\n", filename, err)
		return
	}

	p.items = append(p.items, arr...)
}

// https://www.golinuxcloud.com/golang-json-unmarshal/
func (p *ArxivPapers) ReadFile2(filename string) {
	js, err := parseFile(filename)
	if err != nil {
		log.Printf("[WARN]Cannot open file: %s: %v\n", filename, err)
		return
	}

	log.Printf("%+v\n", js)
}

func (p *ArxivPapers) PrintItems() {
	if len(p.items) == 0 {
		return
	}
	//log.Printf("Papers: %+v\n", p.items)
	log.Printf("Total %d items: \n", len(p.items))
	for _, item := range p.items {
		log.Printf("%v\n", item)
	}
}
