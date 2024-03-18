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

func (r *PostgresRepo) CreateSmSurvey(ctx context.Context, input model.SmSurvey) (*repomodel.BaseIdResponse, error) {

	stmt := table.SmSurvey.
		INSERT(table.SmSurvey.MutableColumns.Except(table.SmSurvey.CreatedAt, table.SmSurvey.UpdatedAt)).
		MODEL(input).
		RETURNING(table.SmSurvey.ID)

	row := model.SmSurvey{}
	err := stmt.Query(r.JetDB, &row)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(row.ID),
	}, nil
}

func (r *PostgresRepo) CreateSmSurveyRange(ctx context.Context, inputs []model.SmSurvey) error {

	stmt := table.SmSurvey.
		INSERT(table.SmSurvey.MutableColumns.Except(table.SmSurvey.CreatedAt, table.SmSurvey.UpdatedAt)).
		MODELS(inputs)

	_, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create survey_tags: %w", err)
	}

	return nil
}

func (r *PostgresRepo) GetSmSurvey(ctx context.Context, req repomodel.GetSmSurveyRequest) ([]repomodel.GetSmSurvey, error) {

	pagination_expr := GenerateDynamicWhereClause(table.SmSurvey.AllColumns, req.PaginationRequest)
	orderby_expr := GenerateDynamicOrderByClause(table.SmSurvey.AllColumns, req.PaginationRequest)

	stmt := SELECT(table.SmSurvey.AllColumns,
		table.Attributes.AllColumns,
		table.SmProject.AllColumns,
		table.SmSurveyStatus.AllColumns,
	).
		FROM(
			table.SmSurvey.
				LEFT_JOIN(table.Attributes, table.Attributes.ID.EQ(table.SmSurvey.DateAttributeMilestoneID)).
				LEFT_JOIN(table.SmProject, table.SmProject.ID.EQ(table.SmSurvey.ProjectID)).
				LEFT_JOIN(table.SmSurveyStatus, table.SmSurveyStatus.ID.EQ(table.SmSurvey.StatusID))).
		WHERE(pagination_expr).
		LIMIT(int64(req.Limit)).
		OFFSET(int64(req.Offset)).
		ORDER_BY(orderby_expr)

	response := []repomodel.GetSmSurvey{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return response, nil
}

func (r *PostgresRepo) GetSmSurveyByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetSmSurvey, error) {

	stmt := SELECT(table.SmSurvey.AllColumns,
		table.Attributes.AllColumns,
		table.SmProject.AllColumns,
		table.SmSurveyStatus.AllColumns,
	).
		FROM(
			table.SmSurvey.
				LEFT_JOIN(table.Attributes, table.Attributes.ID.EQ(table.SmSurvey.DateAttributeMilestoneID)).
				LEFT_JOIN(table.SmProject, table.SmProject.ID.EQ(table.SmSurvey.ProjectID)).
				LEFT_JOIN(table.SmSurveyStatus, table.SmSurveyStatus.ID.EQ(table.SmSurvey.StatusID))).
		WHERE(table.SmSurvey.ID.EQ(Int32(input.ID)))
	response := repomodel.GetSmSurvey{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return &response, nil
}

func (r *PostgresRepo) CountSmSurvey(ctx context.Context) (*int, error) {

	stmt := SELECT(COUNT(table.SmSurvey.ID).AS("count")).
		FROM(table.SmSurvey)

	count := struct {
		Count int64 `json:"count"`
	}{}

	err := stmt.Query(r.JetDB, &count)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	response := int(count.Count)

	return &response, nil
}

func (r *PostgresRepo) UpdateSmSurveyByID(ctx context.Context, input model.SmSurvey) (*repomodel.BaseIdResponse, error) {

	stmt := table.SmSurvey.
		UPDATE(table.SmSurvey.MutableColumns.
			Except(
				table.SmSurvey.CreatedAt,
				table.SmSurvey.CreatedBy)).
		MODEL(input).
		WHERE(table.SmSurvey.ID.EQ(Int32(input.ID))).
		RETURNING(table.SmSurvey.ID)

	response := model.SmSurvey{}
	err := stmt.Query(r.JetDB, &response)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(input.ID),
	}, nil
}

func (r *PostgresRepo) DeleteSmSurveyByID(ctx context.Context, input repomodel.BaseIdRequest) error {

	stmt := table.SmSurvey.
		DELETE().
		WHERE(table.SmSurvey.ID.EQ(Int32(input.ID)))

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
