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

var WorkspaceSettings = newWorkspaceSettingsTable("public", "workspace_settings", "")

type workspaceSettingsTable struct {
	postgres.Table

	// Columns
	ID              postgres.ColumnInteger
	Name            postgres.ColumnString
	AdminCount      postgres.ColumnInteger
	UserCount       postgres.ColumnInteger
	EmployeeCount   postgres.ColumnInteger
	ActiveSeatCount postgres.ColumnInteger
	WorkspaceID     postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type WorkspaceSettingsTable struct {
	workspaceSettingsTable

	EXCLUDED workspaceSettingsTable
}

// AS creates new WorkspaceSettingsTable with assigned alias
func (a WorkspaceSettingsTable) AS(alias string) *WorkspaceSettingsTable {
	return newWorkspaceSettingsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new WorkspaceSettingsTable with assigned schema name
func (a WorkspaceSettingsTable) FromSchema(schemaName string) *WorkspaceSettingsTable {
	return newWorkspaceSettingsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new WorkspaceSettingsTable with assigned table prefix
func (a WorkspaceSettingsTable) WithPrefix(prefix string) *WorkspaceSettingsTable {
	return newWorkspaceSettingsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new WorkspaceSettingsTable with assigned table suffix
func (a WorkspaceSettingsTable) WithSuffix(suffix string) *WorkspaceSettingsTable {
	return newWorkspaceSettingsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newWorkspaceSettingsTable(schemaName, tableName, alias string) *WorkspaceSettingsTable {
	return &WorkspaceSettingsTable{
		workspaceSettingsTable: newWorkspaceSettingsTableImpl(schemaName, tableName, alias),
		EXCLUDED:               newWorkspaceSettingsTableImpl("", "excluded", ""),
	}
}

func newWorkspaceSettingsTableImpl(schemaName, tableName, alias string) workspaceSettingsTable {
	var (
		IDColumn              = postgres.IntegerColumn("id")
		NameColumn            = postgres.StringColumn("name")
		AdminCountColumn      = postgres.IntegerColumn("admin_count")
		UserCountColumn       = postgres.IntegerColumn("user_count")
		EmployeeCountColumn   = postgres.IntegerColumn("employee_count")
		ActiveSeatCountColumn = postgres.IntegerColumn("active_seat_count")
		WorkspaceIDColumn     = postgres.IntegerColumn("workspace_id")
		allColumns            = postgres.ColumnList{IDColumn, NameColumn, AdminCountColumn, UserCountColumn, EmployeeCountColumn, ActiveSeatCountColumn, WorkspaceIDColumn}
		mutableColumns        = postgres.ColumnList{NameColumn, AdminCountColumn, UserCountColumn, EmployeeCountColumn, ActiveSeatCountColumn, WorkspaceIDColumn}
	)

	return workspaceSettingsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		Name:            NameColumn,
		AdminCount:      AdminCountColumn,
		UserCount:       UserCountColumn,
		EmployeeCount:   EmployeeCountColumn,
		ActiveSeatCount: ActiveSeatCountColumn,
		WorkspaceID:     WorkspaceIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}