package repomodel

import "time"

type GetCommunicationTemplatesByEntityIDRequest struct {
	ID int
}

type GetCommunicationTemplateByIDWithEntityRequest struct {
	ID int
}

type CommunicationTemplate struct {
	ID             int        `db:"id"`
	OwnerEntityID  int        `db:"owner_entity_id"`
	HeaderLogoID   int        `db:"header_logo_id"`
	ReminderDaysID int        `db:"reminder_days_id"`
	IsSendReport   *bool      `db:"is_send_report"`
	Name           *string    `db:"name"`
	Description    *string    `db:"description"`
	TimeSendReport *time.Time `db:"time_send_report"`
	CreatedAt      *time.Time `db:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at"`
	CreatedBy      *int       `db:"created_by"`
	UpdatedBy      *int       `db:"updated_by"`
}

type CreateCommunicationTemplate struct {
	OwnerEntityID  int        `db:"owner_entity_id"`
	HeaderLogoID   int        `db:"header_logo_id"`
	ReminderDaysID int        `db:"reminder_days_id"`
	IsSendReport   *bool      `db:"is_send_report"`
	Name           *string    `db:"name"`
	Description    *string    `db:"description"`
	TimeSendReport *time.Time `db:"time_send_report"`
	CreatedBy      *int       `db:"created_by"`
	UpdatedBy      *int       `db:"updated_by"`
}

type UpdateCommunicationTemplate struct {
	ID             int        `db:"id"`
	OwnerEntityID  int        `db:"owner_entity_id"`
	HeaderLogoID   int        `db:"header_logo_id"`
	ReminderDaysID int        `db:"reminder_days_id"`
	IsSendReport   *bool      `db:"is_send_report"`
	Name           *string    `db:"name"`
	Description    *string    `db:"description"`
	TimeSendReport *time.Time `db:"time_send_report"`
	UpdatedAt      *time.Time `db:"updated_at"`
	UpdatedBy      *int       `db:"updated_by"`
}
