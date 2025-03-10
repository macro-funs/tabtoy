package checker

import (
	model2 "github.com/macro-funs/tabkit/model"
	"github.com/macro-funs/tabkit/report"
)

func checkRepeat(globals *model2.Globals) {

	for _, tab := range globals.Datas.AllTables() {

		// 遍历输入数据的每一列
		for _, header := range tab.Headers {

			// 输入的列头，为空表示改列被注释
			if header.TypeInfo == nil {
				continue
			}

			// 这列需要建立索引
			if header.TypeInfo.MakeIndex {

				checker := map[string]*model2.Cell{}

				for row := 1; row < len(tab.Rows); row++ {

					inputCell := tab.GetCell(row, header.Cell.Col)

					// 这行被注释，无效行
					if inputCell == nil {
						continue
					}

					if inputCell.Value == "" {
						continue
					}

					if _, ok := checker[inputCell.Value]; ok {

						report.ReportError("DuplicateValueInMakingIndex", inputCell.String())

					} else {
						checker[inputCell.Value] = inputCell
					}

				}
			}
		}
	}
}
