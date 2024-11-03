package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
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
		log.Printf("ERR: Cannot convert time %s: %v", ap.updatedate, err)
		return
	}

	dt := t.Year()

	arr := strings.Split(ap.categories, " ")
	for _, cate := range arr {
		cate1 := strings.Trim(cate, " \t")
		if cate1 != "" {
			p.AddPaper(dt, ap.title, cate1)
		} else {
			log.Printf("INFO: Cannot parse categories %s for %s\n",
				ap.categories, ap.title)
		}
	}
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
	var ar Archives
	ar.Init()

	total := 0
	for cate, years := range p.cate2years {
		catedesc, exists := ar.Find(cate)
		m := years.GetResult()
		cnt := years.CountPapers()
		total += cnt
		if exists {
			fmt.Printf("Category [%s]: %s, %d papers\n",
				cate, catedesc.name, cnt)
		} else {
			fmt.Printf("Category: %s, %d papers\n",
				cate, cnt)
		}
		SortPrintYear(m)
	}

	log.Printf("Total %d papers.\n", total)
}

func (p *PaperStatistics) ToHtmlChart(desiredcates []string) {
	var ar Archives
	ar.Init()

	yearmin := 3000
	yearmax := 0
	for _, cate := range desiredcates {
		ys := p.cate2years[cate]

		if ys == nil {
			log.Printf("INFO: Cannot find category: %s\n", cate)
			continue
		}

		curmin, curmax := ys.GetMinMax()
		if yearmin > curmin {
			yearmin = curmin
		}

		if yearmax < curmax {
			yearmax = curmax
		}
	}

	var mt Matrix
	mt.Init()

	years := GetYears(yearmin, yearmax)
	mt.SetRows(years)
	mt.SetColumns(desiredcates)

	for _, cate := range desiredcates {
		ys := p.cate2years[cate]

		if ys == nil {
			log.Printf("INFO: Cannot find category 2: %s\n", cate)
			continue
		}

		for y := yearmin; y <= yearmax; y++ {
			papercount := ys.GetCount(y)

			mt.SetValue(strconv.Itoa(y), cate, strconv.Itoa(papercount))
		}
	}

	mt.PrintDesc()
}

func GetYears(yearmin, yearmax int) []string {
	years := []string{}
	for y := yearmin; y <= yearmax; y++ {
		years = append(years, strconv.Itoa(y))
	}
	return years
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

func (p *YearPapers) GetMinMax() (int, int) {
	keys := make([]int, 0, len(p.year2papers))

	for k := range p.year2papers {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return keys[0], keys[len(keys)-1]
}

func (p *YearPapers) GetCount(year int) int {
	ps, exists := p.year2papers[year]
	if !exists {
		return 0
	}

	return ps.Count()
}

func (p *YearPapers) CountPapers() int {
	num := 0
	for _, paper := range p.year2papers {
		num += paper.Count()
	}
	return num
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
