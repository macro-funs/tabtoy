package checker

import (
	model2 "github.com/macro-funs/tabkit/model"
	"github.com/macro-funs/tabkit/report"
)

// 枚举值的解析是放在输出端处理的, 例如json中, 所以在这里提前检查
func checkEnumValue(globals *model2.Globals) {

	for _, tab := range globals.Datas.AllTables() {

		// 遍历输入数据的每一列
		for _, header := range tab.Headers {

			// 输入的列头，为空表示改列被注释
			if header.TypeInfo == nil {
				continue
			}

			if !globals.Types.IsEnumKind(header.TypeInfo.FieldType) {
				continue
			}

			for row := 1; row < len(tab.Rows); row++ {

				inputCell := tab.GetCell(row, header.Cell.Col)

				// 这行被注释，无效行
				if inputCell == nil {
					continue
				}

				if header.TypeInfo.IsArray() {

					for _, v := range inputCell.ValueList {
						checkEnumFieldValue(globals, header, v, inputCell)
					}

				} else {
					checkEnumFieldValue(globals, header, inputCell.Value, inputCell)
				}

			}
		}
	}
}

// 检查枚举值是否存在有效
func checkEnumFieldValue(globals *model2.Globals, header *model2.HeaderField, value string, inputCell *model2.Cell) {

	if inputCell.Value == "" {
		return
	}

	enumValue := globals.Types.GetEnumValue(header.TypeInfo.FieldType, value)
	if enumValue == nil {
		report.ReportError("UnknownEnumValue", header.TypeInfo.FieldType, inputCell.String())
	}

}
