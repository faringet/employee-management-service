//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type TemplTemplateQuestionary struct {
	ID                    int32 `sql:"primary_key"`
	UpdatedAt             *time.Time
	CreatedAt             *time.Time
	UpdatedBy             *int32
	CreatedBy             *int32
	Name                  *string
	Description           *string
	Estimation            *string
	RecomendedFrequencyID *int32
	IsTemplateER          *bool
	TypeQuestionaryID     *int32
	Survey                *string
	OwnerEntityID         *int32
	LogoID                *int32
}
