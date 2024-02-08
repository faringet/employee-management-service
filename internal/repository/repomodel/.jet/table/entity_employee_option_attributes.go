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

var EntityEmployeeOptionAttributes = newEntityEmployeeOptionAttributesTable("public", "entity_employee_option_attributes", "")

type entityEmployeeOptionAttributesTable struct {
	postgres.Table

	// Columns
	ID                postgres.ColumnInteger
	Value             postgres.ColumnString
	CreatedAt         postgres.ColumnTimestamp
	UpdatedAt         postgres.ColumnTimestamp
	EntityAttributeID postgres.ColumnInteger
	EntityUserID      postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type EntityEmployeeOptionAttributesTable struct {
	entityEmployeeOptionAttributesTable

	EXCLUDED entityEmployeeOptionAttributesTable
}

// AS creates new EntityEmployeeOptionAttributesTable with assigned alias
func (a EntityEmployeeOptionAttributesTable) AS(alias string) *EntityEmployeeOptionAttributesTable {
	return newEntityEmployeeOptionAttributesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EntityEmployeeOptionAttributesTable with assigned schema name
func (a EntityEmployeeOptionAttributesTable) FromSchema(schemaName string) *EntityEmployeeOptionAttributesTable {
	return newEntityEmployeeOptionAttributesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EntityEmployeeOptionAttributesTable with assigned table prefix
func (a EntityEmployeeOptionAttributesTable) WithPrefix(prefix string) *EntityEmployeeOptionAttributesTable {
	return newEntityEmployeeOptionAttributesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EntityEmployeeOptionAttributesTable with assigned table suffix
func (a EntityEmployeeOptionAttributesTable) WithSuffix(suffix string) *EntityEmployeeOptionAttributesTable {
	return newEntityEmployeeOptionAttributesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEntityEmployeeOptionAttributesTable(schemaName, tableName, alias string) *EntityEmployeeOptionAttributesTable {
	return &EntityEmployeeOptionAttributesTable{
		entityEmployeeOptionAttributesTable: newEntityEmployeeOptionAttributesTableImpl(schemaName, tableName, alias),
		EXCLUDED:                            newEntityEmployeeOptionAttributesTableImpl("", "excluded", ""),
	}
}

func newEntityEmployeeOptionAttributesTableImpl(schemaName, tableName, alias string) entityEmployeeOptionAttributesTable {
	var (
		IDColumn                = postgres.IntegerColumn("id")
		ValueColumn             = postgres.StringColumn("value")
		CreatedAtColumn         = postgres.TimestampColumn("created_at")
		UpdatedAtColumn         = postgres.TimestampColumn("updated_at")
		EntityAttributeIDColumn = postgres.IntegerColumn("entity_attribute_id")
		EntityUserIDColumn      = postgres.IntegerColumn("entity_user_id")
		allColumns              = postgres.ColumnList{IDColumn, ValueColumn, CreatedAtColumn, UpdatedAtColumn, EntityAttributeIDColumn, EntityUserIDColumn}
		mutableColumns          = postgres.ColumnList{ValueColumn, CreatedAtColumn, UpdatedAtColumn, EntityAttributeIDColumn, EntityUserIDColumn}
	)

	return entityEmployeeOptionAttributesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                IDColumn,
		Value:             ValueColumn,
		CreatedAt:         CreatedAtColumn,
		UpdatedAt:         UpdatedAtColumn,
		EntityAttributeID: EntityAttributeIDColumn,
		EntityUserID:      EntityUserIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
