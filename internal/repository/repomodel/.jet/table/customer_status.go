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

var CustomerStatus = newCustomerStatusTable("public", "customer_status", "")

type customerStatusTable struct {
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

type CustomerStatusTable struct {
	customerStatusTable

	EXCLUDED customerStatusTable
}

// AS creates new CustomerStatusTable with assigned alias
func (a CustomerStatusTable) AS(alias string) *CustomerStatusTable {
	return newCustomerStatusTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CustomerStatusTable with assigned schema name
func (a CustomerStatusTable) FromSchema(schemaName string) *CustomerStatusTable {
	return newCustomerStatusTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CustomerStatusTable with assigned table prefix
func (a CustomerStatusTable) WithPrefix(prefix string) *CustomerStatusTable {
	return newCustomerStatusTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CustomerStatusTable with assigned table suffix
func (a CustomerStatusTable) WithSuffix(suffix string) *CustomerStatusTable {
	return newCustomerStatusTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCustomerStatusTable(schemaName, tableName, alias string) *CustomerStatusTable {
	return &CustomerStatusTable{
		customerStatusTable: newCustomerStatusTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newCustomerStatusTableImpl("", "excluded", ""),
	}
}

func newCustomerStatusTableImpl(schemaName, tableName, alias string) customerStatusTable {
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

	return customerStatusTable{
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
