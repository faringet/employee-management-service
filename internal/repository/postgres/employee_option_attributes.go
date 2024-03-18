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

func (r *PostgresRepo) CreateEmployeeOptionAttributes(ctx context.Context, input model.EmployeeOptionAttributes) (*repomodel.BaseIdResponse, error) {

	stmt := table.EmployeeOptionAttributes.
		INSERT(table.EmployeeOptionAttributes.MutableColumns.Except(table.EmployeeOptionAttributes.CreatedAt, table.EmployeeOptionAttributes.UpdatedAt)).
		MODEL(input).
		RETURNING(table.EmployeeOptionAttributes.ID)

	row := model.EmployeeOptionAttributes{}
	err := stmt.Query(r.JetDB, &row)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(row.ID),
	}, nil
}

func (r *PostgresRepo) CreateEmployeeOptionAttributesRange(ctx context.Context, inputs []model.EmployeeOptionAttributes) error {

	stmt := table.EmployeeOptionAttributes.
		INSERT(table.EmployeeOptionAttributes.MutableColumns.Except(table.EmployeeOptionAttributes.CreatedAt, table.EmployeeOptionAttributes.UpdatedAt)).
		MODELS(inputs)

	_, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create survey_tags: %w", err)
	}

	return nil
}

func (r *PostgresRepo) GetEmployeeOptionAttributes(ctx context.Context, req repomodel.GetEmployeeOptionAttributesRequest) ([]repomodel.GetEmployeeOptionAttributes, error) {

	pagination_expr := GenerateDynamicWhereClause(table.EmployeeOptionAttributes.AllColumns, req.PaginationRequest)
	orderby_expr := GenerateDynamicOrderByClause(table.EmployeeOptionAttributes.AllColumns, req.PaginationRequest)

	stmt := SELECT(table.EmployeeOptionAttributes.AllColumns,
		table.Attributes.AllColumns,
		table.Employees.AllColumns,
	).
		FROM(
			table.EmployeeOptionAttributes.
				LEFT_JOIN(table.Attributes, table.Attributes.ID.EQ(table.EmployeeOptionAttributes.AttributeID)).
				LEFT_JOIN(table.Employees, table.Employees.ID.EQ(table.EmployeeOptionAttributes.EmployeeID))).
		WHERE(pagination_expr).
		LIMIT(int64(req.Limit)).
		OFFSET(int64(req.Offset)).
		ORDER_BY(orderby_expr)

	response := []repomodel.GetEmployeeOptionAttributes{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return response, nil
}

func (r *PostgresRepo) GetEmployeeOptionAttributesByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetEmployeeOptionAttributes, error) {

	stmt := SELECT(table.EmployeeOptionAttributes.AllColumns,
		table.Attributes.AllColumns,
		table.Employees.AllColumns,
	).
		FROM(
			table.EmployeeOptionAttributes.
				LEFT_JOIN(table.Attributes, table.Attributes.ID.EQ(table.EmployeeOptionAttributes.AttributeID)).
				LEFT_JOIN(table.Employees, table.Employees.ID.EQ(table.EmployeeOptionAttributes.EmployeeID))).
		WHERE(table.EmployeeOptionAttributes.ID.EQ(Int32(input.ID)))
	response := repomodel.GetEmployeeOptionAttributes{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return &response, nil
}

func (r *PostgresRepo) CountEmployeeOptionAttributes(ctx context.Context) (*int, error) {

	stmt := SELECT(COUNT(table.EmployeeOptionAttributes.ID).AS("count")).
		FROM(table.EmployeeOptionAttributes)

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

func (r *PostgresRepo) UpdateEmployeeOptionAttributesByID(ctx context.Context, input model.EmployeeOptionAttributes) (*repomodel.BaseIdResponse, error) {

	stmt := table.EmployeeOptionAttributes.
		UPDATE(table.EmployeeOptionAttributes.MutableColumns.
			Except(
				table.EmployeeOptionAttributes.CreatedAt,
				table.EmployeeOptionAttributes.CreatedBy)).
		MODEL(input).
		WHERE(table.EmployeeOptionAttributes.ID.EQ(Int32(input.ID))).
		RETURNING(table.EmployeeOptionAttributes.ID)

	response := model.EmployeeOptionAttributes{}
	err := stmt.Query(r.JetDB, &response)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(input.ID),
	}, nil
}

func (r *PostgresRepo) DeleteEmployeeOptionAttributesByID(ctx context.Context, input repomodel.BaseIdRequest) error {

	stmt := table.EmployeeOptionAttributes.
		DELETE().
		WHERE(table.EmployeeOptionAttributes.ID.EQ(Int32(input.ID)))

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
