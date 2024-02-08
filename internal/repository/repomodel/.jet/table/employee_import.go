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

var EmployeeImport = newEmployeeImportTable("public", "employee_import", "")

type employeeImportTable struct {
	postgres.Table

	// Columns
	ID                    postgres.ColumnInteger
	ImportedAt            postgres.ColumnTimestamp
	CreatedAt             postgres.ColumnTimestamp
	UpdatedAt             postgres.ColumnTimestamp
	RemoveAbsentEmployees postgres.ColumnBool
	EntityID              postgres.ColumnInteger
	ImportTypeID          postgres.ColumnInteger
	CreatedBy             postgres.ColumnInteger
	UpdatedBy             postgres.ColumnInteger
	ChangesImport         postgres.ColumnString
	ErrorsImport          postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type EmployeeImportTable struct {
	employeeImportTable

	EXCLUDED employeeImportTable
}

// AS creates new EmployeeImportTable with assigned alias
func (a EmployeeImportTable) AS(alias string) *EmployeeImportTable {
	return newEmployeeImportTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EmployeeImportTable with assigned schema name
func (a EmployeeImportTable) FromSchema(schemaName string) *EmployeeImportTable {
	return newEmployeeImportTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EmployeeImportTable with assigned table prefix
func (a EmployeeImportTable) WithPrefix(prefix string) *EmployeeImportTable {
	return newEmployeeImportTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EmployeeImportTable with assigned table suffix
func (a EmployeeImportTable) WithSuffix(suffix string) *EmployeeImportTable {
	return newEmployeeImportTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEmployeeImportTable(schemaName, tableName, alias string) *EmployeeImportTable {
	return &EmployeeImportTable{
		employeeImportTable: newEmployeeImportTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newEmployeeImportTableImpl("", "excluded", ""),
	}
}

func newEmployeeImportTableImpl(schemaName, tableName, alias string) employeeImportTable {
	var (
		IDColumn                    = postgres.IntegerColumn("id")
		ImportedAtColumn            = postgres.TimestampColumn("imported_at")
		CreatedAtColumn             = postgres.TimestampColumn("created_at")
		UpdatedAtColumn             = postgres.TimestampColumn("updated_at")
		RemoveAbsentEmployeesColumn = postgres.BoolColumn("remove_absent_employees")
		EntityIDColumn              = postgres.IntegerColumn("entity_id")
		ImportTypeIDColumn          = postgres.IntegerColumn("import_type_id")
		CreatedByColumn             = postgres.IntegerColumn("created_by")
		UpdatedByColumn             = postgres.IntegerColumn("updated_by")
		ChangesImportColumn         = postgres.StringColumn("changes_import")
		ErrorsImportColumn          = postgres.StringColumn("errors_import")
		allColumns                  = postgres.ColumnList{IDColumn, ImportedAtColumn, CreatedAtColumn, UpdatedAtColumn, RemoveAbsentEmployeesColumn, EntityIDColumn, ImportTypeIDColumn, CreatedByColumn, UpdatedByColumn, ChangesImportColumn, ErrorsImportColumn}
		mutableColumns              = postgres.ColumnList{ImportedAtColumn, CreatedAtColumn, UpdatedAtColumn, RemoveAbsentEmployeesColumn, EntityIDColumn, ImportTypeIDColumn, CreatedByColumn, UpdatedByColumn, ChangesImportColumn, ErrorsImportColumn}
	)

	return employeeImportTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                    IDColumn,
		ImportedAt:            ImportedAtColumn,
		CreatedAt:             CreatedAtColumn,
		UpdatedAt:             UpdatedAtColumn,
		RemoveAbsentEmployees: RemoveAbsentEmployeesColumn,
		EntityID:              EntityIDColumn,
		ImportTypeID:          ImportTypeIDColumn,
		CreatedBy:             CreatedByColumn,
		UpdatedBy:             UpdatedByColumn,
		ChangesImport:         ChangesImportColumn,
		ErrorsImport:          ErrorsImportColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
