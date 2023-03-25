// Arxivict project main.go
package main

func main() {
	var ap ArxivPapers
	ap.ReadFile2(`arxivdemo2.json`)
	ap.PrintItems()
}
