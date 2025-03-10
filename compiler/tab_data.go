package compiler

import (
	helper2 "github.com/macro-funs/tabkit/helper"
	model2 "github.com/macro-funs/tabkit/model"
	"github.com/pkg/errors"
	"strings"
)

func readOneRow(sheet helper2.TableSheet, tab *model2.DataTable, row int) bool {

	for _, header := range tab.Headers {

		if header.TypeInfo == nil {
			continue
		}

		// 浮点数用库取时，需要特殊处理
		isFloat := model2.LanguagePrimitive(header.TypeInfo.FieldType, "go") == "float32"

		value := sheet.GetValue(row, header.Cell.Col, &helper2.ValueOption{ValueAsFloat: isFloat})

		// 首列带#时，本行忽略
		if header.Cell.Col == 0 && strings.HasPrefix(value, "#") {
			return false
		}

		cell := tab.MustGetCell(row, header.Cell.Col)
		cell.Value = value
	}

	return true
}

func LoadDataTable(filegetter helper2.FileGetter, fileName, headerType, resolveHeaderType string, typeTab *model2.TypeTable) (ret []*model2.DataTable, err error) {
	file, err := filegetter.GetFile(fileName)
	if err != nil {
		return nil, errors.Wrap(err, fileName)
	}

	for _, sheet := range file.Sheets() {

		tab := model2.NewDataTable()
		tab.HeaderType = headerType
		tab.FileName = fileName
		tab.SheetName = sheet.Name()

		ret = append(ret, tab)

		maxCol := LoadHeader(sheet, tab, resolveHeaderType, typeTab)

		// 遍历所有数据行
		for row := 0; ; row++ {

			if sheet.IsRowEmpty(row, maxCol+1) {
				break
			}

			// 读取每一行
			readOneRow(sheet, tab, row)
		}

	}

	return
}
