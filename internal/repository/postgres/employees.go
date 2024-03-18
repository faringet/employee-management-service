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

func (r *PostgresRepo) CreateEmployees(ctx context.Context, input model.Employees) (*repomodel.BaseIdResponse, error) {

	stmt := table.Employees.
		INSERT(table.Employees.MutableColumns.Except(table.Employees.CreatedAt, table.Employees.UpdatedAt)).
		MODEL(input).
		RETURNING(table.Employees.ID)

	row := model.Employees{}
	err := stmt.Query(r.JetDB, &row)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(row.ID),
	}, nil
}

func (r *PostgresRepo) CreateEmployeesRange(ctx context.Context, inputs []model.Employees) error {

	stmt := table.Employees.
		INSERT(table.Employees.MutableColumns.Except(table.Employees.CreatedAt, table.Employees.UpdatedAt)).
		MODELS(inputs)

	_, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create survey_tags: %w", err)
	}

	return nil
}

func (r *PostgresRepo) GetEmployeesByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetEmployees, error) {

	stmt := SELECT(table.Employees.AllColumns,
		table.SmSurveyRecepients.AllColumns,
		table.EmployeeOptionAttributes.AllColumns,
	).
		FROM(
			table.Employees.
				LEFT_JOIN(table.SmSurveyRecepients, table.Employees.ID.EQ(table.SmSurveyRecepients.EmployeeID)).
				LEFT_JOIN(table.EmployeeOptionAttributes, table.Employees.ID.EQ(table.EmployeeOptionAttributes.EmployeeID))).
		WHERE(table.Employees.ID.EQ(Int32(input.ID)))

	response := repomodel.GetEmployees{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {

		return nil, fmt.Errorf(": % w", err)

	}

	return &response, nil

}

func (r *PostgresRepo) GetEmployees(ctx context.Context, req repomodel.GetEmployeesRequest) ([]repomodel.GetEmployees, error) {

	pagination_expr := GenerateDynamicWhereClause(table.Employees.AllColumns, req.PaginationRequest)
	orderby_expr := GenerateDynamicOrderByClause(table.Employees.AllColumns, req.PaginationRequest)

	stmt := SELECT(table.Employees.AllColumns,
		table.SmSurveyRecepients.AllColumns,
		table.EmployeeOptionAttributes.AllColumns,
	).
		FROM(
			table.Employees.
				LEFT_JOIN(table.SmSurveyRecepients, table.Employees.ID.EQ(table.SmSurveyRecepients.EmployeeID)).
				LEFT_JOIN(table.EmployeeOptionAttributes, table.Employees.ID.EQ(table.EmployeeOptionAttributes.EmployeeID))).
		WHERE(pagination_expr).
		LIMIT(int64(req.Limit)).
		OFFSET(int64(req.Offset)).
		ORDER_BY(orderby_expr)

	response := []repomodel.GetEmployees{}
	err := stmt.Query(r.JetDB, &response)

	log.Println(response)

	if err != nil {
		return nil, fmt.Errorf(": % w", err)
	}

	return response, nil

}

func (r *PostgresRepo) UpdateEmployeesByID(ctx context.Context, input model.Employees) (*repomodel.BaseIdResponse, error) {

	stmt := table.Employees.
		UPDATE(table.Employees.MutableColumns.
			Except(
				table.Employees.CreatedAt,
				table.Employees.CreatedBy)).
		MODEL(input).
		WHERE(table.Employees.ID.EQ(Int32(input.ID))).
		RETURNING(table.Employees.ID)

	response := model.Employees{}
	err := stmt.Query(r.JetDB, &response)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(input.ID),
	}, nil
}

func (r *PostgresRepo) CountEmployees(ctx context.Context) (*int, error) {

	stmt := SELECT(COUNT(table.Employees.ID).AS("count")).
		FROM(table.Employees)

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

func (r *PostgresRepo) DeleteEmployeesByID(ctx context.Context, input repomodel.BaseIdRequest) error {

	stmt := table.Employees.
		DELETE().
		WHERE(table.Employees.ID.EQ(Int32(input.ID)))

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
