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

type DictionaryLanguage struct {
	ID              int32 `sql:"primary_key"`
	Name            *string
	Code            *string
	Description     *string
	UpdatedAt       *time.Time
	CreatedAt       *time.Time
	UpdatedBy       *int32
	CreatedBy       *int32
	QueueNumber     *int32
	IconColor       *string
	IdCustomSvgIcon *int32
	IsDefault       *bool
	Iso             *string
	Active          *bool
}
