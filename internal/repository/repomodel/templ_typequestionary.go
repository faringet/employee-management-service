package repomodel

import "time"

type TemplTypequestionary struct {
	ID              int        `db:"id"`
	Name            *string    `db:"name"`
	Code            *string    `db:"code"`
	Description     *string    `db:"description"`
	UpdatedAt       *time.Time `db:"updated_at"`
	CreatedAt       *time.Time `db:"created_at"`
	UpdatedBy       *int       `db:"updated_by"`
	CreatedBy       *int       `db:"created_by"`
	QueueNumber     *int       `db:"queueNumber"`
	IconColor       *string    `db:"iconColor"`
	IDCustomSvgIcon *int       `db:"idCustomSvgIcon"`
}
