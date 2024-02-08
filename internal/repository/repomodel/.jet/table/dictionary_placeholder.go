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

var DictionaryPlaceholder = newDictionaryPlaceholderTable("public", "dictionary_placeholder", "")

type dictionaryPlaceholderTable struct {
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

type DictionaryPlaceholderTable struct {
	dictionaryPlaceholderTable

	EXCLUDED dictionaryPlaceholderTable
}

// AS creates new DictionaryPlaceholderTable with assigned alias
func (a DictionaryPlaceholderTable) AS(alias string) *DictionaryPlaceholderTable {
	return newDictionaryPlaceholderTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DictionaryPlaceholderTable with assigned schema name
func (a DictionaryPlaceholderTable) FromSchema(schemaName string) *DictionaryPlaceholderTable {
	return newDictionaryPlaceholderTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new DictionaryPlaceholderTable with assigned table prefix
func (a DictionaryPlaceholderTable) WithPrefix(prefix string) *DictionaryPlaceholderTable {
	return newDictionaryPlaceholderTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new DictionaryPlaceholderTable with assigned table suffix
func (a DictionaryPlaceholderTable) WithSuffix(suffix string) *DictionaryPlaceholderTable {
	return newDictionaryPlaceholderTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newDictionaryPlaceholderTable(schemaName, tableName, alias string) *DictionaryPlaceholderTable {
	return &DictionaryPlaceholderTable{
		dictionaryPlaceholderTable: newDictionaryPlaceholderTableImpl(schemaName, tableName, alias),
		EXCLUDED:                   newDictionaryPlaceholderTableImpl("", "excluded", ""),
	}
}

func newDictionaryPlaceholderTableImpl(schemaName, tableName, alias string) dictionaryPlaceholderTable {
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

	return dictionaryPlaceholderTable{
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
