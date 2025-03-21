package checker

import (
	model2 "github.com/macro-funs/tabkit/model"
	"github.com/macro-funs/tabkit/report"
	"strconv"
)

// 检查数据与定义类型是否匹配
func checkDataType(globals *model2.Globals) {

	var currHeader *model2.HeaderField
	var crrCell *model2.Cell

	for _, tab := range globals.Datas.AllTables() {

		// 遍历输入数据的每一列
		for _, header := range tab.Headers {

			// 输入的列头，为空表示改列被注释
			if header.TypeInfo == nil {
				continue
			}

			for row := 1; row < len(tab.Rows); row++ {

				inputCell := tab.GetCell(row, header.Cell.Col)

				// 这行被注释，无效行
				if inputCell == nil {
					continue
				}

				crrCell = inputCell
				currHeader = header

				if header.TypeInfo.IsArray() {
					for _, value := range inputCell.ValueList {

						err := checkSingleValue(header, value)
						if err != nil {
							report.ReportError("DataMissMatchTypeDefine", currHeader.TypeInfo.FieldType, crrCell.String())
						}
					}
				} else if inputCell.Value != "" {
					err := checkSingleValue(header, inputCell.Value)
					if err != nil {
						report.ReportError("DataMissMatchTypeDefine", currHeader.TypeInfo.FieldType, crrCell.String())
					}
				}

			}
		}
	}
}

func checkSingleValue(header *model2.HeaderField, value string) error {
	switch model2.LanguagePrimitive(header.TypeInfo.FieldType, "go") {
	case "int16":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseInt(value, 10, 16)
		if err != nil {
			return err
		}
	case "int32":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
	case "int64":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
	case "uint16":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return err
		}
	case "uint32":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
	case "uint64":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
	case "float32":
		_, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
	case "float64":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
	case "bool":
		_, err := model2.ParseBool(value)
		if err != nil {
			return err
		}
	}

	return nil
}
