package main

import (
	"log"
)

type Version struct {
	version string `json:"version"`
	created string `json:"created"`
}

type Author struct {
	names []string
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
