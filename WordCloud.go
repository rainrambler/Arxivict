package main

import (
	"fmt"
	"log"
	"strings"
)

type WordCloud struct {
	word2count map[string]int
	filtered   map[string]int // Key: words that should be filtered. e.g. "the"
}

func (p *WordCloud) InitParams() {
	p.word2count = make(map[string]int)
	p.filtered = make(map[string]int)
}

func (p *WordCloud) loadConfig(filename string) {
	arr, err := ReadLines(filename)
	if err != nil {
		log.Printf("WARN: Cannot read file %s: %v\n", filename, err)
		return
	}

	for _, item := range arr {
		if len(item) > 0 {
			p.filtered[item] = 1
		}
	}
}

func (p *WordCloud) shouldFilter(enword string) bool {
	if len(enword) == 0 {
		return true
	}
	_, exists := p.filtered[enword]
	return exists
}

func (p *WordCloud) parseSentenceEn(line string) {
	arr := strings.Split(line, " \t")
	for _, item := range arr {
		if item == "" {
			continue
		}

		p.AddWord(item)
	}
}

func (p *WordCloud) AddWord(s string) {
	if p.shouldFilter(s) {
		return
	}
	p.word2count[s] = p.word2count[s] + 1
}

func (p *WordCloud) AddWords(s2c map[string]int) {
	p.InitParams()
	p.loadConfig(`EnCommonWords.txt`)

	for k, v := range s2c {
		if !p.shouldFilter(k) {
			p.word2count[k] = v
		}
	}
}

// Print Top N values (sorted by value), -1 means all
func (p *WordCloud) PrintResult(topn int) {
	if len(p.word2count) == 0 {
		fmt.Println("No result!")
		return
	}
	PrintMapByValueTop(p.word2count, topn)
}

const Multiply = 30

func ConvertJsonHardCode(s2c map[string]int, margin int) string {
	s := ""
	for k, v := range s2c {
		if v > margin {
			line := fmt.Sprintf(`{name:"%s",value:%d},`, k, v)
			s += line
		}
	}

	s = s[:len(s)-1] // remove last comma
	s = "[" + s + "]"
	return s
}

func (p *WordCloud) SaveFile(filename string) {
	tmpl, err := ReadTextFile(`./doc/wordcloudtempl.html`)
	if err != nil {
		fmt.Println("Cannot read template file!")
		return
	}

	for i := 4; i < 30; i++ {
		s := ConvertJsonHardCode(p.word2count, i)
		content := strings.Replace(tmpl, `[$REALDATA$]`, s, 1)
		fullfname := fmt.Sprintf("%s_2_%d.html", filename, i)
		WriteTextFile(fullfname, content)
	}
}

const KEY_MARGIN = 500

func (p *WordCloud) SaveOneFile(filename string) {
	tmpl, err := ReadTextFile(`./doc/wordcloudtempl.html`)
	if err != nil {
		fmt.Println("Cannot read template file!")
		return
	}

	i := KEY_MARGIN * 2
	s := ConvertJsonHardCode(p.word2count, i)
	content := strings.Replace(tmpl, `[$REALDATA$]`, s, 1)
	fullfname := fmt.Sprintf("%s_2_%d.html", filename, i)
	WriteTextFile(fullfname, content)
}

func (p *WordCloud) SaveMultiFiles(filepartname string) {
	tmpl, err := ReadTextFile(`./doc/wordcloudtempl.html`)
	if err != nil {
		fmt.Println("Cannot read template file!")
		return
	}

	for k := 1; k < 6; k++ {
		i := KEY_MARGIN * k
		s := ConvertJsonHardCode(p.word2count, i)
		content := strings.Replace(tmpl, `[$REALDATA$]`, s, 1)
		fullfname := fmt.Sprintf("%s_2_%d.html", filepartname, i)
		WriteTextFile(fullfname, content)
	}
}
