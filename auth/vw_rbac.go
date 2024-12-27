package model

import "github.com/uptrace/bun"

type VwRbac struct {
	Permission ID int `json:"permission id,omitempty"`
	Permission Group int `json:"permission group,omitempty"`
	Security Level string `json:"security level,omitempty"`
	Permission string `json:"permission,omitempty"`
}

type VwRbacSlice []VwRbac

type DBVwRbac struct {
	bun.BaseModel `bun:"table:vw_rbac,alias:vw_rbac"`
	Permission ID int `bun:"Permission ID"`
	Permission Group int `bun:"Permission Group"`
	Security Level string `bun:"Security Level"`
	Permission string `bun:"Permission"`
}

type DBVwRbacSlice []DBVwRbac

func (entry *VwRbac) ToDB() *DBVwRbac {
	if entry == nil {
		return nil
	}
	return &DBVwRbac{
		Permission ID: entry.Permission ID,
		Permission Group: entry.Permission Group,
		Security Level: entry.Security Level,
		Permission: entry.Permission,
	}
}

func (entry *DBVwRbac) ToWeb() *VwRbac {
	if entry == nil {
		return nil
	}
	return &VwRbac{
		Permission ID: entry.Permission ID,
		Permission Group: entry.Permission Group,
		Security Level: entry.Security Level,
		Permission: entry.Permission,
	}
}

func (data VwRbacSlice) ToDB() DBVwRbacSlice {
	result := make(DBVwRbacSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBVwRbacSlice) ToWeb() VwRbacSlice {
	result := make(VwRbacSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
