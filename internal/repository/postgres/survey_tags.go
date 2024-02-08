package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel"
	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"
	table "github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/table"

	_ "github.com/lib/pq"

	// dot import so that jet go code would resemble as much as native SQL
	// dot import is not mandatory
	. "github.com/go-jet/jet/v2/postgres"
)

func (r *PostgresRepo) CreateSurveyTags(ctx context.Context, input model.SurveyTags) (*repomodel.BaseIdResponse, error) {

	stmt := table.SurveyTags.
		INSERT(table.SurveyTags.MutableColumns.Except(table.SurveyTags.CreatedAt, table.SurveyTags.UpdatedAt)).
		MODEL(input).
		RETURNING(table.SurveyTags.ID)

	row := model.SurveyTags{}
	err := stmt.Query(r.JetDB, &row)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(row.ID),
	}, nil
}

func (r *PostgresRepo) CreateSurveyTagsRange(ctx context.Context, inputs []model.SurveyTags) error {

	stmt := table.SurveyTags.
		INSERT(table.SurveyTags.MutableColumns.Except(table.SurveyTags.CreatedAt, table.SurveyTags.UpdatedAt)).
		MODELS(inputs)

	_, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create survey_tags: %w", err)
	}

	return nil
}

func (r *PostgresRepo) GetSurveyTags(ctx context.Context) ([]model.SurveyTags, error) {

	stmt := SELECT(table.SurveyTags.AllColumns).
		FROM(table.SurveyTags)

	response := []model.SurveyTags{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return response, nil
}

func (r *PostgresRepo) GetSurveyTagsPagination(ctx context.Context, req repomodel.GetSurveyTagsRequest) ([]model.SurveyTags, error) {

	pagination_expr := GenerateDynamicWhereClause(table.SurveyTags.AllColumns, req.PaginationRequest)
	orderby_expr := GenerateDynamicOrderByClause(table.SurveyTags.AllColumns, req.PaginationRequest)

	stmt := SELECT(table.SurveyTags.AllColumns).
		FROM(table.SurveyTags).
		WHERE(pagination_expr).
		LIMIT(int64(req.Limit)).
		OFFSET(int64(req.Offset)).
		ORDER_BY(orderby_expr)

	response := []model.SurveyTags{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return response, nil
}

func (r *PostgresRepo) CountSurveyTags(ctx context.Context) (*int, error) {

	stmt := SELECT(COUNT(table.SurveyTags.ID).AS("count")).
		FROM(table.SurveyTags)

	log.Println(stmt.DebugSql())

	response := repomodel.CountResponse{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}

	count := int(response.Count)
	return &count, nil
}

func (r *PostgresRepo) GetSurveyTagsByID(ctx context.Context, input repomodel.BaseIdRequest) (*model.SurveyTags, error) {

	stmt := SELECT(table.SurveyTags.AllColumns).
		FROM(table.SurveyTags).
		WHERE(table.SurveyTags.ID.EQ(Int32(input.ID)))

	response := model.SurveyTags{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return &response, nil
}

func (r *PostgresRepo) UpdateSurveyTagsByID(ctx context.Context, input model.SurveyTags) (*repomodel.BaseIdResponse, error) {

	stmt := table.SurveyTags.
		UPDATE(table.SurveyTags.MutableColumns.
			Except(
				table.SurveyTags.CreatedAt,
				table.SurveyTags.CreatedBy)).
		MODEL(input).
		WHERE(table.SurveyTags.ID.EQ(Int32(input.ID))).
		RETURNING(table.SurveyTags.ID)

	response := model.SurveyTags{}
	err := stmt.Query(r.JetDB, &response)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(input.ID),
	}, nil
}

func (r *PostgresRepo) DeleteSurveyTagsByID(ctx context.Context, input repomodel.BaseIdRequest) error {

	stmt := table.SurveyTags.
		DELETE().
		WHERE(table.SurveyTags.ID.EQ(Int32(input.ID)))

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
