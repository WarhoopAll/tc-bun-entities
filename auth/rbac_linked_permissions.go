package model

import "github.com/uptrace/bun"

type RbacLinkedPermissions struct {
	Id int `json:"id,omitempty"`
	LinkedId int `json:"linkedid,omitempty"`
}

type RbacLinkedPermissionsSlice []RbacLinkedPermissions

type DBRbacLinkedPermissions struct {
	bun.BaseModel `bun:"table:rbac_linked_permissions,alias:rbac_linked_permissions"`
	Id int `bun:"id"`
	LinkedId int `bun:"linkedId"`
}

type DBRbacLinkedPermissionsSlice []DBRbacLinkedPermissions

func (entry *RbacLinkedPermissions) ToDB() *DBRbacLinkedPermissions {
	if entry == nil {
		return nil
	}
	return &DBRbacLinkedPermissions{
		Id: entry.Id,
		LinkedId: entry.LinkedId,
	}
}

func (entry *DBRbacLinkedPermissions) ToWeb() *RbacLinkedPermissions {
	if entry == nil {
		return nil
	}
	return &RbacLinkedPermissions{
		Id: entry.Id,
		LinkedId: entry.LinkedId,
	}
}

func (data RbacLinkedPermissionsSlice) ToDB() DBRbacLinkedPermissionsSlice {
	result := make(DBRbacLinkedPermissionsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBRbacLinkedPermissionsSlice) ToWeb() RbacLinkedPermissionsSlice {
	result := make(RbacLinkedPermissionsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
