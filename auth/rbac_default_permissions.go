package model

import "github.com/uptrace/bun"

type RbacDefaultPermissions struct {
	SecId int `json:"secid,omitempty"`
	PermissionId int `json:"permissionid,omitempty"`
	RealmId int `json:"realmid,omitempty"`
}

type RbacDefaultPermissionsSlice []RbacDefaultPermissions

type DBRbacDefaultPermissions struct {
	bun.BaseModel `bun:"table:rbac_default_permissions,alias:rbac_default_permissions"`
	SecId int `bun:"secId"`
	PermissionId int `bun:"permissionId"`
	RealmId int `bun:"realmId"`
}

type DBRbacDefaultPermissionsSlice []DBRbacDefaultPermissions

func (entry *RbacDefaultPermissions) ToDB() *DBRbacDefaultPermissions {
	if entry == nil {
		return nil
	}
	return &DBRbacDefaultPermissions{
		SecId: entry.SecId,
		PermissionId: entry.PermissionId,
		RealmId: entry.RealmId,
	}
}

func (entry *DBRbacDefaultPermissions) ToWeb() *RbacDefaultPermissions {
	if entry == nil {
		return nil
	}
	return &RbacDefaultPermissions{
		SecId: entry.SecId,
		PermissionId: entry.PermissionId,
		RealmId: entry.RealmId,
	}
}

func (data RbacDefaultPermissionsSlice) ToDB() DBRbacDefaultPermissionsSlice {
	result := make(DBRbacDefaultPermissionsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBRbacDefaultPermissionsSlice) ToWeb() RbacDefaultPermissionsSlice {
	result := make(RbacDefaultPermissionsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
