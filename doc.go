package main

import (
	"fmt"
)

func SumKeywordsInTitle(filename string) {
	var ap ArxivPapers
	ap.ParseLargeFileByLine(filename)

	desired := []string{"cs.LG", "cs.AI", "cs.CR", "cs.DB", "cs.IR"}
	ap.stat.ToHtmlChartPeriod(desired, 2015, 2024)
	//ap.stat.PrintResult()
	//ap.PrintResults()
	//ap.PrintItems()

	fmt.Println("======================")
	ap.PrintResults()
}

func CreateWordCloud(filename, category string) {
	var ap ArxivPapers
	ap.SetCategories([]string{category})
	ap.ParseLargeFileByLine(filename)
	ap.GenWordCloud(filename, category)

	//ap.stat.ToHtmlChartPeriod(desired, 2015, 2024)
}
