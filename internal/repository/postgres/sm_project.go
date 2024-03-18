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

func (r *PostgresRepo) CreateSmProject(ctx context.Context, input model.SmProject) (*repomodel.BaseIdResponse, error) {

	stmt := table.SmProject.
		INSERT(table.SmProject.MutableColumns.Except(table.SmProject.CreatedAt, table.SmProject.UpdatedAt)).
		MODEL(input).
		RETURNING(table.SmProject.ID)

	row := model.SmProject{}
	err := stmt.Query(r.JetDB, &row)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(row.ID),
	}, nil
}

func (r *PostgresRepo) CreateSmProjectRange(ctx context.Context, inputs []model.SmProject) error {

	stmt := table.SmProject.
		INSERT(table.SmProject.MutableColumns.Except(table.SmProject.CreatedAt, table.SmProject.UpdatedAt)).
		MODELS(inputs)

	_, err := stmt.Exec(r.JetDB)
	if err != nil {
		return fmt.Errorf("create survey_tags: %w", err)
	}

	return nil
}

func (r *PostgresRepo) GetSmProjectByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetSmProject, error) {

	stmt := SELECT(table.SmProject.AllColumns,
		table.Attributes.AllColumns,
		table.SmProjectType.AllColumns,
		table.SmSurvey.AllColumns,
		table.SmAttributeTriggers.AllColumns,
	).
		FROM(
			table.SmProject.
				LEFT_JOIN(table.Attributes, table.Attributes.ID.EQ(table.SmProject.DateAttributeMilestoneID)).
				LEFT_JOIN(table.SmProjectType, table.SmProjectType.ID.EQ(table.SmProject.ProjecttypeID)).
				LEFT_JOIN(table.SmSurvey, table.SmProject.ID.EQ(table.SmSurvey.ProjectID)).
				LEFT_JOIN(table.SmAttributeTriggers, table.SmProject.ID.EQ(table.SmAttributeTriggers.ProjectID))).
		WHERE(table.SmProject.ID.EQ(Int32(input.ID)))

	response := repomodel.GetSmProject{}
	err := stmt.Query(r.JetDB, &response)

	if err != nil {

		return nil, fmt.Errorf(": % w", err)

	}

	return &response, nil

}

func (r *PostgresRepo) GetSmProject(ctx context.Context, req repomodel.GetSmProjectRequest) ([]repomodel.GetSmProject, error) {

	pagination_expr := GenerateDynamicWhereClause(table.SmProject.AllColumns, req.PaginationRequest)
	orderby_expr := GenerateDynamicOrderByClause(table.SmProject.AllColumns, req.PaginationRequest)

	stmt := SELECT(table.SmProject.AllColumns,
		table.Attributes.AllColumns,
		table.SmProjectType.AllColumns,
		table.SmSurvey.AllColumns,
		table.SmAttributeTriggers.AllColumns,
	).
		FROM(
			table.SmProject.
				LEFT_JOIN(table.Attributes, table.Attributes.ID.EQ(table.SmProject.DateAttributeMilestoneID)).
				LEFT_JOIN(table.SmProjectType, table.SmProjectType.ID.EQ(table.SmProject.ProjecttypeID)).
				LEFT_JOIN(table.SmSurvey, table.SmProject.ID.EQ(table.SmSurvey.ProjectID)).
				LEFT_JOIN(table.SmAttributeTriggers, table.SmProject.ID.EQ(table.SmAttributeTriggers.ProjectID))).
		WHERE(pagination_expr).
		LIMIT(int64(req.Limit)).
		OFFSET(int64(req.Offset)).
		ORDER_BY(orderby_expr)

	response := []repomodel.GetSmProject{}
	err := stmt.Query(r.JetDB, &response)

	log.Println(response)

	if err != nil {
		return nil, fmt.Errorf(": % w", err)
	}

	return response, nil

}

func (r *PostgresRepo) UpdateSmProjectByID(ctx context.Context, input model.SmProject) (*repomodel.BaseIdResponse, error) {

	stmt := table.SmProject.
		UPDATE(table.SmProject.MutableColumns.
			Except(
				table.SmProject.CreatedAt,
				table.SmProject.CreatedBy)).
		MODEL(input).
		WHERE(table.SmProject.ID.EQ(Int32(input.ID))).
		RETURNING(table.SmProject.ID)

	response := model.SmProject{}
	err := stmt.Query(r.JetDB, &response)
	if err != nil {
		return nil, fmt.Errorf("create survey_tags: %w", err)
	}

	return &repomodel.BaseIdResponse{
		ID: int(input.ID),
	}, nil
}

func (r *PostgresRepo) CountSmProject(ctx context.Context) (*int, error) {

	stmt := SELECT(COUNT(table.SmProject.ID).AS("count")).
		FROM(table.SmProject)

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

func (r *PostgresRepo) DeleteSmProjectByID(ctx context.Context, input repomodel.BaseIdRequest) error {

	stmt := table.SmProject.
		DELETE().
		WHERE(table.SmProject.ID.EQ(Int32(input.ID)))

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
