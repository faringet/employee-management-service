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

func (r *PostgresRepo) CreateSmSurveyStatus(ctx context.Context, input model.SmSurveyStatus) (*repomodel.BaseIdResponse, error) {

	stmt := table.SmSurveyStatus.
		INSERT(table.SmSurveyStatus.MutableColumns.Except(table.SmSurveyStatus.CreatedAt, table.SmSurveyStatus.UpdatedAt)).
		MODEL(input).
		RETURNING(table.SmSurveyStatus.ID)

	row := model.SmSurveyStatus{}
	err := stmt.Query(r.JetDB, &row)
	if err != nil {
		return nil, fmt.Errorf("create SmSurveyStatus: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(row.ID),
	}, nil
}

func (r *PostgresRepo) CreateSmSurveyStatusRange(ctx context.Context, inputs []model.SmSurveyStatus) error {

	stmt := table.SmSurveyStatus.
		INSERT(table.SmSurveyStatus.MutableColumns.Except(table.SmSurveyStatus.CreatedAt, table.SmSurveyStatus.UpdatedAt)).
		MODELS(inputs)

	_, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create SmSurveyStatus: %w", err)
	}

	return nil
}

func (r *PostgresRepo) CountSmSurveyStatus(ctx context.Context) (*int, error) {

	stmt := SELECT(COUNT(table.SmSurveyStatus.ID).AS("count")).
		FROM(table.SmSurveyStatus)

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

func (r *PostgresRepo) GetSmSurveyStatus(ctx context.Context, req repomodel.GetSmSurveyStatusRequest) ([]model.SmSurveyStatus, error) {

	pagination_expr := GenerateDynamicWhereClause(table.SmSurveyStatus.AllColumns, req.PaginationRequest)
	orderby_expr := GenerateDynamicOrderByClause(table.SmSurveyStatus.AllColumns, req.PaginationRequest)

	stmt := SELECT(table.SmSurveyStatus.AllColumns).
		FROM(table.SmSurveyStatus).
		WHERE(pagination_expr).
		LIMIT(int64(req.Limit)).
		OFFSET(int64(req.Offset)).
		ORDER_BY(orderby_expr)

	response := []model.SmSurveyStatus{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select SmSurveyStatus table: %w", err)
	}
	return response, nil
}

func (r *PostgresRepo) GetSmSurveyStatusByID(ctx context.Context, input repomodel.BaseIdRequest) (*model.SmSurveyStatus, error) {

	stmt := SELECT(table.SmSurveyStatus.AllColumns).
		FROM(table.SmSurveyStatus).
		WHERE(table.SmSurveyStatus.ID.EQ(Int32(input.ID)))
	response := model.SmSurveyStatus{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select SmSurveyStatus table: %w", err)
	}
	return &response, nil
}

func (r *PostgresRepo) UpdateSmSurveyStatusByID(ctx context.Context, input model.SmSurveyStatus) (*repomodel.BaseIdResponse, error) {

	stmt := table.SmSurveyStatus.
		UPDATE(table.SmSurveyStatus.MutableColumns.
			Except(
				table.SmSurveyStatus.CreatedAt,
				table.SmSurveyStatus.CreatedBy)).
		MODEL(input).
		WHERE(table.SmSurveyStatus.ID.EQ(Int32(input.ID))).
		RETURNING(table.SmSurveyStatus.ID)

	response := model.SmSurveyStatus{}
	err := stmt.Query(r.JetDB, &response)
	if err != nil {
		return nil, fmt.Errorf("create SmSurveyStatus: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(input.ID),
	}, nil
}

func (r *PostgresRepo) DeleteSmSurveyStatusByID(ctx context.Context, input repomodel.BaseIdRequest) error {

	stmt := table.SmSurveyStatus.
		DELETE().
		WHERE(table.SmSurveyStatus.ID.EQ(Int32(input.ID)))

	res, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create SmSurveyStatus: %w", err)
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("create SmSurveyStatus: %w", err)
	}

	if aff == 0 {
		return fmt.Errorf("not found")
	}
	return nil
}
