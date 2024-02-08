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

var ImportType = newImportTypeTable("public", "import_type", "")

type importTypeTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnInteger
	Name        postgres.ColumnString
	Code        postgres.ColumnString
	Description postgres.ColumnString
	CreatedAt   postgres.ColumnTimestamp
	UpdatedAt   postgres.ColumnTimestamp
	CreatedBy   postgres.ColumnInteger
	UpdatedBy   postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ImportTypeTable struct {
	importTypeTable

	EXCLUDED importTypeTable
}

// AS creates new ImportTypeTable with assigned alias
func (a ImportTypeTable) AS(alias string) *ImportTypeTable {
	return newImportTypeTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ImportTypeTable with assigned schema name
func (a ImportTypeTable) FromSchema(schemaName string) *ImportTypeTable {
	return newImportTypeTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ImportTypeTable with assigned table prefix
func (a ImportTypeTable) WithPrefix(prefix string) *ImportTypeTable {
	return newImportTypeTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ImportTypeTable with assigned table suffix
func (a ImportTypeTable) WithSuffix(suffix string) *ImportTypeTable {
	return newImportTypeTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newImportTypeTable(schemaName, tableName, alias string) *ImportTypeTable {
	return &ImportTypeTable{
		importTypeTable: newImportTypeTableImpl(schemaName, tableName, alias),
		EXCLUDED:        newImportTypeTableImpl("", "excluded", ""),
	}
}

func newImportTypeTableImpl(schemaName, tableName, alias string) importTypeTable {
	var (
		IDColumn          = postgres.IntegerColumn("id")
		NameColumn        = postgres.StringColumn("name")
		CodeColumn        = postgres.StringColumn("code")
		DescriptionColumn = postgres.StringColumn("description")
		CreatedAtColumn   = postgres.TimestampColumn("created_at")
		UpdatedAtColumn   = postgres.TimestampColumn("updated_at")
		CreatedByColumn   = postgres.IntegerColumn("created_by")
		UpdatedByColumn   = postgres.IntegerColumn("updated_by")
		allColumns        = postgres.ColumnList{IDColumn, NameColumn, CodeColumn, DescriptionColumn, CreatedAtColumn, UpdatedAtColumn, CreatedByColumn, UpdatedByColumn}
		mutableColumns    = postgres.ColumnList{NameColumn, CodeColumn, DescriptionColumn, CreatedAtColumn, UpdatedAtColumn, CreatedByColumn, UpdatedByColumn}
	)

	return importTypeTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Name:        NameColumn,
		Code:        CodeColumn,
		Description: DescriptionColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,
		CreatedBy:   CreatedByColumn,
		UpdatedBy:   UpdatedByColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
