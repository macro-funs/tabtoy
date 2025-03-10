package checker

import (
	"github.com/ahmetb/go-linq"
	model2 "github.com/macro-funs/tabkit/model"
	"github.com/macro-funs/tabkit/report"
	"go/token"
)

func CheckType(typeTab *model2.TypeTable) {

	typeTable_CheckField(typeTab)

	typeTable_CheckEnumValueEmpty(typeTab)

	typeTable_CheckDuplicateEnumValue(typeTab)

}

func isValidFieldName(name string) bool {

	return token.IsIdentifier(name)
}

func typeTable_CheckField(typeTab *model2.TypeTable) {
	for _, td := range typeTab.Raw() {

		if !isValidFieldName(td.Define.FieldName) {
			cell := td.Tab.GetValueByName(td.Row, "字段名")
			report.ReportError("InvalidFieldName", cell.String())
		}
	}
}

func typeTable_CheckEnumValueEmpty(typeTab *model2.TypeTable) {
	linq.From(typeTab.Raw()).Where(func(raw interface{}) bool {
		td := raw.(*model2.TypeData)

		return td.Define.Kind == model2.TypeUsage_Enum && td.Define.Value == ""
	}).ForEach(func(raw interface{}) {
		td := raw.(*model2.TypeData)

		cell := td.Tab.GetValueByName(td.Row, "值")

		report.ReportError("EnumValueEmpty", cell.String())
	})
}

func typeTable_CheckDuplicateEnumValue(typeTab *model2.TypeTable) {

	type NameValuePair struct {
		Name  string
		Value string
	}

	checker := map[NameValuePair]*model2.TypeData{}

	for _, td := range typeTab.Raw() {

		if td.Define.IsBuiltin || td.Define.Kind != model2.TypeUsage_Enum {
			continue
		}

		key := NameValuePair{td.Define.ObjectType, td.Define.Value}

		if _, ok := checker[key]; ok {

			cell := td.Tab.GetValueByName(td.Row, "值")

			report.ReportError("DuplicateEnumValue", cell.String())
		}

		checker[key] = td
	}
}
