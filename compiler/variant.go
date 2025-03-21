package compiler

import (
	model2 "github.com/macro-funs/tabkit/model"
	"github.com/macro-funs/tabkit/report"
)

func loadVariantTables(globals *model2.Globals, kvList, dataList *model2.DataTableList) error {
	report.Log.Debugln("Loading tables...")

	// 遍历索引里的每一行配置
	for _, pragma := range globals.IndexList {

		if globals.CanDoAction(model2.ActionNoGenTable, pragma) {
			report.Log.Debugf("   (%s) %s   action=nogentable, ignored(tag: %v)", pragma.TableType, pragma.TableFileName, pragma.Tags)
			continue
		}

		report.Log.Debugf("   (%s) %s", pragma.TableType, pragma.TableFileName)

		switch pragma.Kind {
		case model2.TableKind_Data:

			tablist, err := LoadDataTable(globals.TableGetter, pragma.TableFileName, pragma.TableType, pragma.TableType, globals.Types)

			if err != nil {
				return err
			}

			for _, tab := range tablist {

				dataList.AddDataTable(tab)
			}

		case model2.TableKind_Type:

			err := LoadTypeTable(globals.Types, globals.TableGetter, pragma.TableFileName)

			if err != nil {
				return err
			}

		case model2.TableKind_KeyValue:
			tablist, err := LoadDataTable(globals.TableGetter, pragma.TableFileName, pragma.TableType, "KVDefine", globals.Types)

			if err != nil {
				return err
			}

			for _, tab := range tablist {

				kvList.AddDataTable(tab)
			}

		}
	}

	return nil
}
