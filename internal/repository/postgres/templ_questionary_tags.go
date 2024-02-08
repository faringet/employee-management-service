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

func (r *PostgresRepo) CreateTemplQuestionaryTags(ctx context.Context, input model.TemplQuestionaryTags) (*repomodel.BaseIdResponse, error) {

	stmt := table.TemplQuestionaryTags.
		INSERT(table.TemplQuestionaryTags.MutableColumns.Except(table.TemplQuestionaryTags.CreatedAt, table.TemplQuestionaryTags.UpdatedAt)).
		MODEL(input).
		RETURNING(table.TemplQuestionaryTags.ID)

	row := model.TemplQuestionaryTags{}
	err := stmt.Query(r.JetDB, &row)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(row.ID),
	}, nil
}

func (r *PostgresRepo) CreateTemplQuestionaryTagsRange(ctx context.Context, inputs []model.TemplQuestionaryTags) error {

	stmt := table.TemplQuestionaryTags.
		INSERT(table.TemplQuestionaryTags.MutableColumns.Except(table.TemplQuestionaryTags.CreatedAt, table.TemplQuestionaryTags.UpdatedAt)).
		MODELS(inputs)

	_, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create survey_tags: %w", err)
	}

	return nil
}

func (r *PostgresRepo) GetTemplQuestionaryTags(ctx context.Context, req repomodel.GetTemplQuestionaryTagsRequest) ([]repomodel.GetTemplQuestionaryTags, error) {

	pagination_expr := GenerateDynamicWhereClause(table.TemplQuestionaryTags.AllColumns, req.PaginationRequest)
	orderby_expr := GenerateDynamicOrderByClause(table.TemplQuestionaryTags.AllColumns, req.PaginationRequest)

	stmt := SELECT(table.TemplQuestionaryTags.AllColumns, table.SurveyTags.AllColumns).
		FROM(table.TemplQuestionaryTags.
			INNER_JOIN(table.SurveyTags, table.SurveyTags.ID.EQ(table.TemplQuestionaryTags.TagID))).
		WHERE(pagination_expr).
		LIMIT(int64(req.Limit)).
		OFFSET(int64(req.Offset)).
		ORDER_BY(orderby_expr)

	response := []repomodel.GetTemplQuestionaryTags{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return response, nil
}

func (r *PostgresRepo) GetTemplQuestionaryTagsByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetTemplQuestionaryTags, error) {

	stmt := SELECT(table.TemplQuestionaryTags.AllColumns, table.SurveyTags.AllColumns).
		FROM(table.TemplQuestionaryTags.
			INNER_JOIN(table.SurveyTags, table.SurveyTags.ID.EQ(table.TemplQuestionaryTags.TagID))).
		WHERE(table.TemplQuestionaryTags.ID.EQ(Int32(input.ID)))
	response := repomodel.GetTemplQuestionaryTags{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return &response, nil
}

func (r *PostgresRepo) UpdateTemplQuestionaryTagsByID(ctx context.Context, input model.TemplQuestionaryTags) (*repomodel.BaseIdResponse, error) {

	stmt := table.TemplQuestionaryTags.
		UPDATE(table.TemplQuestionaryTags.MutableColumns.
			Except(
				table.TemplQuestionaryTags.CreatedAt,
				table.TemplQuestionaryTags.CreatedBy)).
		MODEL(input).
		WHERE(table.TemplQuestionaryTags.ID.EQ(Int32(input.ID))).
		RETURNING(table.TemplQuestionaryTags.ID)

	response := model.TemplQuestionaryTags{}
	err := stmt.Query(r.JetDB, &response)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(input.ID),
	}, nil
}

func (r *PostgresRepo) DeleteTemplQuestionaryTagsByID(ctx context.Context, input repomodel.BaseIdRequest) error {

	stmt := table.TemplQuestionaryTags.
		DELETE().
		WHERE(table.TemplQuestionaryTags.ID.EQ(Int32(input.ID)))

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
