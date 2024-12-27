package model

import "github.com/uptrace/bun"

type RbacAccountPermissions struct {
	AccountId int `json:"accountid,omitempty"`
	PermissionId int `json:"permissionid,omitempty"`
	Granted bool `json:"granted,omitempty"`
	RealmId int `json:"realmid,omitempty"`
}

type RbacAccountPermissionsSlice []RbacAccountPermissions

type DBRbacAccountPermissions struct {
	bun.BaseModel `bun:"table:rbac_account_permissions,alias:rbac_account_permissions"`
	AccountId int `bun:"accountId"`
	PermissionId int `bun:"permissionId"`
	Granted bool `bun:"granted"`
	RealmId int `bun:"realmId"`
}

type DBRbacAccountPermissionsSlice []DBRbacAccountPermissions

func (entry *RbacAccountPermissions) ToDB() *DBRbacAccountPermissions {
	if entry == nil {
		return nil
	}
	return &DBRbacAccountPermissions{
		AccountId: entry.AccountId,
		PermissionId: entry.PermissionId,
		Granted: entry.Granted,
		RealmId: entry.RealmId,
	}
}

func (entry *DBRbacAccountPermissions) ToWeb() *RbacAccountPermissions {
	if entry == nil {
		return nil
	}
	return &RbacAccountPermissions{
		AccountId: entry.AccountId,
		PermissionId: entry.PermissionId,
		Granted: entry.Granted,
		RealmId: entry.RealmId,
	}
}

func (data RbacAccountPermissionsSlice) ToDB() DBRbacAccountPermissionsSlice {
	result := make(DBRbacAccountPermissionsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBRbacAccountPermissionsSlice) ToWeb() RbacAccountPermissionsSlice {
	result := make(RbacAccountPermissionsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
