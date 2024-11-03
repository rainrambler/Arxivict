package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// ReadLines reads a whole file into memory and returns a slice of its lines
// for whole file as a string, use ioutil.ReadFile instead.
func ReadLines(fullpath string) ([]string, error) {
	f, err := os.Open(fullpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ReadLines reads a whole file into memory and returns a slice of its lines
// for whole file as a string, use ioutil.ReadFile instead.
func ReadLinesLarge(fullpath string) ([]string, error) {
	f, err := os.Open(fullpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	buf := []byte{}
	scanner := bufio.NewScanner(f)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadLinesInLargeFile(filename string) ([]string, error) {
	readFile, err := os.Open(filename)

	if err != nil {
		//log.Printf("Cannot read file %s: %v\n", filename, err)
		return []string{}, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines, nil
}

func ReadLinesInLargeFile2(filename string) ([]string, error) {
	readFile, err := os.Open(filename)

	if err != nil {
		//log.Printf("Cannot read file %s: %v\n", filename, err)
		return []string{}, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines, nil
}

// WriteLines - each line should NOT end with \n or \r\n
func WriteLines(lines []string, fullpath string) error {
	f, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	defer f.Close()

	//fmt.Printf("[DBG]File: [%s] %+v\n", fullpath, lines)

	w := bufio.NewWriter(f)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}

	return w.Flush()
}

func IoReader(fullpath string) io.ReaderAt {
	r, err := os.Open(fullpath)
	if err != nil {
		panic(err)
	}
	return r
}

func ReadBinFile(fullpath string) ([]byte, error) {
	return ioutil.ReadFile(fullpath)
}

func ReadTextFile(fullpath string) (string, error) {
	content, err := ioutil.ReadFile(fullpath)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(content), nil
}

func ReadTxtFile(filename string) []string {
	linearr := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linearr = append(linearr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return linearr
}

func WriteTextFile(fullpath, content string) error {
	err := ioutil.WriteFile(fullpath, []byte(content), 0777)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// See: https://golangbyexample.com/append-file-golang
func AppendTextFile(fullpath, content string) error {
	file, err := os.OpenFile(fullpath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.WriteString(content); err != nil {
		return err
	}

	return nil
}
