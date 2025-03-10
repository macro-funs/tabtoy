package gosrc

// 报错行号+3
const templateText = `// Generated by github.com/macro-funs/tabkit
// DO NOT EDIT!!
// Version: {{.Version}}
package {{.PackageName}}

import "errors"

type {{.CombineStructName}}EnumValue struct {
	Name  string
	Index int32
}

{{range $sn, $objName := $.Types.EnumNames}}
type {{$objName}} int32
const (	{{range $fi,$field := $.Types.AllFieldByName $objName}}
	{{$objName}}_{{$field.FieldName}} = {{$field.Value}} // {{$field.Name}} {{end}}
)

var (

	{{$objName}}EnumValues = []{{$.CombineStructName}}EnumValue{ {{range $fi,$field := $.Types.AllFieldByName $objName}}
		{ Name: "{{$field.FieldName}}", Index:{{$field.Value}} }, // {{$field.Name}} {{end}}
	}
	{{$objName}}MapperValueByName = map[string]int32{}
	{{$objName}}MapperNameByValue = map[int32]string{}
)

func (self {{$objName}}) String() string {
	name, _ := {{$objName}}MapperNameByValue[int32(self)]
	return name
}
{{end}}

{{range $sn, $objName := $.Types.StructNames}}
type {{$objName}} struct{ {{range $fi,$field := $.Types.AllFieldByName $objName}}
	{{$field.FieldName}} {{GoType $field}} {{GoTabTag $field}} {{end}}
}
{{end}}

// Combine struct
type {{.CombineStructName}} struct { {{range $ti, $tab := $.Datas.AllTables}}
	{{$tab.HeaderType}} []*{{$tab.HeaderType}} // table: {{$tab.HeaderType}} {{end}}

	// Indices {{range $ii, $idx := GetIndices $}}
	{{$idx.Table.HeaderType}}By{{$idx.FieldInfo.FieldName}} map[{{GoType $idx.FieldInfo}}]*{{$idx.Table.HeaderType}}	{{JsonTabOmit}} // table: {{$idx.Table.HeaderType}} {{end}}

	// Handlers
	postHandlers []func(*{{.CombineStructName}}) error {{JsonTabOmit}}
	preHandlers  []func(*{{.CombineStructName}}) error {{JsonTabOmit}}
	
	indexHandler map[string]func() {{JsonTabOmit}}
	resetHandler map[string]func() {{JsonTabOmit}}
}

{{if HasKeyValueTypes $}}
{{range $ti, $name := GetKeyValueTypeNames $}} // table: {{$name}}
func (self*{{$.CombineStructName}}) GetKeyValue_{{$name}}() *{{$name}}{
	return self.{{$name}}[0]
}
{{end}}{{end}}

// 注册加载后回调(用于构建数据)
func (self *{{.CombineStructName}}) RegisterPostEntry(h func(*{{.CombineStructName}}) error) {

	if h == nil {
		panic("empty postload handler")
	}

	self.postHandlers = append(self.postHandlers, h)
}

// 注册加载前回调(用于清除数据)
func (self *{{.CombineStructName}}) RegisterPreEntry(h func(*{{.CombineStructName}}) error) {

	if h == nil {
		panic("empty preload handler")
	}

	self.preHandlers = append(self.preHandlers, h)
}

// 清除索引和数据
func (self *{{.CombineStructName}}) ResetData() error {

	err := self.InvokePreHandler()
	if err != nil {
		return err
	}

	return self.ResetTable("")
}

// 全局表构建索引及通知回调
func (self *{{.CombineStructName}}) BuildData() error {

	err := self.IndexTable("")
	if err != nil {
		return err
	}

	return self.InvokePostHandler()
}

// 调用加载前回调
func (self *{{.CombineStructName}}) InvokePreHandler() error {
	for _, h := range self.preHandlers {
		if err := h(self); err != nil {
			return err
		}
	}

	return nil
}

// 调用加载后回调
func (self *{{.CombineStructName}}) InvokePostHandler() error {
	for _, h := range self.postHandlers {
		if err := h(self); err != nil {
			return err
		}
	}

	return nil
}


// 为表建立索引. 表名为空时, 构建所有表索引
func (self *{{.CombineStructName}}) IndexTable(tableName string) error {

	if tableName == "" {

		for _, h := range self.indexHandler {
			h()
		}
		return nil

	} else {
		if h, ok := self.indexHandler[tableName]; ok {
			h()
		}

		return nil
	}
}

// 重置表格数据
func (self *{{.CombineStructName}}) ResetTable(tableName string) error {
	if tableName == "" {
		for _, h := range self.resetHandler {
			h()
		}

		return nil
	} else {
		if h, ok := self.resetHandler[tableName]; ok {
			h()
			return nil
		}

		return errors.New("reset table failed, table not found: " + tableName)
	}
}


// 初始化表实例
func New{{.CombineStructName}}() *{{.CombineStructName}}{

	self := &{{.CombineStructName}}{
		indexHandler: make(map[string]func()),
		resetHandler: make(map[string]func()),
	}

	{{range $ti, $tab := $.Datas.AllTables}}
	self.indexHandler["{{$tab.HeaderType}}"] = func() {
		{{range $ii, $idx := GetIndicesByTable $tab}}
		for _, v := range self.{{$idx.Table.HeaderType}} {
			self.{{$idx.Table.HeaderType}}By{{$idx.FieldInfo.FieldName}}[v.{{$idx.FieldInfo.FieldName}}] = v
		}{{end}}
	}
	{{end}}


	{{range $ti, $tab := $.Datas.AllTables}}
	self.resetHandler["{{$tab.HeaderType}}"] = func() {
		self.{{$tab.HeaderType}} = nil
		{{range $ii, $idx := GetIndicesByTable $tab}}
		self.{{$idx.Table.HeaderType}}By{{$idx.FieldInfo.FieldName}} = map[{{GoType $idx.FieldInfo}}]*{{$idx.Table.HeaderType}}{} {{end}}
	} {{end}}

	self.ResetData()

	return self
}

func init(){
	{{range $sn, $objName := $.Types.EnumNames}}
	for _, v := range {{$objName}}EnumValues {
		{{$objName}}MapperValueByName[v.Name] = v.Index
		{{$objName}}MapperNameByValue[v.Index] = v.Name
	}
	{{end}}
}

`
