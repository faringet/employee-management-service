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

func (r *PostgresRepo) CreateAttributes(ctx context.Context, input model.Attributes) (*repomodel.BaseIdResponse, error) {

	stmt := table.Attributes.
		INSERT(table.Attributes.MutableColumns.Except(table.Attributes.CreatedAt, table.Attributes.UpdatedAt)).
		MODEL(input).
		RETURNING(table.Attributes.ID)

	row := model.Attributes{}
	err := stmt.Query(r.JetDB, &row)
	if err != nil {
		return nil, fmt.Errorf("create Attributes: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(row.ID),
	}, nil
}

func (r *PostgresRepo) CreateAttributesRange(ctx context.Context, inputs []model.Attributes) error {

	stmt := table.Attributes.
		INSERT(table.Attributes.MutableColumns.Except(table.Attributes.CreatedAt, table.Attributes.UpdatedAt)).
		MODELS(inputs)

	_, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create Attributes: %w", err)
	}

	return nil
}

func (r *PostgresRepo) CountAttributes(ctx context.Context) (*int, error) {

	stmt := SELECT(COUNT(table.Attributes.ID).AS("count")).
		FROM(table.Attributes)

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

func (r *PostgresRepo) GetAttributes(ctx context.Context, req repomodel.GetAttributesRequest) ([]model.Attributes, error) {

	pagination_expr := GenerateDynamicWhereClause(table.Attributes.AllColumns, req.PaginationRequest)
	orderby_expr := GenerateDynamicOrderByClause(table.Attributes.AllColumns, req.PaginationRequest)

	stmt := SELECT(table.Attributes.AllColumns).
		FROM(table.Attributes).
		WHERE(pagination_expr).
		LIMIT(int64(req.Limit)).
		OFFSET(int64(req.Offset)).
		ORDER_BY(orderby_expr)

	response := []model.Attributes{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select Attributes table: %w", err)
	}
	return response, nil
}

func (r *PostgresRepo) GetAttributesByID(ctx context.Context, input repomodel.BaseIdRequest) (*model.Attributes, error) {

	stmt := SELECT(table.Attributes.AllColumns).
		FROM(table.Attributes).
		WHERE(table.Attributes.ID.EQ(Int32(input.ID)))
	response := model.Attributes{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select Attributes table: %w", err)
	}
	return &response, nil
}

func (r *PostgresRepo) UpdateAttributesByID(ctx context.Context, input model.Attributes) (*repomodel.BaseIdResponse, error) {

	stmt := table.Attributes.
		UPDATE(table.Attributes.MutableColumns.
			Except(
				table.Attributes.CreatedAt,
				table.Attributes.CreatedBy)).
		MODEL(input).
		WHERE(table.Attributes.ID.EQ(Int32(input.ID))).
		RETURNING(table.Attributes.ID)

	response := model.Attributes{}
	err := stmt.Query(r.JetDB, &response)
	if err != nil {
		return nil, fmt.Errorf("create Attributes: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(input.ID),
	}, nil
}

func (r *PostgresRepo) DeleteAttributesByID(ctx context.Context, input repomodel.BaseIdRequest) error {

	stmt := table.Attributes.
		DELETE().
		WHERE(table.Attributes.ID.EQ(Int32(input.ID)))

	res, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create Attributes: %w", err)
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("create Attributes: %w", err)
	}

	if aff == 0 {
		return fmt.Errorf("not found")
	}
	return nil
}
