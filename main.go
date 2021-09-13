package main

import (
	"fmt"
	"github.com/klarkxy/gohtml"
)
/* CSS Document */
var tbspace =
`
cellspacing:4;
cellpadding:6;
`
var tbborder =
`
  border: 1px solid black;
  border-collapse: collapse;
  white-space:nowrap;
`
var biaoti =
`
	font-family: 微软雅黑;
	font-size: 20px;
	font-weight: bold;
	color: #255e95;
`
var titfont = "font-family: 微软雅黑;font-size: 12px;font-weight: bold; background-color:#e9e9e9;"

var trfont = "font-family: 微软雅黑;font-size: 8px;;font-weight: normal;"

func TableHtml(caption string, h []string, t [][]string) string {
	table := gohtml.Tag("table").Attr("style", tbborder+tbspace)
	
	table.Tag("caption").Attr("style", biaoti).Text(caption)
	tr := table.Tag("tr").Attr("style", tbborder + titfont)
	for _, col := range h {
		tr.Tag("th").Attr("style", tbborder).Text(col)
	}

	for _, row := range t {
		tr := table.Tag("tr").Attr("style", tbborder)
		for _, col :=range row {
			tr.Tag("th").Attr("style", tbborder + trfont).Text(col)
		}
	}

	return table.String()
}

func main() {
	// Step 1: 创建数组
	values := [][]string{}

	// Step 2: 使用 appped() 函数向空的二维数组添加两行一维数组
	row1 := []string{"2021-09-23","2.43345","3.234"}
	row2 := []string{"2021-09-24","2.53245","5.224"}
	values = append(values, row1)
	values = append(values, row2)

	h := []string{"DATE", "RMBUSD", "RMBEUR"}
	fmt.Println(TableHtml("汇率展示", h ,values))
}
