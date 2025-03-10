package cssrc

import (
	"fmt"
	"github.com/davyxu/tabtoy/gen/bindata"
	model2 "github.com/davyxu/tabtoy/model"
	"github.com/davyxu/tabtoy/util"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

func wrapSingleValue(globals *model2.Globals, valueType *model2.TypeDefine, value string) string {
	switch {
	case valueType.FieldType == "string": // 字符串

		// C#特殊优化
		if value == "" {
			return "string.Empty"
		}

		return util.StringWrap(util.StringEscape(value))
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
	UsefulFunc["CSType"] = func(tf *model2.TypeDefine) string {

		convertedType := model2.LanguagePrimitive(tf.FieldType, "cs")

		if tf.IsArray() {
			return fmt.Sprintf("List<%s>", convertedType)
		}

		return convertedType
	}

	UsefulFunc["CSTag"] = func(globals *model2.Globals, fieldIndex int, tf *model2.TypeDefine) string {

		tag := bindata.MakeTag(globals, tf, fieldIndex)

		return fmt.Sprintf("0x%x", tag)
	}

	UsefulFunc["CSStructTag"] = func() string {

		return fmt.Sprintf("0x%x", bindata.MakeTagStructArray())
	}

	UsefulFunc["CSReader"] = func(globals *model2.Globals, tf *model2.TypeDefine) (ret string) {

		convertedType := model2.LanguagePrimitive(tf.FieldType, "cs")

		switch {
		case convertedType == "float":
			ret = "Float"
		case convertedType == "double":
			ret = "Double"
		case convertedType == "string":
			ret = "String"
		case convertedType == "bool":
			ret = "Bool"
		case globals.Types.IsEnumKind(tf.FieldType):
			ret = "Enum"
		default:
			ret = convertedType
		}

		return
	}

	UsefulFunc["CSDefaultValue"] = func(globals *model2.Globals, tf *model2.TypeDefine) string {

		convertedType := model2.LanguagePrimitive(tf.FieldType, "cs")

		if tf.IsArray() {
			return fmt.Sprintf("new List<%s>()", convertedType)
		} else {
			return wrapSingleValue(globals, tf, "")
		}

	}

	UsefulFunc["IsWarpFieldName"] = func(globals *model2.Globals, tf *model2.TypeDefine) bool {

		if globals.CanDoAction(model2.ActionNoGennFieldCsharp, tf) {
			return false
		}
		return true
	}

}
