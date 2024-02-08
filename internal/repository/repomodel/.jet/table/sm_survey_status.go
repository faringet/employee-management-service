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

var SmSurveyStatus = newSmSurveyStatusTable("public", "sm_survey_status", "")

type smSurveyStatusTable struct {
	postgres.Table

	// Columns
	ID              postgres.ColumnInteger
	Name            postgres.ColumnString
	Code            postgres.ColumnString
	Description     postgres.ColumnString
	UpdatedAt       postgres.ColumnTimestamp
	CreatedAt       postgres.ColumnTimestamp
	UpdatedBy       postgres.ColumnInteger
	CreatedBy       postgres.ColumnInteger
	QueueNumber     postgres.ColumnInteger
	IconColor       postgres.ColumnString
	IdCustomSvgIcon postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SmSurveyStatusTable struct {
	smSurveyStatusTable

	EXCLUDED smSurveyStatusTable
}

// AS creates new SmSurveyStatusTable with assigned alias
func (a SmSurveyStatusTable) AS(alias string) *SmSurveyStatusTable {
	return newSmSurveyStatusTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SmSurveyStatusTable with assigned schema name
func (a SmSurveyStatusTable) FromSchema(schemaName string) *SmSurveyStatusTable {
	return newSmSurveyStatusTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SmSurveyStatusTable with assigned table prefix
func (a SmSurveyStatusTable) WithPrefix(prefix string) *SmSurveyStatusTable {
	return newSmSurveyStatusTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SmSurveyStatusTable with assigned table suffix
func (a SmSurveyStatusTable) WithSuffix(suffix string) *SmSurveyStatusTable {
	return newSmSurveyStatusTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSmSurveyStatusTable(schemaName, tableName, alias string) *SmSurveyStatusTable {
	return &SmSurveyStatusTable{
		smSurveyStatusTable: newSmSurveyStatusTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newSmSurveyStatusTableImpl("", "excluded", ""),
	}
}

func newSmSurveyStatusTableImpl(schemaName, tableName, alias string) smSurveyStatusTable {
	var (
		IDColumn              = postgres.IntegerColumn("id")
		NameColumn            = postgres.StringColumn("name")
		CodeColumn            = postgres.StringColumn("code")
		DescriptionColumn     = postgres.StringColumn("description")
		UpdatedAtColumn       = postgres.TimestampColumn("updated_at")
		CreatedAtColumn       = postgres.TimestampColumn("created_at")
		UpdatedByColumn       = postgres.IntegerColumn("updated_by")
		CreatedByColumn       = postgres.IntegerColumn("created_by")
		QueueNumberColumn     = postgres.IntegerColumn("queueNumber")
		IconColorColumn       = postgres.StringColumn("iconColor")
		IdCustomSvgIconColumn = postgres.IntegerColumn("idCustomSvgIcon")
		allColumns            = postgres.ColumnList{IDColumn, NameColumn, CodeColumn, DescriptionColumn, UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, QueueNumberColumn, IconColorColumn, IdCustomSvgIconColumn}
		mutableColumns        = postgres.ColumnList{NameColumn, CodeColumn, DescriptionColumn, UpdatedAtColumn, CreatedAtColumn, UpdatedByColumn, CreatedByColumn, QueueNumberColumn, IconColorColumn, IdCustomSvgIconColumn}
	)

	return smSurveyStatusTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		Name:            NameColumn,
		Code:            CodeColumn,
		Description:     DescriptionColumn,
		UpdatedAt:       UpdatedAtColumn,
		CreatedAt:       CreatedAtColumn,
		UpdatedBy:       UpdatedByColumn,
		CreatedBy:       CreatedByColumn,
		QueueNumber:     QueueNumberColumn,
		IconColor:       IconColorColumn,
		IdCustomSvgIcon: IdCustomSvgIconColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
