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

var QuestionQuestiontype = newQuestionQuestiontypeTable("public", "question_questiontype", "")

type questionQuestiontypeTable struct {
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

type QuestionQuestiontypeTable struct {
	questionQuestiontypeTable

	EXCLUDED questionQuestiontypeTable
}

// AS creates new QuestionQuestiontypeTable with assigned alias
func (a QuestionQuestiontypeTable) AS(alias string) *QuestionQuestiontypeTable {
	return newQuestionQuestiontypeTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new QuestionQuestiontypeTable with assigned schema name
func (a QuestionQuestiontypeTable) FromSchema(schemaName string) *QuestionQuestiontypeTable {
	return newQuestionQuestiontypeTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new QuestionQuestiontypeTable with assigned table prefix
func (a QuestionQuestiontypeTable) WithPrefix(prefix string) *QuestionQuestiontypeTable {
	return newQuestionQuestiontypeTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new QuestionQuestiontypeTable with assigned table suffix
func (a QuestionQuestiontypeTable) WithSuffix(suffix string) *QuestionQuestiontypeTable {
	return newQuestionQuestiontypeTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newQuestionQuestiontypeTable(schemaName, tableName, alias string) *QuestionQuestiontypeTable {
	return &QuestionQuestiontypeTable{
		questionQuestiontypeTable: newQuestionQuestiontypeTableImpl(schemaName, tableName, alias),
		EXCLUDED:                  newQuestionQuestiontypeTableImpl("", "excluded", ""),
	}
}

func newQuestionQuestiontypeTableImpl(schemaName, tableName, alias string) questionQuestiontypeTable {
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

	return questionQuestiontypeTable{
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