package sql2struct

import (
	"fmt"
	"github.com/eighteen20/goTools/internal/word"
	"html/template"
	"os"
)

/*
 定义生成结构体的模板
*/
const structTpl = `
type {{.TableName | ToCamelCase}} struct {
	{{range .Columns}} {{ $length := len .Comment}} {{ if gt $length 0 }} 
	// {{.Comment}} {{else}} // {{.Name}} 
	{{end}}
	{{ $typeLen := len .Type}}{{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}    {{ .Type}}    {{.Tag}}{{ else }}{{.Name}}{{ end }} 
	{{end}}
}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}
`

type StructTemplate struct {
	StructTpl string
}

/*
 存储转换够的Go结构体中的所有字段信息
*/
type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

/*
 存储最终用于渲染的模板对象信息
*/
type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{StructTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     fmt.Sprintf("`json:"+"%s"+"`", column.ColumnName),
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("spl2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.StructTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}

	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}
