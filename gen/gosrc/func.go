package gosrc

import (
	"fmt"
	model2 "github.com/macro-funs/tabkit/model"
	"strings"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

// 将定义用的类型，转换为不同语言对应的复合类型

func init() {
	UsefulFunc["GoType"] = func(tf *model2.TypeDefine) string {

		convertedType := model2.LanguagePrimitive(tf.FieldType, "go")

		if tf.IsArray() {
			return "[]" + convertedType
		}

		return convertedType
	}

	UsefulFunc["GoTabTag"] = func(fieldType *model2.TypeDefine) string {

		var sb strings.Builder

		var kv []string

		if fieldType.Name != "" {
			kv = append(kv, fmt.Sprintf("tb_name:\"%s\"", fieldType.Name))
		}

		if len(kv) > 0 {
			sb.WriteString("`")

			for _, s := range kv {
				sb.WriteString(s)
			}

			sb.WriteString("`")
		}

		return sb.String()
	}

	UsefulFunc["JsonTabOmit"] = func() string {
		return "`json:\"-\"`"
	}

}
