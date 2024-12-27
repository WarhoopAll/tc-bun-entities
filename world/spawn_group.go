package model

import "github.com/uptrace/bun"

type SpawnGroup struct {
	GroupId int `json:"groupid,omitempty"`
	SpawnType int8 `json:"spawntype,omitempty"`
	SpawnId int `json:"spawnid,omitempty"`
}

type SpawnGroupSlice []SpawnGroup

type DBSpawnGroup struct {
	bun.BaseModel `bun:"table:spawn_group,alias:spawn_group"`
	GroupId int `bun:"groupId"`
	SpawnType int8 `bun:"spawnType"`
	SpawnId int `bun:"spawnId"`
}

type DBSpawnGroupSlice []DBSpawnGroup

func (entry *SpawnGroup) ToDB() *DBSpawnGroup {
	if entry == nil {
		return nil
	}
	return &DBSpawnGroup{
		GroupId: entry.GroupId,
		SpawnType: entry.SpawnType,
		SpawnId: entry.SpawnId,
	}
}

func (entry *DBSpawnGroup) ToWeb() *SpawnGroup {
	if entry == nil {
		return nil
	}
	return &SpawnGroup{
		GroupId: entry.GroupId,
		SpawnType: entry.SpawnType,
		SpawnId: entry.SpawnId,
	}
}

func (data SpawnGroupSlice) ToDB() DBSpawnGroupSlice {
	result := make(DBSpawnGroupSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpawnGroupSlice) ToWeb() SpawnGroupSlice {
	result := make(SpawnGroupSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
