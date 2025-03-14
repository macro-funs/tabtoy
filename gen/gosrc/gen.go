package gosrc

import (
	"github.com/davyxu/protoplus/codegen"
	"github.com/macro-funs/tabkit/gen"
	"github.com/macro-funs/tabkit/model"
	"github.com/macro-funs/tabkit/report"
)

func Generate(globals *model.Globals) (data []byte, err error) {

	cg := codegen.NewCodeGen("gosrc").
		RegisterTemplateFunc(codegen.UsefulFunc).
		RegisterTemplateFunc(gen.UsefulFunc).
		RegisterTemplateFunc(UsefulFunc)

	err = cg.ParseTemplate(templateText, globals).Error()
	if err != nil {
		return
	}

	err = cg.FormatGoCode().Error()
	if err != nil {
		report.Log.Infoln(string(cg.Code()))
		return
	}

	err = cg.WriteBytes(&data).Error()

	return
}
