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

var EntityAttributes = newEntityAttributesTable("public", "entity_attributes", "")

type entityAttributesTable struct {
	postgres.Table

	// Columns
	ID              postgres.ColumnInteger
	Name            postgres.ColumnString
	Type            postgres.ColumnString
	EntityID        postgres.ColumnInteger
	CreatedAt       postgres.ColumnTimestamp
	UpdatedAt       postgres.ColumnTimestamp
	CreatedBy       postgres.ColumnInteger
	UpdatedBy       postgres.ColumnInteger
	IDParentattr    postgres.ColumnInteger
	IDAttributetype postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type EntityAttributesTable struct {
	entityAttributesTable

	EXCLUDED entityAttributesTable
}

// AS creates new EntityAttributesTable with assigned alias
func (a EntityAttributesTable) AS(alias string) *EntityAttributesTable {
	return newEntityAttributesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EntityAttributesTable with assigned schema name
func (a EntityAttributesTable) FromSchema(schemaName string) *EntityAttributesTable {
	return newEntityAttributesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EntityAttributesTable with assigned table prefix
func (a EntityAttributesTable) WithPrefix(prefix string) *EntityAttributesTable {
	return newEntityAttributesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EntityAttributesTable with assigned table suffix
func (a EntityAttributesTable) WithSuffix(suffix string) *EntityAttributesTable {
	return newEntityAttributesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEntityAttributesTable(schemaName, tableName, alias string) *EntityAttributesTable {
	return &EntityAttributesTable{
		entityAttributesTable: newEntityAttributesTableImpl(schemaName, tableName, alias),
		EXCLUDED:              newEntityAttributesTableImpl("", "excluded", ""),
	}
}

func newEntityAttributesTableImpl(schemaName, tableName, alias string) entityAttributesTable {
	var (
		IDColumn              = postgres.IntegerColumn("id")
		NameColumn            = postgres.StringColumn("name")
		TypeColumn            = postgres.StringColumn("type")
		EntityIDColumn        = postgres.IntegerColumn("entity_id")
		CreatedAtColumn       = postgres.TimestampColumn("created_at")
		UpdatedAtColumn       = postgres.TimestampColumn("updated_at")
		CreatedByColumn       = postgres.IntegerColumn("created_by")
		UpdatedByColumn       = postgres.IntegerColumn("updated_by")
		IDParentattrColumn    = postgres.IntegerColumn("id_parentAttr")
		IDAttributetypeColumn = postgres.IntegerColumn("id_attributetype")
		allColumns            = postgres.ColumnList{IDColumn, NameColumn, TypeColumn, EntityIDColumn, CreatedAtColumn, UpdatedAtColumn, CreatedByColumn, UpdatedByColumn, IDParentattrColumn, IDAttributetypeColumn}
		mutableColumns        = postgres.ColumnList{NameColumn, TypeColumn, EntityIDColumn, CreatedAtColumn, UpdatedAtColumn, CreatedByColumn, UpdatedByColumn, IDParentattrColumn, IDAttributetypeColumn}
	)

	return entityAttributesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		Name:            NameColumn,
		Type:            TypeColumn,
		EntityID:        EntityIDColumn,
		CreatedAt:       CreatedAtColumn,
		UpdatedAt:       UpdatedAtColumn,
		CreatedBy:       CreatedByColumn,
		UpdatedBy:       UpdatedByColumn,
		IDParentattr:    IDParentattrColumn,
		IDAttributetype: IDAttributetypeColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
