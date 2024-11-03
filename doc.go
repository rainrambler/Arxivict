package main

func SumKeywordsInTitle(filename string) {
	var ap ArxivPapers
	ap.ParseLargeFileByLine(filename)

	desired := []string{"cs.LG", "cs.AI", "cs.CR", "cs.DB", "cs.IR"}
	ap.stat.ToHtmlChart(desired)
	//ap.stat.PrintResult()
	//ap.PrintResults()
	//ap.PrintItems()

}
