package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
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

func (p *Matrix) PrintDesc() {
	for i := 0; i < p.MaxCol; i++ {
		colname := p.Columns[i]
		for j := 0; j < p.MaxRow; j++ {
			rolname := p.Rows[j]

			v := p.getValueInner(j, i)
			fmt.Printf("[%s, %s]:%s\n", rolname, colname, v)
		}
	}
}

func (p *Matrix) convSeries() []interface{} {
	all := []interface{}{}

	for i := 0; i < p.MaxCol; i++ {
		colname := p.Columns[i]

		varr := []int{}
		for j := 0; j < p.MaxRow; j++ {
			v := p.getValueInner(j, i)

			vint, _ := strconv.Atoi(v)
			varr = append(varr, vint)
		}

		series1 := map[string]interface{}{
			"name": colname,
			"type": "line",
			//"data": ArrToJsonStr(varr),
			"data": varr,
		}

		all = append(all, series1)
	}

	return all
}

func ArrToJsonStr(arr []int) string {
	la := len(arr)
	if la == 0 {
		return "[]"
	} else if la == 1 {
		return fmt.Sprintf("[%d]", arr[0])
	}

	s := "["
	for i := 0; i < la-1; i++ {
		s += strconv.Itoa(arr[i]) + ", "
	}

	s += strconv.Itoa(arr[la-1]) + "]" // last item
	return s
}

func (p *Matrix) ToChart() {
	data := map[string]interface{}{
		"title": map[string]interface{}{
			"text": "Arxiv Statistics",
		},
		"tooltip": map[string]interface{}{
			"trigger": "axis",
		},
		"legend": map[string]interface{}{
			"data": p.Columns,
		},
		"grid": map[string]interface{}{
			"left":         "3%",
			"right":        "4%",
			"bottom":       "3%",
			"containLabel": true,
		},
		"xAxis": map[string]interface{}{
			"type":        "category",
			"boundaryGap": false,
			"data":        p.Rows,
		},
		"yAxis": map[string]interface{}{
			"type": "value",
		},

		"series": []interface{}{},
	}

	data["series"] = p.convSeries()

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Printf("json data: %s\n", jsonData)
}
