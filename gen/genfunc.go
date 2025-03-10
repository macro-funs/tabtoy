package gen

import (
	"github.com/macro-funs/tabkit/model"
)

type GenSingleFile func(globals *model.Globals) (data []byte, err error)

type GenCustom func(globals *model.Globals, param string) (err error)
