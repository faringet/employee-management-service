//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var TemplateDictionary = newTemplateDictionaryTable("public", "Template_Dictionary", "")

type templateDictionaryTable struct {
	postgres.Table

	// Columns
	ID              postgres.ColumnInteger
	Name            postgres.ColumnString
	Code            postgres.ColumnString
	Description     postgres.ColumnString
	UpdatedAt       postgres.ColumnTimestamp
	CreatedAt       postgres.ColumnTimestamp
	UpdatedBy       postgres.ColumnInteger
	CreatedBy       postgres.ColumnInteger
	QueueNumber     postgres.ColumnInteger
	IconColor       postgres.ColumnString
	IdCustomSvgIcon postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TemplateDictionaryTable struct {
	templateDictionaryTable

	EXCLUDED templateDictionaryTable
}

// AS creates new TemplateDictionaryTable with assigned alias
func (a TemplateDictionaryTable) AS(alias string) *TemplateDictionaryTable {
	return newTemplateDictionaryTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TemplateDictionaryTable with assigned schema name
func (a TemplateDictionaryTable) FromSchema(schemaName string) *TemplateDictionaryTable {
	return newTemplateDictionaryTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TemplateDictionaryTable with assigned table prefix
func (a TemplateDictionaryTable) WithPrefix(prefix string) *TemplateDictionaryTable {
	return newTemplateDictionaryTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TemplateDictionaryTable with assigned table suffix
func (a TemplateDictionaryTable) WithSuffix(suffix string) *TemplateDictionaryTable {
	return newTemplateDictionaryTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTemplateDictionaryTable(schemaName, tableName, alias string) *TemplateDictionaryTable {
	return &TemplateDictionaryTable{
		templateDictionaryTable: newTemplateDictionaryTableImpl(schemaName, tableName, alias),
		EXCLUDED:                newTemplateDictionaryTableImpl("", "excluded", ""),
	}
}

func newTemplateDictionaryTableImpl(schemaName, tableName, alias string) templateDictionaryTable {
	var (
		IDColumn              = postgres.IntegerColumn("id")
		NameColumn            = postgres.StringColumn("name")
		CodeColumn            = postgres.StringColumn("code")
		DescriptionColumn     = postgres.StringColumn("description")
		UpdatedAtColumn       = postgres.TimestampColumn("updated_at")
		CreatedAtColumn       = postgres.TimestampColumn("created_at")
		UpdatedByColumn       = postgres.IntegerColumn("updated_by")
		CreatedByColumn       = postgres.IntegerColumn("created_by")
		QueueNumberColumn     = postgres.IntegerColumn("queueNumber")
		IconColorColumn       = postgres.StringColumn("iconColor")
		IdCustomSvgIconColumn = postgres.IntegerColumn("idCustomSvgIcon")
		allColumns            = postgres.ColumnList{IDColumn, NameColumn, CodeColumn, DescriptionColumn, UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, QueueNumberColumn, IconColorColumn, IdCustomSvgIconColumn}
		mutableColumns        = postgres.ColumnList{NameColumn, CodeColumn, DescriptionColumn, UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, QueueNumberColumn, IconColorColumn, IdCustomSvgIconColumn}
	)

	return templateDictionaryTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		Name:            NameColumn,
		Code:            CodeColumn,
		Description:     DescriptionColumn,
		UpdatedAt:       UpdatedAtColumn,
		CreatedAt:       CreatedAtColumn,
		UpdatedBy:       UpdatedByColumn,
		CreatedBy:       CreatedByColumn,
		QueueNumber:     QueueNumberColumn,
		IconColor:       IconColorColumn,
		IdCustomSvgIcon: IdCustomSvgIconColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}