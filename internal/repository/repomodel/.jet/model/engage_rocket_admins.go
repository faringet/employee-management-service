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

type EngageRocketAdmins struct {
	ID        int32 `sql:"primary_key"`
	UserID    int32
	UpdatedAt *time.Time
	CreatedAt *time.Time
	UpdatedBy *int32
	CreatedBy *int32
}
