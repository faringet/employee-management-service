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

var TemplCommsFooter = newTemplCommsFooterTable("public", "templ_comms_footer", "")

type templCommsFooterTable struct {
	postgres.Table

	// Columns
	ID              postgres.ColumnInteger
	UpdatedAt       postgres.ColumnTimestamp
	CreatedAt       postgres.ColumnTimestamp
	UpdatedBy       postgres.ColumnInteger
	CreatedBy       postgres.ColumnInteger
	TemplateCommsID postgres.ColumnInteger
	LanguageID      postgres.ColumnInteger
	FooterEmail     postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TemplCommsFooterTable struct {
	templCommsFooterTable

	EXCLUDED templCommsFooterTable
}

// AS creates new TemplCommsFooterTable with assigned alias
func (a TemplCommsFooterTable) AS(alias string) *TemplCommsFooterTable {
	return newTemplCommsFooterTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TemplCommsFooterTable with assigned schema name
func (a TemplCommsFooterTable) FromSchema(schemaName string) *TemplCommsFooterTable {
	return newTemplCommsFooterTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TemplCommsFooterTable with assigned table prefix
func (a TemplCommsFooterTable) WithPrefix(prefix string) *TemplCommsFooterTable {
	return newTemplCommsFooterTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TemplCommsFooterTable with assigned table suffix
func (a TemplCommsFooterTable) WithSuffix(suffix string) *TemplCommsFooterTable {
	return newTemplCommsFooterTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTemplCommsFooterTable(schemaName, tableName, alias string) *TemplCommsFooterTable {
	return &TemplCommsFooterTable{
		templCommsFooterTable: newTemplCommsFooterTableImpl(schemaName, tableName, alias),
		EXCLUDED:              newTemplCommsFooterTableImpl("", "excluded", ""),
	}
}

func newTemplCommsFooterTableImpl(schemaName, tableName, alias string) templCommsFooterTable {
	var (
		IDColumn              = postgres.IntegerColumn("id")
		UpdatedAtColumn       = postgres.TimestampColumn("updated_at")
		CreatedAtColumn       = postgres.TimestampColumn("created_at")
		UpdatedByColumn       = postgres.IntegerColumn("updated_by")
		CreatedByColumn       = postgres.IntegerColumn("created_by")
		TemplateCommsIDColumn = postgres.IntegerColumn("template_comms_id")
		LanguageIDColumn      = postgres.IntegerColumn("language_id")
		FooterEmailColumn     = postgres.StringColumn("footer_email")
		allColumns            = postgres.ColumnList{IDColumn, UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, TemplateCommsIDColumn, LanguageIDColumn, FooterEmailColumn}
		mutableColumns        = postgres.ColumnList{UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, TemplateCommsIDColumn, LanguageIDColumn, FooterEmailColumn}
	)

	return templCommsFooterTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		UpdatedAt:       UpdatedAtColumn,
		CreatedAt:       CreatedAtColumn,
		UpdatedBy:       UpdatedByColumn,
		CreatedBy:       CreatedByColumn,
		TemplateCommsID: TemplateCommsIDColumn,
		LanguageID:      LanguageIDColumn,
		FooterEmail:     FooterEmailColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
