//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type AccountRole struct {
	ID            int32 `sql:"primary_key"`
	AccessLevelID int32
	RealID        *int32
	RoleID        int32
	RegionID      *int32
}