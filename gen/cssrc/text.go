package cssrc

// 报错行号+3
const templateText = `// Generated by github.com/macro-funs/tabkit
// DO NOT EDIT!!
// Version: {{.Version}}
using System;
using System.Collections.Generic;

namespace {{.PackageName}}
{ 	{{range $sn, $objName := $.Types.EnumNames}}
	public enum {{$objName}}
	{ {{range $fi,$field := $.Types.AllFieldByName $objName}}
		{{$field.FieldName}} = {{$field.Value}}, // {{$field.Name}} {{end}}
	}
	{{end}}	{{range $sn, $objName := $.Types.StructNames}}
	public partial class {{$objName}} : tabtoy.ITableSerializable
	{ 
		{{range $fi,$field := $.Types.AllFieldByName $objName}}{{if IsWarpFieldName $ $field}}public {{CSType $field}} {{$field.FieldName}} = {{CSDefaultValue $ $field}};
		{{end}}{{end}}
		{{if $.GenBinary}}#region Deserialize Code
		public void Deserialize( tabtoy.TableReader reader )
		{
			UInt32 mamaSaidTagNameShouldBeLong = 0;
            while ( reader.ReadTag(ref mamaSaidTagNameShouldBeLong) )
            {
 				switch (mamaSaidTagNameShouldBeLong)
				{ 
					{{range $fi,$field := $.Types.AllFieldByName $objName}}{{if IsWarpFieldName $ $field}}case {{CSTag $ $fi $field}}:
                	{
						reader.Read{{CSReader $ $field}}( ref {{$field.FieldName}} );
                	}
					break;
					{{end}}{{end}}
                    default:
                    {
                        reader.SkipFiled(mamaSaidTagNameShouldBeLong);                            
                    }
                    break;
				}
			}
		}
		#endregion {{end}}
	}
	{{end}}

	// Combine struct
	public partial class {{.CombineStructName}}
	{ {{range $ti, $tab := $.Datas.AllTables}}
		// table: {{$tab.HeaderType}}
		public List<{{$tab.HeaderType}}> {{$tab.HeaderType}} = new List<{{$tab.HeaderType}}>(); {{end}}

		// Indices
		{{range $ii, $idx := GetIndices $}}{{if IsWarpFieldName $ $idx.FieldInfo}}public Dictionary<{{CSType $idx.FieldInfo}},{{$idx.Table.HeaderType}}> {{$idx.Table.HeaderType}}By{{$idx.FieldInfo.FieldName}} = new Dictionary<{{CSType $idx.FieldInfo}},{{$idx.Table.HeaderType}}>();
		{{end}}{{end}}
		
		{{if HasKeyValueTypes $}}
		{{range $ti, $name := GetKeyValueTypeNames $}}
		// table: {{$name}}
		public {{$name}} GetKeyValue_{{$name}}()
		{
			return {{$name}}[0];
		}{{end}}{{end}}

		public void ResetData( )
		{   {{range $ti, $tab := $.Datas.AllTables}}
			{{$tab.HeaderType}}.Clear(); {{end}} 
			{{range $ii, $idx := GetIndices $}}{{if IsWarpFieldName $ $idx.FieldInfo}}{{$idx.Table.HeaderType}}By{{$idx.FieldInfo.FieldName}}.Clear();
			{{end}}{{end}}	
		}
		{{if $.GenBinary}}
		public void Deserialize( tabtoy.TableReader reader )
		{	
			reader.ReadHeader();

			UInt32 mamaSaidTagNameShouldBeLong = 0;
            while ( reader.ReadTag(ref mamaSaidTagNameShouldBeLong) )
            {
				if (mamaSaidTagNameShouldBeLong == 0x6f0000)
				{
                    var tabName = string.Empty;
                    reader.ReadString(ref tabName);
					switch (tabName)
					{ {{range $ti, $tab := $.Datas.AllTables}}
						case "{{$tab.HeaderType}}":
						{
							reader.ReadStruct(ref {{$tab.HeaderType}});	
						}
						break;{{end}}
						default:
						{
							reader.SkipFiled(mamaSaidTagNameShouldBeLong);                            
						}
						break;
					}
				}
			}
		}

		public void IndexData( string tabName = "")
		{ {{range $ii, $idx := GetIndices $}}	
			if (tabName == "" || tabName == "{{$idx.Table.HeaderType}}")
			{
				{{if IsWarpFieldName $ $idx.FieldInfo}}foreach( var kv in {{$idx.Table.HeaderType}} )
				{
					{{$idx.Table.HeaderType}}By{{$idx.FieldInfo.FieldName}}[kv.{{$idx.FieldInfo.FieldName}}] = kv;
				}
				{{end}}
			}
			{{end}}
		}
		{{end}}
	}
}
`
