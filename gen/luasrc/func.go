package luasrc

import (
	"github.com/macro-funs/tabkit/gen"
	model2 "github.com/macro-funs/tabkit/model"
	"strings"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

func WrapValue(globals *model2.Globals, cell *model2.Cell, valueType *model2.TypeDefine) string {
	if valueType.IsArray() {

		var sb strings.Builder
		sb.WriteString("{")

		if cell != nil {
			for index, elementValue := range cell.ValueList {
				if index > 0 {
					sb.WriteString(",")
				}
				sb.WriteString(gen.WrapSingleValue(globals, valueType, elementValue))
			}
		}

		sb.WriteString("}")

		return sb.String()

	} else {

		var value string
		if cell != nil {
			value = cell.Value
		}

		return gen.WrapSingleValue(globals, valueType, value)
	}
}

func init() {
	UsefulFunc["WrapTabValue"] = func(globals *model2.Globals, dataTable *model2.DataTable, allHeaders []*model2.TypeDefine, row, col int) (ret string) {
		// 找到完整的表头（按完整表头遍历）
		header := allHeaders[col]

		if header == nil {
			return ""
		}

		// 在单元格找到值
		valueCell := dataTable.GetCell(row, col)

		if valueCell != nil {

			return WrapValue(globals, valueCell, header)
		} else {
			// 这个表中没有这列数据
			return WrapValue(globals, nil, header)
		}
	}

	UsefulFunc["IsWrapFieldName"] = func(globals *model2.Globals, dataTable *model2.DataTable, allHeaders []*model2.TypeDefine, row, col int) (ret bool) {
		// 找到完整的表头（按完整表头遍历）
		header := allHeaders[col]

		if header == nil {
			return false
		}

		if globals.CanDoAction(model2.ActionNoGennFieldLua, header) {
			return false
		}

		return true
	}

}
