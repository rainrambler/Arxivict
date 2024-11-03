package main

import (
	"log"
)

type Matrix struct {
	MaxRow   int
	MaxCol   int
	Rows     []string
	Columns  []string
	grid2Val map[int]string
	row2num  map[string]int
	col2num  map[string]int
}

func (p *Matrix) Init() {
	p.Columns = []string{}
	p.Rows = []string{}
	p.MaxCol = 0
	p.MaxRow = 0

	p.grid2Val = make(map[int]string)
	p.row2num = make(map[string]int)
	p.col2num = make(map[string]int)
}

func (p *Matrix) SetRows(arr []string) {
	for i, v := range arr {
		p.Rows = append(p.Rows, v)
		p.row2num[v] = i
	}
	p.MaxRow = len(p.Rows)
}

func (p *Matrix) SetColumns(arr []string) {
	for i, v := range arr {
		p.Columns = append(p.Columns, v)
		p.col2num[v] = i
	}
	p.MaxCol = len(p.Columns)
}

func (p *Matrix) SetValue(row, col, v string) bool {
	rownum, exists := p.row2num[row]
	if !exists {
		log.Printf("Cannot save [%s, %s]: %s, %s not found!\n",
			row, col, v, row)
		return false
	}

	colnum, exists := p.col2num[col]
	if !exists {
		log.Printf("Cannot save [%s, %s]: %s, %s not found!\n",
			row, col, v, col)
		return false
	}

	return p.setValueInner(rownum, colnum, v)
}

func (p *Matrix) GetValue(row, col string) string {
	rownum, exists := p.row2num[row]
	if !exists {
		log.Printf("Cannot find [%s, %s], %s not found!\n",
			row, col, row)
		return ""
	}

	colnum, exists := p.col2num[col]
	if !exists {
		log.Printf("Cannot find [%s, %s], %s not found!\n",
			row, col, col)
		return ""
	}

	return p.getValueInner(rownum, colnum)
}

const MaxColNum = 100

func (p *Matrix) setValueInner(row, col int, v string) bool {
	if (row < 0) || (col < 0) {
		log.Printf("INFO: Invalid in matrix: [%d, %d]: %s\n",
			row, col, v)
		return false
	}
	grid := row*MaxColNum + col

	p.grid2Val[grid] = v
	return true
}

func (p *Matrix) getValueInner(row, col int) string {
	if (row < 0) || (col < 0) {
		log.Printf("INFO: Invalid in matrix: [%d, %d]\n",
			row, col)
		return ""
	}
	grid := row*MaxColNum + col

	v, exists := p.grid2Val[grid]
	if !exists {
		return ""
	}

	return v
}
