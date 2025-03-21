package pbsrc

import (
	"github.com/davyxu/protoplus/codegen"
	"github.com/macro-funs/tabkit/gen"
	"github.com/macro-funs/tabkit/model"
)

func Generate(globals *model.Globals) (data []byte, err error) {

	cg := codegen.NewCodeGen("pbsrc").
		RegisterTemplateFunc(codegen.UsefulFunc).
		RegisterTemplateFunc(gen.UsefulFunc).
		RegisterTemplateFunc(UsefulFunc)

	err = cg.ParseTemplate(templateText, globals).Error()
	if err != nil {
		return
	}

	err = cg.WriteBytes(&data).Error()

	return
}
