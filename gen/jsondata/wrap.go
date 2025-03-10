package jsondata

import (
	model2 "github.com/davyxu/tabtoy/model"
	"strconv"
)

func wrapValue(globals *model2.Globals, valueCell *model2.Cell, valueType *model2.TypeDefine) interface{} {
	if valueType.IsArray() {

		var vlist = make([]interface{}, 0)
		// 空的单元格，导出空数组，除非强制指定填充默认值
		if valueCell != nil {

			for _, elementValue := range valueCell.ValueList {

				vlist = append(vlist, wrapSingleValue(globals, valueType, elementValue))
			}
		}

		return vlist

	} else {

		var value string
		if valueCell != nil {
			value = valueCell.Value
		}

		return wrapSingleValue(globals, valueType, value)
	}
}

func wrapSingleValue(globals *model2.Globals, valueType *model2.TypeDefine, value string) interface{} {

	goType := model2.LanguagePrimitive(valueType.FieldType, "go")

	switch {
	case goType == "string": // 字符串
		return value // json自己会做转义, 所以这里无需转义
	case goType == "float32":

		if value == "" {
			return float32(0)
		}

		f64, _ := strconv.ParseFloat(value, 32)
		return float32(f64)
	case goType == "float64":

		if value == "" {
			return float64(0)
		}

		f64, _ := strconv.ParseFloat(value, 64)
		return f64
	case globals.Types.IsEnumKind(valueType.FieldType): // 枚举
		enumValue := globals.Types.ResolveEnumValue(valueType.FieldType, value)
		i, _ := strconv.Atoi(enumValue)
		return int32(i)
	case goType == "bool":

		v, _ := model2.ParseBool(value)
		if v {
			return true
		}

		return false
	case goType == "int16":
		i64, _ := strconv.ParseInt(value, 10, 16)
		return int16(i64)
	case goType == "int32":
		i64, _ := strconv.ParseInt(value, 10, 32)
		return int32(i64)
	case goType == "int64":
		i64, _ := strconv.ParseInt(value, 10, 64)
		return i64
	case goType == "uint16":
		i64, _ := strconv.ParseInt(value, 10, 16)
		return uint16(i64)
	case goType == "uint32":
		i64, _ := strconv.ParseUint(value, 10, 32)
		return uint32(i64)
	case goType == "uint64":
		i64, _ := strconv.ParseUint(value, 10, 64)
		return i64
	}

	if value == "" {
		return model2.FetchDefaultValue(valueType.FieldType)
	}

	return value
}
