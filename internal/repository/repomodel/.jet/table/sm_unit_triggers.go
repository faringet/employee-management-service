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

var SmUnitTriggers = newSmUnitTriggersTable("public", "sm_unit_triggers", "")

type smUnitTriggersTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnInteger
	UpdatedAt postgres.ColumnTimestamp
	CreatedAt postgres.ColumnTimestamp
	UpdatedBy postgres.ColumnInteger
	CreatedBy postgres.ColumnInteger
	ProjectID postgres.ColumnInteger
	UnitID    postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SmUnitTriggersTable struct {
	smUnitTriggersTable

	EXCLUDED smUnitTriggersTable
}

// AS creates new SmUnitTriggersTable with assigned alias
func (a SmUnitTriggersTable) AS(alias string) *SmUnitTriggersTable {
	return newSmUnitTriggersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SmUnitTriggersTable with assigned schema name
func (a SmUnitTriggersTable) FromSchema(schemaName string) *SmUnitTriggersTable {
	return newSmUnitTriggersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SmUnitTriggersTable with assigned table prefix
func (a SmUnitTriggersTable) WithPrefix(prefix string) *SmUnitTriggersTable {
	return newSmUnitTriggersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SmUnitTriggersTable with assigned table suffix
func (a SmUnitTriggersTable) WithSuffix(suffix string) *SmUnitTriggersTable {
	return newSmUnitTriggersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSmUnitTriggersTable(schemaName, tableName, alias string) *SmUnitTriggersTable {
	return &SmUnitTriggersTable{
		smUnitTriggersTable: newSmUnitTriggersTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newSmUnitTriggersTableImpl("", "excluded", ""),
	}
}

func newSmUnitTriggersTableImpl(schemaName, tableName, alias string) smUnitTriggersTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		UpdatedAtColumn = postgres.TimestampColumn("updated_at")
		CreatedAtColumn = postgres.TimestampColumn("created_at")
		UpdatedByColumn = postgres.IntegerColumn("updated_by")
		CreatedByColumn = postgres.IntegerColumn("created_by")
		ProjectIDColumn = postgres.IntegerColumn("project_id")
		UnitIDColumn    = postgres.IntegerColumn("unit_id")
		allColumns      = postgres.ColumnList{IDColumn, UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, ProjectIDColumn, UnitIDColumn}
		mutableColumns  = postgres.ColumnList{UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, ProjectIDColumn, UnitIDColumn}
	)

	return smUnitTriggersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		UpdatedAt: UpdatedAtColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedBy: UpdatedByColumn,
		CreatedBy: CreatedByColumn,
		ProjectID: ProjectIDColumn,
		UnitID:    UnitIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}