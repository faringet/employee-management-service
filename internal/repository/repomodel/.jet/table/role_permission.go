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

var RolePermission = newRolePermissionTable("public", "role_permission", "")

type rolePermissionTable struct {
	postgres.Table

	// Columns
	ID                postgres.ColumnInteger
	RoleID            postgres.ColumnInteger
	LevelPermissionID postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type RolePermissionTable struct {
	rolePermissionTable

	EXCLUDED rolePermissionTable
}

// AS creates new RolePermissionTable with assigned alias
func (a RolePermissionTable) AS(alias string) *RolePermissionTable {
	return newRolePermissionTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RolePermissionTable with assigned schema name
func (a RolePermissionTable) FromSchema(schemaName string) *RolePermissionTable {
	return newRolePermissionTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RolePermissionTable with assigned table prefix
func (a RolePermissionTable) WithPrefix(prefix string) *RolePermissionTable {
	return newRolePermissionTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RolePermissionTable with assigned table suffix
func (a RolePermissionTable) WithSuffix(suffix string) *RolePermissionTable {
	return newRolePermissionTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRolePermissionTable(schemaName, tableName, alias string) *RolePermissionTable {
	return &RolePermissionTable{
		rolePermissionTable: newRolePermissionTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newRolePermissionTableImpl("", "excluded", ""),
	}
}

func newRolePermissionTableImpl(schemaName, tableName, alias string) rolePermissionTable {
	var (
		IDColumn                = postgres.IntegerColumn("id")
		RoleIDColumn            = postgres.IntegerColumn("role_id")
		LevelPermissionIDColumn = postgres.IntegerColumn("level_permission_id")
		allColumns              = postgres.ColumnList{IDColumn, RoleIDColumn, LevelPermissionIDColumn}
		mutableColumns          = postgres.ColumnList{RoleIDColumn, LevelPermissionIDColumn}
	)

	return rolePermissionTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                IDColumn,
		RoleID:            RoleIDColumn,
		LevelPermissionID: LevelPermissionIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}