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

func (r *PostgresRepo) CreateSmAttributeTriggers(ctx context.Context, input model.SmAttributeTriggers) (*repomodel.BaseIdResponse, error) {

	stmt := table.SmAttributeTriggers.
		INSERT(table.SmAttributeTriggers.MutableColumns.Except(table.SmAttributeTriggers.CreatedAt, table.SmAttributeTriggers.UpdatedAt)).
		MODEL(input).
		RETURNING(table.SmAttributeTriggers.ID)

	row := model.SmAttributeTriggers{}
	err := stmt.Query(r.JetDB, &row)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(row.ID),
	}, nil
}

func (r *PostgresRepo) CreateSmAttributeTriggersRange(ctx context.Context, inputs []model.SmAttributeTriggers) error {

	stmt := table.SmAttributeTriggers.
		INSERT(table.SmAttributeTriggers.MutableColumns.Except(table.SmAttributeTriggers.CreatedAt, table.SmAttributeTriggers.UpdatedAt)).
		MODELS(inputs)

	_, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create survey_tags: %w", err)
	}

	return nil
}

func (r *PostgresRepo) GetSmAttributeTriggers(ctx context.Context, req repomodel.GetSmAttributeTriggersRequest) ([]repomodel.GetSmAttributeTriggers, error) {

	pagination_expr := GenerateDynamicWhereClause(table.SmAttributeTriggers.AllColumns, req.PaginationRequest)
	orderby_expr := GenerateDynamicOrderByClause(table.SmAttributeTriggers.AllColumns, req.PaginationRequest)

	stmt := SELECT(table.SmAttributeTriggers.AllColumns,
		table.Attributes.AllColumns,
		table.SmProject.AllColumns,
	).
		FROM(
			table.SmAttributeTriggers.
				LEFT_JOIN(table.Attributes, table.Attributes.ID.EQ(table.SmAttributeTriggers.AttributeID)).
				LEFT_JOIN(table.SmProject, table.SmProject.ID.EQ(table.SmAttributeTriggers.ProjectID))).
		WHERE(pagination_expr).
		LIMIT(int64(req.Limit)).
		OFFSET(int64(req.Offset)).
		ORDER_BY(orderby_expr)

	response := []repomodel.GetSmAttributeTriggers{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return response, nil
}

func (r *PostgresRepo) GetSmAttributeTriggersByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetSmAttributeTriggers, error) {

	stmt := SELECT(table.SmAttributeTriggers.AllColumns,
		table.Attributes.AllColumns,
		table.SmProject.AllColumns,
	).
		FROM(
			table.SmAttributeTriggers.
				LEFT_JOIN(table.Attributes, table.Attributes.ID.EQ(table.SmAttributeTriggers.AttributeID)).
				LEFT_JOIN(table.SmProject, table.SmProject.ID.EQ(table.SmAttributeTriggers.ProjectID))).
		WHERE(table.SmAttributeTriggers.ID.EQ(Int32(input.ID)))
	response := repomodel.GetSmAttributeTriggers{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {
		return nil, fmt.Errorf("select survey_tags table: %w", err)
	}
	return &response, nil
}

func (r *PostgresRepo) CountSmAttributeTriggers(ctx context.Context) (*int, error) {

	stmt := SELECT(COUNT(table.SmAttributeTriggers.ID).AS("count")).
		FROM(table.SmAttributeTriggers)

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

func (r *PostgresRepo) UpdateSmAttributeTriggersByID(ctx context.Context, input model.SmAttributeTriggers) (*repomodel.BaseIdResponse, error) {

	stmt := table.SmAttributeTriggers.
		UPDATE(table.SmAttributeTriggers.MutableColumns.
			Except(
				table.SmAttributeTriggers.CreatedAt,
				table.SmAttributeTriggers.CreatedBy)).
		MODEL(input).
		WHERE(table.SmAttributeTriggers.ID.EQ(Int32(input.ID))).
		RETURNING(table.SmAttributeTriggers.ID)

	response := model.SmAttributeTriggers{}
	err := stmt.Query(r.JetDB, &response)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(input.ID),
	}, nil
}

func (r *PostgresRepo) DeleteSmAttributeTriggersByID(ctx context.Context, input repomodel.BaseIdRequest) error {

	stmt := table.SmAttributeTriggers.
		DELETE().
		WHERE(table.SmAttributeTriggers.ID.EQ(Int32(input.ID)))

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
