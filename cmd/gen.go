package cmd

import (
	"fmt"
	"github.com/kzkzzzz/dbtogo/common"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type (
	ColumnInfo struct {
		Table   string
		Name    string
		Comment string
		Type    string
		GoName  string
		GoType  string
	}
	Gen interface {
		GetColumns() []*ColumnInfo
	}
)

func Run(gen Gen) {
	columns := gen.GetColumns()
	tc := make(map[string][]*ColumnInfo, 0)
	for _, column := range columns {
		if column.Comment == "" {
			column.Comment = column.GoName
		}
		tc[column.Table] = append(tc[column.Table], column)
	}

	//fmt.Println(tc)

	line := strings.Repeat("-", 16)

	for _, table := range cmdParam.Table {
		tColumns, ok := tc[table]
		if !ok {
			continue
		}
		tableCamel := common.StrToCamelCase(table)
		build := make([]string, 0)
		build = append(build, fmt.Sprintf("package main\ntype %s struct {\n", tableCamel))
		for _, column := range tColumns {
			build = append(build, fmt.Sprintf(
				" %s\t%s `json:\"%s\"` // %s\n",
				column.GoName, column.GoType, column.Name, column.Comment,
			))
		}
		build = append(build, fmt.Sprintf("}\n"))

		tableNameFunc := fmt.Sprintf("func (%s *%s) TableName() string {\n return \"%s\" \n}",
			strings.ToLower(tableCamel[:1]), tableCamel, table,
		)
		//fmt.Println(tableNameFunc)
		build = append(build, tableNameFunc)

		code := strings.Join(build, "")

		source, _ := format.Source([]byte(code))

		tLine := fmt.Sprintf("%s %s %s", line, table, line)

		fmt.Printf("\n%s\n%s\n%s\n", tLine, string(source), tLine)

		if cmdParam.Output != "" {
			filename := filepath.Join(cmdParam.Output, fmt.Sprintf("%s.go", table))

			err := ioutil.WriteFile(filename, source, 0644)
			if err != nil {
				common.Log.Errorf("写入文件%s失败: %s", filename, err)
			} else {
				common.Log.Infof("写入文件%s", filename)
			}

		}
	}
}
