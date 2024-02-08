package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/engagerocketco/templates-api-svc/internal/repository"
	communicationtemplate "github.com/engagerocketco/templates-api-svc/internal/repository/postgres/tables/communicationtemplate"
	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel"
)

func (p *PostgresRepo) CreateCommunicationTemplate(ctx context.Context, createTemplComms *repomodel.CreateCommunicationTemplate) (*repomodel.CommunicationTemplate, error) {
	var templComms repomodel.CommunicationTemplate

	_, err := DialectGoQu().DB(p.DB).
		Insert(goqu.T(communicationtemplate.TableName)).
		Rows(createTemplComms).Returning(
		communicationtemplate.ColumnID,
		communicationtemplate.ColumnOwnerEntityID,
		communicationtemplate.ColumnHeaderLogoID,
		communicationtemplate.ColumnReminderDaysID,
		communicationtemplate.ColumnIsSendReport,
		communicationtemplate.ColumnName,
		communicationtemplate.ColumnDescription,
		communicationtemplate.ColumnTimeSendReport,
		communicationtemplate.ColumnCreatedAt,
		communicationtemplate.ColumnUpdatedAt,
		communicationtemplate.ColumnCreatedBy,
		communicationtemplate.ColumnUpdatedBy,
	).Executor().ScanStructContext(ctx, &templComms)

	if err != nil {
		return nil, fmt.Errorf("could not insert new communication template record: %w", err)
	}

	return &templComms, nil
}

func (p *PostgresRepo) GetCommunicationTemplateByID(ctx context.Context, id int) (*repomodel.CommunicationTemplate, error) {
	var templComms repomodel.CommunicationTemplate

	found, err := DialectGoQu().DB(p.DB).From(communicationtemplate.TableName).Select(
		communicationtemplate.ColumnID,
		communicationtemplate.ColumnOwnerEntityID,
		communicationtemplate.ColumnHeaderLogoID,
		communicationtemplate.ColumnReminderDaysID,
		communicationtemplate.ColumnIsSendReport,
		communicationtemplate.ColumnName,
		communicationtemplate.ColumnDescription,
		communicationtemplate.ColumnTimeSendReport,
		communicationtemplate.ColumnCreatedAt,
		communicationtemplate.ColumnUpdatedAt,
		communicationtemplate.ColumnCreatedBy,
		communicationtemplate.ColumnUpdatedBy,
	).Where(
		goqu.C(communicationtemplate.ColumnID).Eq(id),
	).ScanStructContext(ctx, &templComms)

	if err != nil {
		return nil, fmt.Errorf("failed to select communication template record: %w", err)
	}

	if !found {
		return nil, repository.ErrCommunicationTemplateNotFound
	}

	return &templComms, err
}

func (p *PostgresRepo) GetCommunicationTemplatesByEntityID(ctx context.Context, id int) ([]repomodel.CommunicationTemplate, error) {
	var templComms []repomodel.CommunicationTemplate

	err := DialectGoQu().DB(p.DB).From(communicationtemplate.TableName).Select(
		communicationtemplate.ColumnID,
		communicationtemplate.ColumnOwnerEntityID,
		communicationtemplate.ColumnHeaderLogoID,
		communicationtemplate.ColumnReminderDaysID,
		communicationtemplate.ColumnIsSendReport,
		communicationtemplate.ColumnName,
		communicationtemplate.ColumnDescription,
		communicationtemplate.ColumnTimeSendReport,
		communicationtemplate.ColumnCreatedAt,
		communicationtemplate.ColumnUpdatedAt,
		communicationtemplate.ColumnCreatedBy,
		communicationtemplate.ColumnUpdatedBy,
	).Where(
		goqu.C(communicationtemplate.ColumnOwnerEntityID).Eq(id),
	).ScanStructsContext(ctx, &templComms)

	if err != nil {
		return nil, fmt.Errorf("failed to select communication template record: %w", err)
	}

	if templComms == nil {
		return nil, repository.ErrCommunicationTemplateNotFound
	}

	return templComms, err
}

func (p *PostgresRepo) UpdateCommunicationTemplateByID(ctx context.Context, updateTemplComms *repomodel.UpdateCommunicationTemplate) (*repomodel.CommunicationTemplate, error) {
	var templComms repomodel.CommunicationTemplate
	columnsToUpdate := goqu.Record{}

	if updateTemplComms.OwnerEntityID != 0 {
		columnsToUpdate[communicationtemplate.ColumnOwnerEntityID] = updateTemplComms.OwnerEntityID
	}

	if updateTemplComms.ReminderDaysID != 0 {
		columnsToUpdate[communicationtemplate.ColumnHeaderLogoID] = updateTemplComms.HeaderLogoID
	}

	if updateTemplComms.ReminderDaysID != 0 {
		columnsToUpdate[communicationtemplate.ColumnReminderDaysID] = updateTemplComms.ReminderDaysID
	}

	if updateTemplComms.IsSendReport != nil {
		columnsToUpdate[communicationtemplate.ColumnIsSendReport] = updateTemplComms.IsSendReport
	}

	if updateTemplComms.Name != nil {
		columnsToUpdate[communicationtemplate.ColumnName] = updateTemplComms.Name
	}

	if updateTemplComms.Description != nil {
		columnsToUpdate[communicationtemplate.ColumnDescription] = updateTemplComms.Description
	}

	if updateTemplComms.TimeSendReport != nil {
		columnsToUpdate[communicationtemplate.ColumnTimeSendReport] = updateTemplComms.TimeSendReport
	}

	columnsToUpdate[communicationtemplate.ColumnUpdatedAt] = time.Now().UTC()

	found, err := DialectGoQu().DB(p.DB).
		Update(goqu.T(communicationtemplate.TableName)).
		Set(columnsToUpdate).
		Where(
			goqu.C(communicationtemplate.ColumnID).Eq(updateTemplComms.ID),
		).Returning(
		communicationtemplate.ColumnID,
		communicationtemplate.ColumnOwnerEntityID,
		communicationtemplate.ColumnHeaderLogoID,
		communicationtemplate.ColumnReminderDaysID,
		communicationtemplate.ColumnIsSendReport,
		communicationtemplate.ColumnName,
		communicationtemplate.ColumnDescription,
		communicationtemplate.ColumnTimeSendReport,
		communicationtemplate.ColumnCreatedAt,
		communicationtemplate.ColumnUpdatedAt,
		communicationtemplate.ColumnCreatedBy,
		communicationtemplate.ColumnUpdatedBy,
	).Executor().ScanStructContext(ctx, &templComms)

	if err != nil {
		return nil, fmt.Errorf("failed to update communication template record: %w", err)
	}

	if !found {
		return nil, repository.ErrCommunicationTemplateNotFound
	}

	return &templComms, nil
}

func (p *PostgresRepo) DeleteCommunicationTemplateByID(ctx context.Context, id int) error {
	var foundID int

	found, err := DialectGoQu().DB(p.DB).
		Delete(goqu.T(communicationtemplate.TableName)).
		Where(
			goqu.C(communicationtemplate.ColumnID).Eq(id),
		).
		Returning(communicationtemplate.ColumnID).
		Executor().
		ScanValContext(ctx, &foundID)

	if err != nil {
		return fmt.Errorf("failed to delete communication template record: %w", err)
	}

	if !found {
		return repository.ErrCommunicationTemplateNotFound
	}

	return nil
}
