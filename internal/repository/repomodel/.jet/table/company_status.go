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

var CompanyStatus = newCompanyStatusTable("public", "company_status", "")

type companyStatusTable struct {
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

type CompanyStatusTable struct {
	companyStatusTable

	EXCLUDED companyStatusTable
}

// AS creates new CompanyStatusTable with assigned alias
func (a CompanyStatusTable) AS(alias string) *CompanyStatusTable {
	return newCompanyStatusTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CompanyStatusTable with assigned schema name
func (a CompanyStatusTable) FromSchema(schemaName string) *CompanyStatusTable {
	return newCompanyStatusTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CompanyStatusTable with assigned table prefix
func (a CompanyStatusTable) WithPrefix(prefix string) *CompanyStatusTable {
	return newCompanyStatusTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CompanyStatusTable with assigned table suffix
func (a CompanyStatusTable) WithSuffix(suffix string) *CompanyStatusTable {
	return newCompanyStatusTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCompanyStatusTable(schemaName, tableName, alias string) *CompanyStatusTable {
	return &CompanyStatusTable{
		companyStatusTable: newCompanyStatusTableImpl(schemaName, tableName, alias),
		EXCLUDED:           newCompanyStatusTableImpl("", "excluded", ""),
	}
}

func newCompanyStatusTableImpl(schemaName, tableName, alias string) companyStatusTable {
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

	return companyStatusTable{
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
