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

type OrganizationUnit struct {
	ID          int32 `sql:"primary_key"`
	ParentID    *int32
	Name        *string
	Description *string
	Level       *int32
	Details     *string
	Emails      *string
	EntityID    *int32
	UpdatedAt   *time.Time
	CreatedAt   *time.Time
	UpdatedBy   *int32
	CreatedBy   *int32
}