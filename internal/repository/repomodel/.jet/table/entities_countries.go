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

var EntitiesCountries = newEntitiesCountriesTable("public", "entities_countries", "")

type entitiesCountriesTable struct {
	postgres.Table

	// Columns
	ID         postgres.ColumnInteger
	EntitiesID postgres.ColumnInteger
	CountryID  postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type EntitiesCountriesTable struct {
	entitiesCountriesTable

	EXCLUDED entitiesCountriesTable
}

// AS creates new EntitiesCountriesTable with assigned alias
func (a EntitiesCountriesTable) AS(alias string) *EntitiesCountriesTable {
	return newEntitiesCountriesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EntitiesCountriesTable with assigned schema name
func (a EntitiesCountriesTable) FromSchema(schemaName string) *EntitiesCountriesTable {
	return newEntitiesCountriesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EntitiesCountriesTable with assigned table prefix
func (a EntitiesCountriesTable) WithPrefix(prefix string) *EntitiesCountriesTable {
	return newEntitiesCountriesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EntitiesCountriesTable with assigned table suffix
func (a EntitiesCountriesTable) WithSuffix(suffix string) *EntitiesCountriesTable {
	return newEntitiesCountriesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEntitiesCountriesTable(schemaName, tableName, alias string) *EntitiesCountriesTable {
	return &EntitiesCountriesTable{
		entitiesCountriesTable: newEntitiesCountriesTableImpl(schemaName, tableName, alias),
		EXCLUDED:               newEntitiesCountriesTableImpl("", "excluded", ""),
	}
}

func newEntitiesCountriesTableImpl(schemaName, tableName, alias string) entitiesCountriesTable {
	var (
		IDColumn         = postgres.IntegerColumn("id")
		EntitiesIDColumn = postgres.IntegerColumn("entities_id")
		CountryIDColumn  = postgres.IntegerColumn("country_id")
		allColumns       = postgres.ColumnList{IDColumn, EntitiesIDColumn, CountryIDColumn}
		mutableColumns   = postgres.ColumnList{EntitiesIDColumn, CountryIDColumn}
	)

	return entitiesCountriesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		EntitiesID: EntitiesIDColumn,
		CountryID:  CountryIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
