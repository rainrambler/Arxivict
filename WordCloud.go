package main

import (
	"fmt"
	"strings"
)

type WordCloud struct {
	char2count map[string]int
	word2count map[string]int
}

func (p *WordCloud) InitParams() {
	p.char2count = make(map[string]int)
	p.word2count = make(map[string]int)
}

func (p *WordCloud) parseSentence(line string) {
	rs := []rune(line)
	for _, r := range rs {
		p.AddChar(r)
	}

	rcount := len(rs)
	for i := 0; i < rcount-1; i++ {
		pair := rs[i : i+2]
		p.AddWord(string(pair))
	}
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

func (p *WordCloud) AddChar(r rune) {
	s := string(r)
	p.char2count[s] = p.char2count[s] + 1
}

func (p *WordCloud) AddWord(s string) {
	p.word2count[s] = p.word2count[s] + 1
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

const KEY_MARGIN = 1000

func (p *WordCloud) SaveOneFile(filename string) {
	tmpl, err := ReadTextFile(`./doc/wordcloudtempl.html`)
	if err != nil {
		fmt.Println("Cannot read template file!")
		return
	}

	i := KEY_MARGIN
	s := ConvertJsonHardCode(p.word2count, i)
	content := strings.Replace(tmpl, `[$REALDATA$]`, s, 1)
	fullfname := fmt.Sprintf("%s_2_%d.html", filename, i)
	WriteTextFile(fullfname, content)
}
