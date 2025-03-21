package compiler

import (
	model2 "github.com/macro-funs/tabkit/model"
	"github.com/macro-funs/tabkit/report"
	"strings"
)

func transposeKVtoData(symbols *model2.TypeTable, kvtab *model2.DataTable) (ret *model2.DataTable) {

	ret = model2.NewDataTable()
	ret.HeaderType = kvtab.HeaderType
	ret.OriginalHeaderType = kvtab.HeaderType
	ret.FileName = kvtab.FileName
	ret.SheetName = kvtab.SheetName

	// 添加表头
	ret.AddRow()

	// 添加数据行
	ret.AddRow()

	// 遍历KV表的每一行
	for row := 1; row < len(kvtab.Rows); row++ {

		fieldName := kvtab.GetValueByName(row, "字段名")
		fieldType := kvtab.GetValueByName(row, "字段类型")
		name := kvtab.GetValueByName(row, "标识名")

		arraySplitter := kvtab.GetValueByName(row, "数组切割")

		tags := kvtab.GetValueByName(row, "标记")

		var tf model2.TypeDefine
		tf.Kind = model2.TypeUsage_HeaderStruct
		tf.ObjectType = kvtab.HeaderType

		tf.Name = name.Value

		if !model2.PrimitiveExists(fieldType.Value) && !symbols.ObjectExists(fieldType.Value) { // 对象检查
			report.ReportError("UnknownFieldType", fieldType.Value, fieldType.String())
		}

		tf.FieldName = fieldName.Value
		tf.FieldType = fieldType.Value
		tf.ArraySplitter = arraySplitter.Value

		// 将KV表的Tags转换过去
		if tags != nil && tags.Value != "" {
			tagsType := kvtab.HeaderByName("标记")
			tf.Tags = strings.Split(tags.Value, tagsType.TypeInfo.ArraySplitter)
		}

		if symbols.FieldByName(tf.ObjectType, tf.FieldName) != nil {
			report.ReportError("DuplicateKVField", fieldName.String())
		}

		symbols.AddField(&tf, kvtab, row)

		// 输出表的表头原始数据
		headerCell := ret.AddCell(0)
		headerCell.Value = fieldName.Value

		header := ret.MustGetHeader(headerCell.Col)
		header.Cell.Value = fieldName.Value
		header.TypeInfo = &tf

		inputValueCell := kvtab.GetValueByName(row, "值")

		outputValueCell := ret.AddCell(1)
		outputValueCell.CopyFrom(inputValueCell)

	}

	return
}
