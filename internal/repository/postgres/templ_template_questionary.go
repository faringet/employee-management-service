package postgres

import (
	"context"
	"fmt"

	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel"
	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"
	table "github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/table"

	_ "github.com/lib/pq"

	// dot import so that jet go code would resemble as much as native SQL
	// dot import is not mandatory
	. "github.com/go-jet/jet/v2/postgres"
)

func (r *PostgresRepo) CreateTemplTemplateQuestionary(ctx context.Context, input model.TemplTemplateQuestionary) (*repomodel.BaseIdResponse, error) {

	stmt := table.TemplTemplateQuestionary.
		INSERT(table.TemplTemplateQuestionary.MutableColumns.Except(table.TemplTemplateQuestionary.CreatedAt, table.TemplTemplateQuestionary.UpdatedAt)).
		MODEL(input).
		RETURNING(table.TemplTemplateQuestionary.ID)

	row := model.TemplTemplateQuestionary{}
	err := stmt.Query(r.JetDB, &row)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(row.ID),
	}, nil
}

func (r *PostgresRepo) CreateTemplTemplateQuestionaryRange(ctx context.Context, inputs []model.TemplTemplateQuestionary) error {

	stmt := table.TemplTemplateQuestionary.
		INSERT(table.TemplTemplateQuestionary.MutableColumns.Except(table.TemplTemplateQuestionary.CreatedAt, table.TemplTemplateQuestionary.UpdatedAt)).
		MODELS(inputs)

	_, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create survey_tags: %w", err)
	}

	return nil
}

func (r *PostgresRepo) GetTemplTemplateQuestionary(ctx context.Context, req repomodel.GetTemplTemplateQuestionaryRequest) ([]repomodel.GetTemplTemplateQuestionary, error) {

	pagination_expr := GenerateDynamicWhereClause(table.TemplTemplateQuestionary.AllColumns, req.PaginationRequest)
	orderby_expr := GenerateDynamicOrderByClause(table.TemplTemplateQuestionary.AllColumns, req.PaginationRequest)

	stmt := SELECT(
		table.TemplTemplateQuestionary.AllColumns,
		table.TemplRecomendedFrequancy.AllColumns,
		table.TemplQuestionaryTags.AllColumns,
		table.SurveyTags.AllColumns).
		FROM(
			table.TemplTemplateQuestionary.
				INNER_JOIN(table.TemplRecomendedFrequancy, table.TemplRecomendedFrequancy.ID.EQ(table.TemplTemplateQuestionary.RecomendedFrequencyID)).
				INNER_JOIN(table.TemplQuestionaryTags, table.TemplQuestionaryTags.TemplateQuestionaryID.EQ(table.TemplTemplateQuestionary.ID)).
				INNER_JOIN(table.SurveyTags, table.SurveyTags.ID.EQ(table.TemplQuestionaryTags.TagID))).
		WHERE(pagination_expr).
		LIMIT(int64(req.Limit)).
		OFFSET(int64(req.Offset)).
		ORDER_BY(orderby_expr)

	response := []repomodel.GetTemplTemplateQuestionary{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return response, nil
}

func (r *PostgresRepo) GetTemplTemplateQuestionaryByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetTemplTemplateQuestionary, error) {

	stmt := SELECT(
		table.TemplTemplateQuestionary.AllColumns,
		table.TemplRecomendedFrequancy.AllColumns,
		table.TemplQuestionaryTags.AllColumns,
		table.SurveyTags.AllColumns).
		FROM(
			table.TemplTemplateQuestionary.
				INNER_JOIN(table.TemplRecomendedFrequancy, table.TemplRecomendedFrequancy.ID.EQ(table.TemplTemplateQuestionary.RecomendedFrequencyID)).
				INNER_JOIN(table.TemplQuestionaryTags, table.TemplQuestionaryTags.TemplateQuestionaryID.EQ(table.TemplTemplateQuestionary.ID)).
				INNER_JOIN(table.SurveyTags, table.SurveyTags.ID.EQ(table.TemplQuestionaryTags.TagID))).
		WHERE(table.TemplTemplateQuestionary.ID.EQ(Int32(input.ID)))

	response := repomodel.GetTemplTemplateQuestionary{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return &response, nil
}

func (r *PostgresRepo) UpdateTemplTemplateQuestionaryByID(ctx context.Context, input model.TemplTemplateQuestionary) (*repomodel.BaseIdResponse, error) {

	stmt := table.TemplTemplateQuestionary.
		UPDATE(table.TemplTemplateQuestionary.MutableColumns.
			Except(
				table.TemplTemplateQuestionary.CreatedAt,
				table.TemplTemplateQuestionary.CreatedBy)).
		MODEL(input).
		WHERE(table.TemplTemplateQuestionary.ID.EQ(Int32(input.ID))).
		RETURNING(table.TemplTemplateQuestionary.ID)

	response := model.TemplTemplateQuestionary{}
	err := stmt.Query(r.JetDB, &response)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(input.ID),
	}, nil
}

func (r *PostgresRepo) DeleteTemplTemplateQuestionaryByID(ctx context.Context, input repomodel.BaseIdRequest) error {

	stmt := table.TemplTemplateQuestionary.
		DELETE().
		WHERE(table.TemplTemplateQuestionary.ID.EQ(Int32(input.ID)))

	res, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create survey_tags: %w", err)
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("create survey_tags: %w", err)
	}

	if aff == 0 {
		return fmt.Errorf("not found")
	}
	return nil
}
