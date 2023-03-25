// Arxivict project main.go
package main

func main() {
	var ap ArxivPapers
	ap.ReadFile(`arxivdemo.json`)
	ap.PrintItems()
}
