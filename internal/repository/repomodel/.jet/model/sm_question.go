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

type SmQuestion struct {
	ID                  int32 `sql:"primary_key"`
	UpdatedAt           *time.Time
	CreatedAt           *time.Time
	UpdatedBy           *int32
	CreatedBy           *int32
	Name                *string
	Title               *string
	IsQuestionCondition *bool
	IsCalculateScore    *bool
	TypeQuestionID      int32
	DimensionID         int32
	SurveyID            int32
}
