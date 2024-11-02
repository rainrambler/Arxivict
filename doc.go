package main

func SumKeywordsInTitle(filename string) {
	var ap ArxivPapers
	ap.ParseLargeFile(filename)
	ap.stat.PrintResult()
	//ap.PrintResults()
	//ap.PrintItems()

}
