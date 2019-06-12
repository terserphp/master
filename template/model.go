package template

var Model = `// auto-generated by terser-cli {{.GenerateTime}}
// struct: {{.StructName}} {{.Comment}}
package model

import (
	"github.com/jinzhu/gorm"
{{if .HasDateTime}}	"time"{{end}}
)

// {{.StructName}} struct {{.Comment}}
type {{.StructName}} struct { 
{{range .ColumnList}}	{{.PropertyName}}	{{if .IsNullable}}*{{end}}{{.GoDataType}}		{{$.LabelTag}}gorm:"column:{{.Name}};type:{{.ColumnType}};{{if .DefaultValue}}default:{{if .IsNumeral}}{{.DefaultValue}}{{else}}'{{.DefaultValue}}'{{end}};{{end}}{{if .IsPrimaryKey}}primary_key;{{end}}" json:"{{.Name}}"{{$.LabelTag}}	// {{.Comment}}{{if .IsNullable}} [null type]{{end}}{{if .DefaultValue}} [default: {{.DefaultValue}}]{{end}}
{{end}}}

//  {{.StructName}} 表字段名常量
const (
	{{.StructName}}_TableName  = "{{.Name}}" // 表名: {{.Comment}}
{{range  .ColumnList}}	{{$.StructName}}_{{.PropertyName}} TableField = "{{.Name}}" // 字段名: {{.Comment}}
{{end}})

// {{.StructName}} 数据库{{.Name}}查询对象
func New{{.StructName}}Query() (db *gorm.DB) {
	db = GetDefaultDB()

	return db.Table({{.StructName}}_TableName)
}

// 获取{{.StructName}}表名
func ({{.VarName}} *{{.StructName}}) TableName() string {
	return {{.StructName}}_TableName
}

func ({{.VarName}} *{{.StructName}}) BeforeCreate(scope *gorm.Scope) error {
	// 执行create时为某些字段赋初值
	//scope.Set("字段名", 默认值)
	return nil
}

func ({{.VarName}} *{{.StructName}}) Create({{if or .CreateUserKey}}adminUserID string{{end}}) error {
{{if .CreateUserKey}}	{{.VarName}}.{{.CreateUserKey}} = adminUserID
	{{if .UpdateUserKey}}{{.VarName}}.{{.UpdateUserKey}} = adminUserID{{end}}

{{end}}	return New{{.StructName}}Query().
	    Create(&{{.VarName}}).Error
}

func ({{.VarName}} *{{.StructName}}) Save({{if .UpdateUserKey}}adminUserID string{{end}}) error {
{{if .UpdateUserKey}}	{{.VarName}}.{{.UpdateUserKey}} = adminUserID

{{end}}	return New{{.StructName}}Query().
	    Save(&{{.VarName}}).Error
}
{{ if .LogicDeleteKey }}
func ({{ .VarName }} *{{ .StructName }}) Delete({{ if .UpdateUserKey }}adminUserID string{{ end }}) error {
	{{ .VarName }}.{{ .LogicDeleteKey }} = 1
{{ if .UpdateUserKey }}	{{ .VarName }}.{{ .UpdateUserKey }} = adminUserID

{{ end }}	return {{ .VarName }}.Save({{ if .UpdateUserKey }}adminUserID{{ end }})
}
{{ else }}
func ({{ .VarName }} *{{ .StructName }}) Delete() error {
	return New{{ .StructName }}Query().
	    Delete(&{{ .VarName }}).Error
}
{{ end }}
{{if .HasPrimaryKey}}
func ({{.VarName}} *{{.StructName}}) Load() error {
	return New{{.StructName}}Query().
		Query(&{{.VarName}}).First(&{{.VarName}}).Error
}

func Find{{.StructName}}ByID({{range .PrimaryKeys}}{{if gt .Index 0}}, {{end}}{{.VarName}} {{.GoDataType}}{{end}}) ({{.VarName}} *{{.StructName}}, err error) {
	{{.VarName}} = &{{.StructName}}{}
	err = New{{.StructName}}Query().
		First(&{{.VarName}}, "{{range .PrimaryKeys}}{{if gt .Index 0}} and {{end}}{{.Name}} = ?{{end}}{{if .LogicDeleteKey}} and is_deleted = 0 {{end}}"{{range .PrimaryKeys}}, {{.VarName}}{{end}}).Error

	return {{.VarName}}, err
}

func Delete{{.StructName}}ByID({{range .PrimaryKeys}}{{if gt .Index 0}}, {{end}}{{.VarName}} {{.GoDataType}}{{end}}{{if and .UpdateUserKey .LogicDeleteKey}}, adminUserID string{{end}}) (err error) {
{{if .LogicDeleteKey}}	update := map[string]interface{}{
		{{.StructName}}_{{.LogicDeleteKey}}:  1,{{if .UpdateUserKey}}
		{{.StructName}}_{{.UpdateUserKey}}: adminUserID,{{end}}
	}

	return New{{.StructName}}Query().
		Query("{{range .PrimaryKeys}}{{if gt .Index 0}} and {{end}}{{.Name}} = ?{{end}} "{{range .PrimaryKeys}}, {{.VarName}}{{end}}).
		Updates(update).Error
{{else}}	return New{{.StructName}}Query().
		Query("{{range .PrimaryKeys}}{{if gt .Index 0}} and {{end}}{{.Name}} = ?{{end}} "{{range .PrimaryKeys}}, {{.VarName}}{{end}}).
		Delete({{.StructName}}{}).Error
{{end}}}

func Update{{.StructName}}ByID({{range .PrimaryKeys}}{{if gt .Index 0}}, {{end}}{{.VarName}} {{.GoDataType}}{{end}}, updates interface{}) (err error) {
	return New{{.StructName}}Query().
		Query("{{range .PrimaryKeys}}{{if gt .Index 0}} and {{end}}{{.Name}} = ?{{end}} "{{range .PrimaryKeys}}, {{.VarName}}{{end}}).
		Updates(updates).Error
}
{{end}}

`
