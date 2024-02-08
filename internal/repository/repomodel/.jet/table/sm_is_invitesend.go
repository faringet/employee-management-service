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

var SmIsInvitesend = newSmIsInvitesendTable("public", "sm_is_invitesend", "")

type smIsInvitesendTable struct {
	postgres.Table

	// Columns
	ID                postgres.ColumnInteger
	UpdatedAt         postgres.ColumnTimestamp
	CreatedBy         postgres.ColumnInteger
	CreatedAt         postgres.ColumnTimestamp
	UpdatedBy         postgres.ColumnInteger
	SurveyRecipientID postgres.ColumnInteger
	DateTimeSend      postgres.ColumnTimestamp
	IsInvite          postgres.ColumnBool
	EmailtemplateID   postgres.ColumnInteger
	IsReminder        postgres.ColumnBool

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SmIsInvitesendTable struct {
	smIsInvitesendTable

	EXCLUDED smIsInvitesendTable
}

// AS creates new SmIsInvitesendTable with assigned alias
func (a SmIsInvitesendTable) AS(alias string) *SmIsInvitesendTable {
	return newSmIsInvitesendTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SmIsInvitesendTable with assigned schema name
func (a SmIsInvitesendTable) FromSchema(schemaName string) *SmIsInvitesendTable {
	return newSmIsInvitesendTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SmIsInvitesendTable with assigned table prefix
func (a SmIsInvitesendTable) WithPrefix(prefix string) *SmIsInvitesendTable {
	return newSmIsInvitesendTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SmIsInvitesendTable with assigned table suffix
func (a SmIsInvitesendTable) WithSuffix(suffix string) *SmIsInvitesendTable {
	return newSmIsInvitesendTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSmIsInvitesendTable(schemaName, tableName, alias string) *SmIsInvitesendTable {
	return &SmIsInvitesendTable{
		smIsInvitesendTable: newSmIsInvitesendTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newSmIsInvitesendTableImpl("", "excluded", ""),
	}
}

func newSmIsInvitesendTableImpl(schemaName, tableName, alias string) smIsInvitesendTable {
	var (
		IDColumn                = postgres.IntegerColumn("id")
		UpdatedAtColumn         = postgres.TimestampColumn("updated_at")
		CreatedByColumn         = postgres.IntegerColumn("created_by")
		CreatedAtColumn         = postgres.TimestampColumn("created_at")
		UpdatedByColumn         = postgres.IntegerColumn("updated_by")
		SurveyRecipientIDColumn = postgres.IntegerColumn("survey_recipient_id")
		DateTimeSendColumn      = postgres.TimestampColumn("dateTimeSend")
		IsInviteColumn          = postgres.BoolColumn("IsInvite")
		EmailtemplateIDColumn   = postgres.IntegerColumn("emailtemplate_id")
		IsReminderColumn        = postgres.BoolColumn("isReminder")
		allColumns              = postgres.ColumnList{IDColumn, UpdatedAtColumn, CreatedByColumn, CreatedAtColumn, UpdatedByColumn, SurveyRecipientIDColumn, DateTimeSendColumn, IsInviteColumn, EmailtemplateIDColumn, IsReminderColumn}
		mutableColumns          = postgres.ColumnList{UpdatedAtColumn, CreatedByColumn, CreatedAtColumn, UpdatedByColumn, SurveyRecipientIDColumn, DateTimeSendColumn, IsInviteColumn, EmailtemplateIDColumn, IsReminderColumn}
	)

	return smIsInvitesendTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                IDColumn,
		UpdatedAt:         UpdatedAtColumn,
		CreatedBy:         CreatedByColumn,
		CreatedAt:         CreatedAtColumn,
		UpdatedBy:         UpdatedByColumn,
		SurveyRecipientID: SurveyRecipientIDColumn,
		DateTimeSend:      DateTimeSendColumn,
		IsInvite:          IsInviteColumn,
		EmailtemplateID:   EmailtemplateIDColumn,
		IsReminder:        IsReminderColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
