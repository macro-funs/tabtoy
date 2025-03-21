package checker

import (
	model2 "github.com/macro-funs/tabkit/model"
	"github.com/macro-funs/tabkit/report"
)

func PreCheck(dataList *model2.DataTableList) {

	type ArrayFieldDefine struct {
		FieldName  string
		ObjectName string
	}

	fieldCountByField := map[ArrayFieldDefine]int{}

	// 合并前的数据表
	for _, tab := range dataList.AllTables() {

		// 遍历输入数据的每一列
		for _, header := range tab.Headers {

			// 输入的列头，为空表示改列被注释
			if header.TypeInfo == nil {
				continue
			}

			fieldKey := ArrayFieldDefine{
				FieldName:  header.TypeInfo.FieldName,
				ObjectName: header.TypeInfo.Name,
			}

			if header.TypeInfo.IsArray() {
				arrayFieldCount := tab.ArrayFieldCount(header)
				if preFieldCount, ok := fieldCountByField[fieldKey]; ok {

					if preFieldCount != arrayFieldCount {
						report.ReportError("ArrayMultiColumnDefineNotMatch")
					}
				} else {
					fieldCountByField[fieldKey] = arrayFieldCount
				}
			}

		}
	}
}

// merge后检查
func PostCheck(globals *model2.Globals) {

	checkEnumValue(globals)
	checkRepeat(globals)
	checkDataType(globals)
}
