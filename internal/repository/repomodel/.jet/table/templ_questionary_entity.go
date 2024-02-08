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

var TemplQuestionaryEntity = newTemplQuestionaryEntityTable("public", "templ_questionary_entity", "")

type templQuestionaryEntityTable struct {
	postgres.Table

	// Columns
	ID                    postgres.ColumnInteger
	UpdatedAt             postgres.ColumnTimestamp
	CreatedAt             postgres.ColumnTimestamp
	UpdatedBy             postgres.ColumnInteger
	CreatedBy             postgres.ColumnInteger
	TemplateQuestionaryID postgres.ColumnInteger
	EntityID              postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TemplQuestionaryEntityTable struct {
	templQuestionaryEntityTable

	EXCLUDED templQuestionaryEntityTable
}

// AS creates new TemplQuestionaryEntityTable with assigned alias
func (a TemplQuestionaryEntityTable) AS(alias string) *TemplQuestionaryEntityTable {
	return newTemplQuestionaryEntityTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TemplQuestionaryEntityTable with assigned schema name
func (a TemplQuestionaryEntityTable) FromSchema(schemaName string) *TemplQuestionaryEntityTable {
	return newTemplQuestionaryEntityTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TemplQuestionaryEntityTable with assigned table prefix
func (a TemplQuestionaryEntityTable) WithPrefix(prefix string) *TemplQuestionaryEntityTable {
	return newTemplQuestionaryEntityTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TemplQuestionaryEntityTable with assigned table suffix
func (a TemplQuestionaryEntityTable) WithSuffix(suffix string) *TemplQuestionaryEntityTable {
	return newTemplQuestionaryEntityTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTemplQuestionaryEntityTable(schemaName, tableName, alias string) *TemplQuestionaryEntityTable {
	return &TemplQuestionaryEntityTable{
		templQuestionaryEntityTable: newTemplQuestionaryEntityTableImpl(schemaName, tableName, alias),
		EXCLUDED:                    newTemplQuestionaryEntityTableImpl("", "excluded", ""),
	}
}

func newTemplQuestionaryEntityTableImpl(schemaName, tableName, alias string) templQuestionaryEntityTable {
	var (
		IDColumn                    = postgres.IntegerColumn("id")
		UpdatedAtColumn             = postgres.TimestampColumn("updated_at")
		CreatedAtColumn             = postgres.TimestampColumn("created_at")
		UpdatedByColumn             = postgres.IntegerColumn("updated_by")
		CreatedByColumn             = postgres.IntegerColumn("created_by")
		TemplateQuestionaryIDColumn = postgres.IntegerColumn("template_questionary_id")
		EntityIDColumn              = postgres.IntegerColumn("entity_id")
		allColumns                  = postgres.ColumnList{IDColumn, UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, TemplateQuestionaryIDColumn, EntityIDColumn}
		mutableColumns              = postgres.ColumnList{UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, TemplateQuestionaryIDColumn, EntityIDColumn}
	)

	return templQuestionaryEntityTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                    IDColumn,
		UpdatedAt:             UpdatedAtColumn,
		CreatedAt:             CreatedAtColumn,
		UpdatedBy:             UpdatedByColumn,
		CreatedBy:             CreatedByColumn,
		TemplateQuestionaryID: TemplateQuestionaryIDColumn,
		EntityID:              EntityIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
