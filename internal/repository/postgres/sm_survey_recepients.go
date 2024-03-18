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

func (r *PostgresRepo) CreateSmSurveyRecepients(ctx context.Context, input model.SmSurveyRecepients) (*repomodel.BaseIdResponse, error) {

	stmt := table.SmSurveyRecepients.
		INSERT(table.SmSurveyRecepients.MutableColumns.Except(table.SmSurveyRecepients.CreatedAt, table.SmSurveyRecepients.UpdatedAt)).
		MODEL(input).
		RETURNING(table.SmSurveyRecepients.ID)

	row := model.SmSurveyRecepients{}
	err := stmt.Query(r.JetDB, &row)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(row.ID),
	}, nil
}

func (r *PostgresRepo) CreateSmSurveyRecepientsRange(ctx context.Context, inputs []model.SmSurveyRecepients) error {

	stmt := table.SmSurveyRecepients.
		INSERT(table.SmSurveyRecepients.MutableColumns.Except(table.SmSurveyRecepients.CreatedAt, table.SmSurveyRecepients.UpdatedAt)).
		MODELS(inputs)

	_, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create survey_tags: %w", err)
	}

	return nil
}

func (r *PostgresRepo) GetSmSurveyRecepients(ctx context.Context, req repomodel.GetSmSurveyRecepientsRequest) ([]repomodel.GetSmSurveyRecepients, error) {

	pagination_expr := GenerateDynamicWhereClause(table.SmSurveyRecepients.AllColumns, req.PaginationRequest)
	orderby_expr := GenerateDynamicOrderByClause(table.SmSurveyRecepients.AllColumns, req.PaginationRequest)

	stmt := SELECT(table.SmSurveyRecepients.AllColumns,
		table.SmSurvey.AllColumns,
		table.Employees.AllColumns,
	).
		FROM(
			table.SmSurveyRecepients.
				LEFT_JOIN(table.SmSurvey, table.SmSurvey.ID.EQ(table.SmSurveyRecepients.SurveyID)).
				LEFT_JOIN(table.Employees, table.Employees.ID.EQ(table.SmSurveyRecepients.EmployeeID))).
		WHERE(pagination_expr).
		LIMIT(int64(req.Limit)).
		OFFSET(int64(req.Offset)).
		ORDER_BY(orderby_expr)

	response := []repomodel.GetSmSurveyRecepients{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return response, nil
}

func (r *PostgresRepo) GetSmSurveyRecepientsByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetSmSurveyRecepients, error) {

	stmt := SELECT(table.SmSurveyRecepients.AllColumns,
		table.SmSurvey.AllColumns,
		table.Employees.AllColumns,
	).
		FROM(
			table.SmSurveyRecepients.
				LEFT_JOIN(table.SmSurvey, table.SmSurvey.ID.EQ(table.SmSurveyRecepients.SurveyID)).
				LEFT_JOIN(table.Employees, table.Employees.ID.EQ(table.SmSurveyRecepients.EmployeeID))).
		WHERE(table.SmSurveyRecepients.ID.EQ(Int32(input.ID)))
	response := repomodel.GetSmSurveyRecepients{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return &response, nil
}

func (r *PostgresRepo) CountSmSurveyRecepients(ctx context.Context) (*int, error) {

	stmt := SELECT(COUNT(table.SmSurveyRecepients.ID).AS("count")).
		FROM(table.SmSurveyRecepients)

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

func (r *PostgresRepo) UpdateSmSurveyRecepientsByID(ctx context.Context, input model.SmSurveyRecepients) (*repomodel.BaseIdResponse, error) {

	stmt := table.SmSurveyRecepients.
		UPDATE(table.SmSurveyRecepients.MutableColumns.
			Except(
				table.SmSurveyRecepients.CreatedAt,
				table.SmSurveyRecepients.CreatedBy)).
		MODEL(input).
		WHERE(table.SmSurveyRecepients.ID.EQ(Int32(input.ID))).
		RETURNING(table.SmSurveyRecepients.ID)

	response := model.SmSurveyRecepients{}
	err := stmt.Query(r.JetDB, &response)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(input.ID),
	}, nil
}

func (r *PostgresRepo) DeleteSmSurveyRecepientsByID(ctx context.Context, input repomodel.BaseIdRequest) error {

	stmt := table.SmSurveyRecepients.
		DELETE().
		WHERE(table.SmSurveyRecepients.ID.EQ(Int32(input.ID)))

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
