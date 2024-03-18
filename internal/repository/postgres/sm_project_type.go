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

func (r *PostgresRepo) CreateSmProjectType(ctx context.Context, input model.SmProjectType) (*repomodel.BaseIdResponse, error) {

	stmt := table.SmProjectType.
		INSERT(table.SmProjectType.MutableColumns.Except(table.SmProjectType.CreatedAt, table.SmProjectType.UpdatedAt)).
		MODEL(input).
		RETURNING(table.SmProjectType.ID)

	row := model.SmProjectType{}
	err := stmt.Query(r.JetDB, &row)
	if err != nil {
		return nil, fmt.Errorf("create SmProjectType: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(row.ID),
	}, nil
}

func (r *PostgresRepo) CreateSmProjectTypeRange(ctx context.Context, inputs []model.SmProjectType) error {

	stmt := table.SmProjectType.
		INSERT(table.SmProjectType.MutableColumns.Except(table.SmProjectType.CreatedAt, table.SmProjectType.UpdatedAt)).
		MODELS(inputs)

	_, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create SmProjectType: %w", err)
	}

	return nil
}

func (r *PostgresRepo) CountSmProjectType(ctx context.Context) (*int, error) {

	stmt := SELECT(COUNT(table.SmProjectType.ID).AS("count")).
		FROM(table.SmProjectType)

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

func (r *PostgresRepo) GetSmProjectType(ctx context.Context, req repomodel.GetSmProjectTypeRequest) ([]model.SmProjectType, error) {

	pagination_expr := GenerateDynamicWhereClause(table.SmProjectType.AllColumns, req.PaginationRequest)
	orderby_expr := GenerateDynamicOrderByClause(table.SmProjectType.AllColumns, req.PaginationRequest)

	stmt := SELECT(table.SmProjectType.AllColumns).
		FROM(table.SmProjectType).
		WHERE(pagination_expr).
		LIMIT(int64(req.Limit)).
		OFFSET(int64(req.Offset)).
		ORDER_BY(orderby_expr)

	response := []model.SmProjectType{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select SmProjectType table: %w", err)
	}
	return response, nil
}

func (r *PostgresRepo) GetSmProjectTypeByID(ctx context.Context, input repomodel.BaseIdRequest) (*model.SmProjectType, error) {

	stmt := SELECT(table.SmProjectType.AllColumns).
		FROM(table.SmProjectType).
		WHERE(table.SmProjectType.ID.EQ(Int32(input.ID)))
	response := model.SmProjectType{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select SmProjectType table: %w", err)
	}
	return &response, nil
}

func (r *PostgresRepo) UpdateSmProjectTypeByID(ctx context.Context, input model.SmProjectType) (*repomodel.BaseIdResponse, error) {

	stmt := table.SmProjectType.
		UPDATE(table.SmProjectType.MutableColumns.
			Except(
				table.SmProjectType.CreatedAt,
				table.SmProjectType.CreatedBy)).
		MODEL(input).
		WHERE(table.SmProjectType.ID.EQ(Int32(input.ID))).
		RETURNING(table.SmProjectType.ID)

	response := model.SmProjectType{}
	err := stmt.Query(r.JetDB, &response)
	if err != nil {
		return nil, fmt.Errorf("create SmProjectType: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(input.ID),
	}, nil
}

func (r *PostgresRepo) DeleteSmProjectTypeByID(ctx context.Context, input repomodel.BaseIdRequest) error {

	stmt := table.SmProjectType.
		DELETE().
		WHERE(table.SmProjectType.ID.EQ(Int32(input.ID)))

	res, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create SmProjectType: %w", err)
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("create SmProjectType: %w", err)
	}

	if aff == 0 {
		return fmt.Errorf("not found")
	}
	return nil
}
