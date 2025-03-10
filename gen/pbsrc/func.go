package pbsrc

import (
	"fmt"
	model2 "github.com/davyxu/tabtoy/model"
	"strings"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

// 将定义用的类型，转换为不同语言对应的复合类型

func init() {
	UsefulFunc["PbType"] = func(tf *model2.TypeDefine) string {

		pbType := model2.LanguagePrimitive(tf.FieldType, "pb")

		if tf.IsArray() {
			return "repeated " + pbType
		}

		return pbType
	}

	UsefulFunc["PbTag"] = func(fieldIndex int, fieldType *model2.TypeDefine) string {

		var sb strings.Builder
		fmt.Fprintf(&sb, "= %d", fieldIndex+1)
		return sb.String()
	}

	UsefulFunc["PbCombineField"] = func(fieldIndex int) string {

		var sb strings.Builder
		fmt.Fprintf(&sb, "= %d", fieldIndex+1)
		return sb.String()
	}
}
