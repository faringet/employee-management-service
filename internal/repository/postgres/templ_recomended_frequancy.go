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

func (r *PostgresRepo) CreateTemplRecomendedFrequancy(ctx context.Context, input model.TemplRecomendedFrequancy) (*repomodel.BaseIdResponse, error) {

	stmt := table.TemplRecomendedFrequancy.
		INSERT(table.TemplRecomendedFrequancy.MutableColumns.Except(table.TemplRecomendedFrequancy.CreatedAt, table.TemplRecomendedFrequancy.UpdatedAt)).
		MODEL(input).
		RETURNING(table.TemplRecomendedFrequancy.ID)

	row := model.TemplRecomendedFrequancy{}
	err := stmt.Query(r.JetDB, &row)
	if err != nil {
		return nil, fmt.Errorf("create TemplRecomendedFrequancy: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(row.ID),
	}, nil
}

func (r *PostgresRepo) CreateTemplRecomendedFrequancyRange(ctx context.Context, inputs []model.TemplRecomendedFrequancy) error {

	stmt := table.TemplRecomendedFrequancy.
		INSERT(table.TemplRecomendedFrequancy.MutableColumns.Except(table.TemplRecomendedFrequancy.CreatedAt, table.TemplRecomendedFrequancy.UpdatedAt)).
		MODELS(inputs)

	_, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create TemplRecomendedFrequancy: %w", err)
	}

	return nil
}

func (r *PostgresRepo) GetTemplRecomendedFrequancy(ctx context.Context, req repomodel.GetTemplRecomendedFrequancyRequest) ([]model.TemplRecomendedFrequancy, error) {

	pagination_expr := GenerateDynamicWhereClause(table.TemplRecomendedFrequancy.AllColumns, req.PaginationRequest)
	orderby_expr := GenerateDynamicOrderByClause(table.TemplRecomendedFrequancy.AllColumns, req.PaginationRequest)

	stmt := SELECT(table.TemplRecomendedFrequancy.AllColumns).
		FROM(table.TemplRecomendedFrequancy).
		WHERE(pagination_expr).
		LIMIT(int64(req.Limit)).
		OFFSET(int64(req.Offset)).
		ORDER_BY(orderby_expr)

	response := []model.TemplRecomendedFrequancy{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select TemplRecomendedFrequancy table: %w", err)
	}
	return response, nil
}

func (r *PostgresRepo) GetTemplRecomendedFrequancyByID(ctx context.Context, input repomodel.BaseIdRequest) (*model.TemplRecomendedFrequancy, error) {

	stmt := SELECT(table.TemplRecomendedFrequancy.AllColumns).
		FROM(table.TemplRecomendedFrequancy).
		WHERE(table.TemplRecomendedFrequancy.ID.EQ(Int32(input.ID)))
	response := model.TemplRecomendedFrequancy{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select TemplRecomendedFrequancy table: %w", err)
	}
	return &response, nil
}

func (r *PostgresRepo) UpdateTemplRecomendedFrequancyByID(ctx context.Context, input model.TemplRecomendedFrequancy) (*repomodel.BaseIdResponse, error) {

	stmt := table.TemplRecomendedFrequancy.
		UPDATE(table.TemplRecomendedFrequancy.MutableColumns.
			Except(
				table.TemplRecomendedFrequancy.CreatedAt,
				table.TemplRecomendedFrequancy.CreatedBy)).
		MODEL(input).
		WHERE(table.TemplRecomendedFrequancy.ID.EQ(Int32(input.ID))).
		RETURNING(table.TemplRecomendedFrequancy.ID)

	response := model.TemplRecomendedFrequancy{}
	err := stmt.Query(r.JetDB, &response)
	if err != nil {
		return nil, fmt.Errorf("create TemplRecomendedFrequancy: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(input.ID),
	}, nil
}

func (r *PostgresRepo) DeleteTemplRecomendedFrequancyByID(ctx context.Context, input repomodel.BaseIdRequest) error {

	stmt := table.TemplRecomendedFrequancy.
		DELETE().
		WHERE(table.TemplRecomendedFrequancy.ID.EQ(Int32(input.ID)))

	res, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create TemplRecomendedFrequancy: %w", err)
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("create TemplRecomendedFrequancy: %w", err)
	}

	if aff == 0 {
		return fmt.Errorf("not found")
	}
	return nil
}
