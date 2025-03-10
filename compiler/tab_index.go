package compiler

import (
	model2 "github.com/macro-funs/tabkit/model"
	"path/filepath"
	"sort"
	"strings"
)

func parseIndexRow(tab *model2.DataTable, symbols *model2.TypeTable) (pragmaList []*model2.IndexDefine) {

	for row := 1; row < len(tab.Rows); row++ {

		var pragma model2.IndexDefine
		if !ParseRow(&pragma, tab, row, symbols) {
			continue
		}

		if pragma.Kind == model2.TableKind_Type {
			pragma.TableType = "TypeDefine"
		}

		if pragma.TableType == "" {

			_, name := filepath.Split(pragma.TableFileName)

			pragma.TableType = strings.TrimSuffix(name, filepath.Ext(pragma.TableFileName))
		}

		pragmaList = append(pragmaList, &pragma)
	}

	return
}

func LoadIndexTable(globals *model2.Globals, fileName string) error {

	if fileName == "" {
		return nil
	}

	// 加载原始数据
	tabs, err := LoadDataTable(globals.IndexGetter, fileName, "IndexDefine", "IndexDefine", globals.Types)

	if err != nil {
		return err
	}

	var pragmaList []*model2.IndexDefine

	for _, tab := range tabs {
		pragmaList = append(pragmaList, parseIndexRow(tab, globals.Types)...)
	}

	// 按表类型排序，保证类型表先读取
	sort.Slice(pragmaList, func(i, j int) bool {
		a := pragmaList[i]
		b := pragmaList[j]

		if a.Kind != b.Kind {
			return a.Kind < b.Kind
		}

		if a.TableType != b.TableType {
			return a.TableType < b.TableType
		}

		return a.TableFileName < b.TableFileName

	})

	globals.IndexList = pragmaList

	return nil
}
