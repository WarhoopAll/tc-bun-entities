package model

import "github.com/uptrace/bun"

type SpawnGroupTemplate struct {
	GroupId int `json:"groupid,omitempty"`
	GroupName string `json:"groupname,omitempty"`
	GroupFlags int `json:"groupflags,omitempty"`
}

type SpawnGroupTemplateSlice []SpawnGroupTemplate

type DBSpawnGroupTemplate struct {
	bun.BaseModel `bun:"table:spawn_group_template,alias:spawn_group_template"`
	GroupId int `bun:"groupId"`
	GroupName string `bun:"groupName"`
	GroupFlags int `bun:"groupFlags"`
}

type DBSpawnGroupTemplateSlice []DBSpawnGroupTemplate

func (entry *SpawnGroupTemplate) ToDB() *DBSpawnGroupTemplate {
	if entry == nil {
		return nil
	}
	return &DBSpawnGroupTemplate{
		GroupId: entry.GroupId,
		GroupName: entry.GroupName,
		GroupFlags: entry.GroupFlags,
	}
}

func (entry *DBSpawnGroupTemplate) ToWeb() *SpawnGroupTemplate {
	if entry == nil {
		return nil
	}
	return &SpawnGroupTemplate{
		GroupId: entry.GroupId,
		GroupName: entry.GroupName,
		GroupFlags: entry.GroupFlags,
	}
}

func (data SpawnGroupTemplateSlice) ToDB() DBSpawnGroupTemplateSlice {
	result := make(DBSpawnGroupTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpawnGroupTemplateSlice) ToWeb() SpawnGroupTemplateSlice {
	result := make(SpawnGroupTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
