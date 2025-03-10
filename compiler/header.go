package compiler

import (
	"github.com/davyxu/tabtoy/helper"
	model2 "github.com/davyxu/tabtoy/model"
	"github.com/davyxu/tabtoy/report"
	"strings"
)

func LoadHeader(sheet helper.TableSheet, tab *model2.DataTable, resolveTableType string, typeTab *model2.TypeTable) (maxCol int) {
	for col := 0; ; col++ {
		headerValue := sheet.GetValue(0, col, nil)
		if headerValue == "" {
			break
		}
		maxCol = col
		if strings.HasPrefix(headerValue, "#") {
			continue
		}
		header := tab.MustGetHeader(col)
		header.Cell.CopyFrom(&model2.Cell{
			Value: headerValue,
			Col:   col,
			Row:   0,
			Table: tab,
		})

	}
	resolveHeaderFields(tab, resolveTableType, typeTab)
	checkHeaderTypes(tab, typeTab)
	return
}

func checkHeaderTypes(tab *model2.DataTable, symbols *model2.TypeTable) {
	for _, header := range tab.Headers {
		if header.TypeInfo == nil {
			continue
		}
		// 原始类型检查
		if !model2.PrimitiveExists(header.TypeInfo.FieldType) &&
			!symbols.ObjectExists(header.TypeInfo.FieldType) { // 对象检查
			report.ReportError("UnknownFieldType", header.TypeInfo.FieldType, header.Cell.String())
		}
	}
}

func headerValueExists(offset int, name string, headers []*model2.HeaderField) bool {
	for i := offset; i < len(headers); i++ {
		if headers[i].Cell.Value == name {
			return true
		}
	}
	return false
}

func resolveHeaderFields(tab *model2.DataTable, tableObjectType string, typeTab *model2.TypeTable) {
	tab.OriginalHeaderType = tableObjectType
	for index, header := range tab.Headers {
		if header.Cell.Value == "" {
			continue
		}
		tf := typeTab.FieldByName(tableObjectType, header.Cell.Value)
		if tf == nil {
			report.ReportError("HeaderFieldNotDefined", header.Cell.String(), tableObjectType)
		}
		if headerValueExists(index+1, header.Cell.Value, tab.Headers) && !tf.IsArray() {
			report.ReportError("DuplicateHeaderField", header.Cell.String())
		}
		// 解析好的类型
		header.TypeInfo = tf
	}
}
