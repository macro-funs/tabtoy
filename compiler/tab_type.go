package compiler

import (
	"github.com/macro-funs/tabkit/helper"
	model2 "github.com/macro-funs/tabkit/model"
	"github.com/macro-funs/tabkit/report"
)

func LoadTypeTable(typeTab *model2.TypeTable, indexGetter helper.FileGetter, fileName string) error {

	tabs, err := LoadDataTable(indexGetter, fileName, "TypeDefine", "TypeDefine", typeTab)

	if err != nil {
		return err
	}

	for _, tab := range tabs {

		for row := 1; row < len(tab.Rows); row++ {

			var objtype model2.TypeDefine

			if !ParseRow(&objtype, tab, row, typeTab) {
				continue
			}

			if objtype.Kind == model2.TypeUsage_None {
				report.ReportError("UnknownTypeKind", objtype.ObjectType, objtype.FieldName)
			}

			if typeTab.FieldByName(objtype.ObjectType, objtype.FieldName) != nil {
				cell := tab.GetValueByName(row, "字段名")
				if cell != nil {
					report.ReportError("DuplicateTypeFieldName", cell.String(), objtype.ObjectType, objtype.FieldName)
				} else {
					report.ReportError("InvalidTypeTable", objtype.ObjectType, objtype.FieldName, tab.FileName)
				}

			}

			typeTab.AddField(&objtype, tab, row)
		}

	}

	return nil
}
