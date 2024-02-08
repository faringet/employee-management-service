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

var SmQuestionaryTranslate = newSmQuestionaryTranslateTable("public", "sm_questionary_translate", "")

type smQuestionaryTranslateTable struct {
	postgres.Table

	// Columns
	ID         postgres.ColumnInteger
	UpdatedAt  postgres.ColumnTimestamp
	CreatedAt  postgres.ColumnTimestamp
	UpdatedBy  postgres.ColumnInteger
	CreatedBy  postgres.ColumnInteger
	LanguageID postgres.ColumnInteger
	SurveyID   postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SmQuestionaryTranslateTable struct {
	smQuestionaryTranslateTable

	EXCLUDED smQuestionaryTranslateTable
}

// AS creates new SmQuestionaryTranslateTable with assigned alias
func (a SmQuestionaryTranslateTable) AS(alias string) *SmQuestionaryTranslateTable {
	return newSmQuestionaryTranslateTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SmQuestionaryTranslateTable with assigned schema name
func (a SmQuestionaryTranslateTable) FromSchema(schemaName string) *SmQuestionaryTranslateTable {
	return newSmQuestionaryTranslateTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SmQuestionaryTranslateTable with assigned table prefix
func (a SmQuestionaryTranslateTable) WithPrefix(prefix string) *SmQuestionaryTranslateTable {
	return newSmQuestionaryTranslateTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SmQuestionaryTranslateTable with assigned table suffix
func (a SmQuestionaryTranslateTable) WithSuffix(suffix string) *SmQuestionaryTranslateTable {
	return newSmQuestionaryTranslateTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSmQuestionaryTranslateTable(schemaName, tableName, alias string) *SmQuestionaryTranslateTable {
	return &SmQuestionaryTranslateTable{
		smQuestionaryTranslateTable: newSmQuestionaryTranslateTableImpl(schemaName, tableName, alias),
		EXCLUDED:                    newSmQuestionaryTranslateTableImpl("", "excluded", ""),
	}
}

func newSmQuestionaryTranslateTableImpl(schemaName, tableName, alias string) smQuestionaryTranslateTable {
	var (
		IDColumn         = postgres.IntegerColumn("id")
		UpdatedAtColumn  = postgres.TimestampColumn("updated_at")
		CreatedAtColumn  = postgres.TimestampColumn("created_at")
		UpdatedByColumn  = postgres.IntegerColumn("updated_by")
		CreatedByColumn  = postgres.IntegerColumn("created_by")
		LanguageIDColumn = postgres.IntegerColumn("language_id")
		SurveyIDColumn   = postgres.IntegerColumn("survey_id")
		allColumns       = postgres.ColumnList{IDColumn, UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, LanguageIDColumn, SurveyIDColumn}
		mutableColumns   = postgres.ColumnList{UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, LanguageIDColumn, SurveyIDColumn}
	)

	return smQuestionaryTranslateTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		UpdatedAt:  UpdatedAtColumn,
		CreatedAt:  CreatedAtColumn,
		UpdatedBy:  UpdatedByColumn,
		CreatedBy:  CreatedByColumn,
		LanguageID: LanguageIDColumn,
		SurveyID:   SurveyIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}