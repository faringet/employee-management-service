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

var SmCommsEmail = newSmCommsEmailTable("public", "sm_comms_email", "")

type smCommsEmailTable struct {
	postgres.Table

	// Columns
	ID              postgres.ColumnInteger
	UpdatedAt       postgres.ColumnTimestamp
	CreatedAt       postgres.ColumnTimestamp
	UpdatedBy       postgres.ColumnInteger
	CreatedBy       postgres.ColumnInteger
	BodyEmail       postgres.ColumnString
	SubjectEmail    postgres.ColumnString
	CommsReminderID postgres.ColumnInteger
	IsForReport     postgres.ColumnBool
	LanguageID      postgres.ColumnInteger
	SurveyID        postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SmCommsEmailTable struct {
	smCommsEmailTable

	EXCLUDED smCommsEmailTable
}

// AS creates new SmCommsEmailTable with assigned alias
func (a SmCommsEmailTable) AS(alias string) *SmCommsEmailTable {
	return newSmCommsEmailTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SmCommsEmailTable with assigned schema name
func (a SmCommsEmailTable) FromSchema(schemaName string) *SmCommsEmailTable {
	return newSmCommsEmailTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SmCommsEmailTable with assigned table prefix
func (a SmCommsEmailTable) WithPrefix(prefix string) *SmCommsEmailTable {
	return newSmCommsEmailTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SmCommsEmailTable with assigned table suffix
func (a SmCommsEmailTable) WithSuffix(suffix string) *SmCommsEmailTable {
	return newSmCommsEmailTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSmCommsEmailTable(schemaName, tableName, alias string) *SmCommsEmailTable {
	return &SmCommsEmailTable{
		smCommsEmailTable: newSmCommsEmailTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newSmCommsEmailTableImpl("", "excluded", ""),
	}
}

func newSmCommsEmailTableImpl(schemaName, tableName, alias string) smCommsEmailTable {
	var (
		IDColumn              = postgres.IntegerColumn("id")
		UpdatedAtColumn       = postgres.TimestampColumn("updated_at")
		CreatedAtColumn       = postgres.TimestampColumn("created_at")
		UpdatedByColumn       = postgres.IntegerColumn("updated_by")
		CreatedByColumn       = postgres.IntegerColumn("created_by")
		BodyEmailColumn       = postgres.StringColumn("body_email")
		SubjectEmailColumn    = postgres.StringColumn("subject_email")
		CommsReminderIDColumn = postgres.IntegerColumn("comms_reminder_id")
		IsForReportColumn     = postgres.BoolColumn("is_for_report")
		LanguageIDColumn      = postgres.IntegerColumn("language_id")
		SurveyIDColumn        = postgres.IntegerColumn("survey_id")
		allColumns            = postgres.ColumnList{IDColumn, UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, BodyEmailColumn, SubjectEmailColumn, CommsReminderIDColumn, IsForReportColumn, LanguageIDColumn, SurveyIDColumn}
		mutableColumns        = postgres.ColumnList{UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, BodyEmailColumn, SubjectEmailColumn, CommsReminderIDColumn, IsForReportColumn, LanguageIDColumn, SurveyIDColumn}
	)

	return smCommsEmailTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		UpdatedAt:       UpdatedAtColumn,
		CreatedAt:       CreatedAtColumn,
		UpdatedBy:       UpdatedByColumn,
		CreatedBy:       CreatedByColumn,
		BodyEmail:       BodyEmailColumn,
		SubjectEmail:    SubjectEmailColumn,
		CommsReminderID: CommsReminderIDColumn,
		IsForReport:     IsForReportColumn,
		LanguageID:      LanguageIDColumn,
		SurveyID:        SurveyIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}