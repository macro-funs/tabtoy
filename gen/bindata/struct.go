package bindata

import (
	model2 "github.com/macro-funs/tabkit/model"
)

// 写入表的一行
func writeStruct(globals *model2.Globals, tab *model2.DataTable, row int) (*BinaryWriter, error) {

	structWriter := NewBinaryWriter()

	// 一个结构体
	for _, header := range tab.Headers {

		if header == nil {
			continue
		}

		if globals.CanDoAction(model2.ActionNoGenFieldBinary, header) {
			continue
		}

		cell := tab.GetCell(row, header.Cell.Col)

		if cell == nil {
			continue
		}

		goType := model2.LanguagePrimitive(header.TypeInfo.FieldType, "go")

		// 写入字段
		if header.TypeInfo.IsArray() {

			for _, elementValue := range cell.ValueList {

				if err := writePair(globals, structWriter, header.TypeInfo, goType, elementValue, header.Cell.Col); err != nil {
					return nil, err
				}
			}

		} else {

			// 空格不输出
			if cell.Value != "" {

				if err := writePair(globals, structWriter, header.TypeInfo, goType, cell.Value, header.Cell.Col); err != nil {
					return nil, err
				}
			}

		}

	}

	return structWriter, nil
}
