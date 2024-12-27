package model

import "github.com/uptrace/bun"

type RbacPermissions struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type RbacPermissionsSlice []RbacPermissions

type DBRbacPermissions struct {
	bun.BaseModel `bun:"table:rbac_permissions,alias:rbac_permissions"`
	Id int `bun:"id"`
	Name string `bun:"name"`
}

type DBRbacPermissionsSlice []DBRbacPermissions

func (entry *RbacPermissions) ToDB() *DBRbacPermissions {
	if entry == nil {
		return nil
	}
	return &DBRbacPermissions{
		Id: entry.Id,
		Name: entry.Name,
	}
}

func (entry *DBRbacPermissions) ToWeb() *RbacPermissions {
	if entry == nil {
		return nil
	}
	return &RbacPermissions{
		Id: entry.Id,
		Name: entry.Name,
	}
}

func (data RbacPermissionsSlice) ToDB() DBRbacPermissionsSlice {
	result := make(DBRbacPermissionsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBRbacPermissionsSlice) ToWeb() RbacPermissionsSlice {
	result := make(RbacPermissionsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
