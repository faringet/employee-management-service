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

type TemplQuestion struct {
	ID                  int32 `sql:"primary_key"`
	UpdatedAt           *time.Time
	CreatedAt           *time.Time
	UpdatedBy           *int32
	CreatedBy           *int32
	Name                *string
	Title               *string
	TemplateID          int32
	TypeQuestionID      int32
	DimensionID         *int32
	IsQuestionCondition *bool
	IsCalculateScore    *bool
	QuestionJson        *string
}
