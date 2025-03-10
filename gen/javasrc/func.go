package javasrc

import (
	"fmt"
	model2 "github.com/macro-funs/tabkit/model"
	"github.com/macro-funs/tabkit/util"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

// 将定义用的类型，转换为不同语言对应的复合类型

func wrapSingleValue(globals *model2.Globals, valueType *model2.TypeDefine, value string) string {
	switch {
	case valueType.FieldType == "string": // 字符串
		return util.StringWrap(util.StringEscape(value))
	case valueType.FieldType == "float32":
		return value
	case globals.Types.IsEnumKind(valueType.FieldType): // 枚举
		t := globals.Types.ResolveEnum(valueType.FieldType, value)
		if t != nil {
			return t.Define.ObjectType + "." + t.Define.FieldName
		}

		return ""
	case valueType.FieldType == "bool":

		v, _ := model2.ParseBool(value)
		if v {
			return "true"
		}

		return "false"
	}

	if value == "" {
		return model2.FetchDefaultValue(valueType.FieldType)
	}

	return value
}

func init() {
	UsefulFunc["JavaType"] = func(tf *model2.TypeDefine, requireRef bool) string {

		convertedType := model2.LanguagePrimitive(tf.FieldType, "java")

		if requireRef {
			// https://www.geeksforgeeks.org/difference-between-an-integer-and-int-in-java/
			switch convertedType {
			case "int":
				convertedType = "Integer"
			case "short":
				convertedType = "Short"
			case "long":
				convertedType = "Integer"
			case "float":
				convertedType = "Float"
			case "double":
				convertedType = "Double"
			case "boolean":
				convertedType = "Boolean"
			}
		}

		if tf.IsArray() {
			return convertedType + "[]"
		}

		return convertedType
	}

	UsefulFunc["JavaDefaultValue"] = func(globals *model2.Globals, tf *model2.TypeDefine) string {

		convertedType := model2.LanguagePrimitive(tf.FieldType, "java")

		if tf.IsArray() {
			return fmt.Sprintf("new %s[]{}", convertedType)
		} else {
			return wrapSingleValue(globals, tf, "")
		}

		return convertedType
	}

}
