package main

import (
	"fmt"
	"log"
	"time"
)

type PaperStatistics struct {
	cate2years map[string]*YearPapers
}

func (p *PaperStatistics) Init() {
	p.cate2years = make(map[string]*YearPapers)
}

const (
	// https://gosamples.dev/date-time-format-cheatsheet/
	// YYYY-MM-DD: 2022-03-23
	YYYYMMDD = "2006-01-02"
)

func (p *PaperStatistics) AddOnePaper(ap *ArxivPaper) {
	if ap == nil {
		return
	}

	t, err := time.Parse(YYYYMMDD, ap.updatedate)
	if err != nil {
		log.Fatal("Cannot convert time %s: %v", ap.updatedate, err)
		return
	}

	dt := t.Year()
	p.AddPaper(dt, ap.title, ap.categories)
}

func (p *PaperStatistics) AddPaper(year int, title, cate string) {
	ps, exists := p.cate2years[cate]
	if exists {
		ps.AddPaper(year, title)
		return
	}

	ps = new(YearPapers)
	ps.year2papers = make(map[int]*Papers)
	ps.AddPaper(year, title)
	p.cate2years[cate] = ps
}

func (p *PaperStatistics) PrintResult() {
	for cate, years := range p.cate2years {
		fmt.Printf("Category: %s\n", cate)

		m := years.GetResult()
		SortPrintYear(m)
	}
}

func SortPrintYear(m map[int]int) {
	PrintSortedMapByKey("year", m)
}

type YearPapers struct {
	year2papers map[int]*Papers
}

func (p *YearPapers) AddPaper(year int, title string) {
	ps, exists := p.year2papers[year]
	if exists {
		ps.AddPaper(title)
		return
	}

	ps = new(Papers)
	ps.paper2int = make(map[string]int)
	ps.AddPaper(title)
	p.year2papers[year] = ps
}

// Returns: Key: Year, Value: Paper count
func (p *YearPapers) GetResult() map[int]int {
	year2num := map[int]int{}

	for year, papers := range p.year2papers {
		year2num[year] = papers.Count()
	}

	return year2num
}

type Papers struct {
	paper2int map[string]int // Key: paper name, value: no use
}

func (p *Papers) AddPaper(title string) bool {
	_, exists := p.paper2int[title]
	if exists {
		return false
	}

	p.paper2int[title] = 1
	return true
}

func (p *Papers) Count() int {
	return len(p.paper2int)
}
