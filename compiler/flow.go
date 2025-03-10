package compiler

import (
	checker2 "github.com/macro-funs/tabkit/checker"
	"github.com/macro-funs/tabkit/helper"
	model2 "github.com/macro-funs/tabkit/model"
	report2 "github.com/macro-funs/tabkit/report"
)

func Compile(globals *model2.Globals) (ret error) {

	defer func() {

		switch err := recover().(type) {
		case *report2.TableError:
			ret = err
		case nil:
		default:
			panic(err)
		}

	}()

	model2.InitBuiltinTypes(globals.Types)

	report2.Log.Debugf("Loading Index file: '%s'... ", globals.IndexFile)
	err := LoadIndexTable(globals, globals.IndexFile)

	if err != nil {
		return err
	}

	// 测试时, 这个Getter会被提前设置为MemFile, 普通导出时, 这个Getter为空
	if globals.TableGetter == nil {
		tabLoader := helper.NewFileLoader(!globals.ParaLoading, globals.CacheDir)

		if globals.ParaLoading {
			for _, pragma := range globals.IndexList {
				tabLoader.AddFile(pragma.TableFileName)
			}

			tabLoader.Commit()
		}

		globals.TableGetter = tabLoader
	}

	var kvList, dataList model2.DataTableList

	// 加载多种表
	err = loadVariantTables(globals, &kvList, &dataList)

	if err != nil {
		return err
	}

	report2.Log.Debugln("Checking types...")
	checker2.CheckType(globals.Types)
	checker2.PreCheck(&dataList)

	if kvList.Count() > 0 {
		report2.Log.Debugln("Merge key-value tables...")

		// 合并所有的KV表行
		var mergedKV model2.DataTableList
		MergeData(&kvList, &mergedKV, globals.Types)

		// 完整KV表转置为普通数据表
		for _, tab := range mergedKV.AllTables() {

			dataList.AddDataTable(transposeKVtoData(globals.Types, tab))
		}
	}

	// KV转置后, 再检查一次
	checker2.CheckType(globals.Types)

	report2.Log.Debugln("Merge data tables...")

	// 合并所有的数据表
	MergeData(&dataList, &globals.Datas, globals.Types)

	checker2.PostCheck(globals)

	return nil
}
