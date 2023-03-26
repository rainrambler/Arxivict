package main

import (
	"log"
)

type Version struct {
	version string `json:"version"`
	created string `json:"created"`
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
	updatedate     string
	authors_parsed *Authors
}

type ArxivPapers struct {
	items []*ArxivPaper
}

// https://www.golinuxcloud.com/golang-json-unmarshal/
func (p *ArxivPapers) ReadFile(filename string) {
	js, err := parseFile(filename)
	if err != nil {
		log.Printf("[WARN]Cannot open file: %s: %v\n", filename, err)
		return
	}

	p.convert(js)
	//log.Printf("%+v\n", js)
}

func (p *ArxivPapers) convert(jc *JsonContent) {
	content := jc.Data
	arr := content.([]interface{})
	for _, v := range arr {
		paper := convPaper(v)
		if paper != nil {
			p.AddPaper(paper)

			//log.Printf("[DBG]Add paper: %+v\n", *paper)
		}
	}
	//log.Printf("Content: %v\n", content)
	log.Printf("Total %d papers.\n", len(p.items))
}

func (p *ArxivPapers) AddPaper(paper *ArxivPaper) {
	if paper == nil {
		return
	}
	p.items = append(p.items, paper)
}

func convPaper(content interface{}) *ArxivPaper {
	paper := new(ArxivPaper)
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			paper.id = v.(string)
		case "submitter":
			paper.submitter = v.(string)
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
	return paper
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
