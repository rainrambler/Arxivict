package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// {"version":"v1","created":"Mon, 2 Apr 2007 19:18:42 GMT"}
type Version struct {
	version string
	created string
}

type Versions struct {
	allVersions []*Version
}

func (p *Versions) Add(ver *Version) {
	if ver == nil {
		return
	}
	p.allVersions = append(p.allVersions, ver)
}

type Author struct {
	names []string
}

type Authors struct {
	allItems []*Author
}

func (p *Authors) Add(item *Author) {
	if item == nil {
		return
	}
	p.allItems = append(p.allItems, item)
}

type ArxivPaper struct {
	id             string
	submitter      string
	authors        string
	title          string
	comments       string
	journalref     string
	doi            string
	reportno       string
	categories     string
	license        string
	abstract       string
	versions       *Versions
	updatedate     string // "update_date":"2008-11-26"
	authors_parsed *Authors
}

type ArxivPapers struct {
	items          []*ArxivPaper
	category2count map[string]int
	subctg2count   map[string]int
	license2count  map[string]int
	key2count      map[string]int

	stat *PaperStatistics
}

func (p *ArxivPapers) Init() {
	p.category2count = make(map[string]int)
	p.subctg2count = make(map[string]int)
	p.license2count = make(map[string]int)
	p.key2count = make(map[string]int)

	p.stat = new(PaperStatistics)
	p.stat.Init()
}

func (p *ArxivPapers) ParseLargeFileByLine(filename string) {
	p.Init()

	f, err := os.Open(filename)
	if err != nil {
		log.Printf("Err: Cannot read file %s: %v\n", filename, err)
		return
	}
	defer f.Close()

	buf := []byte{}
	scanner := bufio.NewScanner(f)
	scanner.Buffer(buf, 1024*1024)

	totallines := 0
	for scanner.Scan() {
		line := scanner.Text()

		content, err := parseLine(line)
		if err != nil {
			log.Printf("Err: Cannot read line %d %s: %v\n", totallines+1,
				line, err)
		}

		paper := convPaper(content.Data)
		if paper != nil {
			p.addPaperMeta(paper)
		} else {
			log.Printf("INFO: Cannot convert paper on line: %d\n", totallines+1)
		}

		totallines++
	}

	log.Printf("INFO: Total %d lines.\n", totallines)
}

func (p *ArxivPapers) addPaperMeta(paper *ArxivPaper) {
	if paper == nil {
		return
	}
	allcats := strings.Split(paper.categories, " ")
	for _, oneCat := range allcats {
		p.subctg2count[oneCat] = p.subctg2count[oneCat] + 1

		parent := getCategory(oneCat)
		p.category2count[parent] = p.category2count[parent] + 1
	}

	p.license2count[paper.license] = p.license2count[paper.license] + 1

	arr := strings.Split(paper.title, " ")
	for _, key := range arr {
		keynew := PurifyKey(key)
		if len(keynew) > 0 {
			p.key2count[keynew] = p.key2count[keynew] + 1
		}
	}

	p.stat.AddOnePaper(paper)
}

func PurifyKey(s string) string {
	sn := strings.Trim(s, " \t\r\n")
	return strings.ToLower(sn)
}

const MinPaperCount = 10

func (p *ArxivPapers) PrintResults() {
	log.Printf("Categories count: %d\n", len(p.category2count))
	for k, v := range p.category2count {
		if v >= MinPaperCount {
			log.Printf("%s:%d\n", k, v)
		}
	}
	log.Println("=============================")
	log.Printf("Sub-categories count: %d\n", len(p.subctg2count))
	for k, v := range p.subctg2count {
		log.Printf("%s:%d\n", k, v)
	}
	log.Println("=============================")
	log.Printf("Licenses count: %d\n", len(p.license2count))
	for k, v := range p.license2count {
		log.Printf("%s:%d\n", k, v)
	}

	PrintMapByValueTop(p.key2count, 10000)
}

func convPaper(content interface{}) *ArxivPaper {
	var paper ArxivPaper
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			paper.id = v.(string)
		case "submitter":
			paper.submitter = toNilOrString(v)
		case "authors":
			paper.authors = v.(string)
		case "title":
			paper.title = v.(string)
		case "comments":
			paper.comments = toNilOrString(v)
		case "journal-ref":
			paper.journalref = toNilOrString(v)
		case "doi":
			paper.doi = toNilOrString(v)
		case "report-no":
			paper.reportno = toNilOrString(v)
		case "categories":
			paper.categories = v.(string)
		case "license":
			paper.license = toNilOrString(v)
		case "abstract":
			paper.abstract = v.(string)
		case "update_date":
			paper.updatedate = v.(string)
		case "versions":
			paper.versions = convertVersions(v)
		case "authors_parsed":
			paper.authors_parsed = convertAuthorsParsed(v)
		default:
			log.Printf("[DBG][convPaper]Unknown [%v]:%v\n", k, v)
		}
	}
	return &paper
}

func convertVersions(content interface{}) *Versions {
	var inst Versions
	arr := content.([]interface{})
	for _, v := range arr {
		ver := convertVersion(v)
		inst.Add(ver)
	}

	return &inst
}

func convertVersion(content interface{}) *Version {
	var inst Version
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "version":
			inst.version = v.(string)
		case "created":
			inst.created = v.(string)
		default:
			log.Printf("[DBG][convertVersion]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertAuthorsParsed(content interface{}) *Authors {
	var inst Authors
	arr := content.([]interface{})
	for _, v := range arr {
		ver := convertAuthor(v)
		inst.Add(ver)
	}

	return &inst
}

func convertAuthor(content interface{}) *Author {
	var inst Author
	arr := content.([]interface{})
	for _, v := range arr {
		inst.names = append(inst.names, v.(string))
	}

	return &inst
}

func toNilOrString(content interface{}) string {
	if content == nil {
		return ""
	}

	return content.(string)
}

// cs.GL ==> cs
// plasm-ph ==> plasm-ph
func getCategory(subcategory string) string {
	idx := strings.Index(subcategory, ".")
	if idx == -1 {
		return subcategory
	}

	arr := strings.Split(subcategory, ".")
	return arr[0] + "]"
}

func isValidCategory(catname string) bool {
	return strings.HasPrefix(catname, "[") &&
		strings.HasSuffix(catname, "]")
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
