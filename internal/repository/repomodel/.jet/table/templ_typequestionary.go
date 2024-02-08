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

var TemplTypequestionary = newTemplTypequestionaryTable("public", "templ_typequestionary", "")

type templTypequestionaryTable struct {
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

type TemplTypequestionaryTable struct {
	templTypequestionaryTable

	EXCLUDED templTypequestionaryTable
}

// AS creates new TemplTypequestionaryTable with assigned alias
func (a TemplTypequestionaryTable) AS(alias string) *TemplTypequestionaryTable {
	return newTemplTypequestionaryTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TemplTypequestionaryTable with assigned schema name
func (a TemplTypequestionaryTable) FromSchema(schemaName string) *TemplTypequestionaryTable {
	return newTemplTypequestionaryTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TemplTypequestionaryTable with assigned table prefix
func (a TemplTypequestionaryTable) WithPrefix(prefix string) *TemplTypequestionaryTable {
	return newTemplTypequestionaryTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TemplTypequestionaryTable with assigned table suffix
func (a TemplTypequestionaryTable) WithSuffix(suffix string) *TemplTypequestionaryTable {
	return newTemplTypequestionaryTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTemplTypequestionaryTable(schemaName, tableName, alias string) *TemplTypequestionaryTable {
	return &TemplTypequestionaryTable{
		templTypequestionaryTable: newTemplTypequestionaryTableImpl(schemaName, tableName, alias),
		EXCLUDED:                  newTemplTypequestionaryTableImpl("", "excluded", ""),
	}
}

func newTemplTypequestionaryTableImpl(schemaName, tableName, alias string) templTypequestionaryTable {
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

	return templTypequestionaryTable{
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