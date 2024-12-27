package model

import "github.com/uptrace/bun"

type InstanceSpawnGroups struct {
	InstanceMapId int16 `json:"instancemapid,omitempty"`
	BossStateId int8 `json:"bossstateid,omitempty"`
	BossStates int8 `json:"bossstates,omitempty"`
	SpawnGroupId int `json:"spawngroupid,omitempty"`
	Flags int8 `json:"flags,omitempty"`
}

type InstanceSpawnGroupsSlice []InstanceSpawnGroups

type DBInstanceSpawnGroups struct {
	bun.BaseModel `bun:"table:instance_spawn_groups,alias:instance_spawn_groups"`
	InstanceMapId int16 `bun:"instanceMapId"`
	BossStateId int8 `bun:"bossStateId"`
	BossStates int8 `bun:"bossStates"`
	SpawnGroupId int `bun:"spawnGroupId"`
	Flags int8 `bun:"flags"`
}

type DBInstanceSpawnGroupsSlice []DBInstanceSpawnGroups

func (entry *InstanceSpawnGroups) ToDB() *DBInstanceSpawnGroups {
	if entry == nil {
		return nil
	}
	return &DBInstanceSpawnGroups{
		InstanceMapId: entry.InstanceMapId,
		BossStateId: entry.BossStateId,
		BossStates: entry.BossStates,
		SpawnGroupId: entry.SpawnGroupId,
		Flags: entry.Flags,
	}
}

func (entry *DBInstanceSpawnGroups) ToWeb() *InstanceSpawnGroups {
	if entry == nil {
		return nil
	}
	return &InstanceSpawnGroups{
		InstanceMapId: entry.InstanceMapId,
		BossStateId: entry.BossStateId,
		BossStates: entry.BossStates,
		SpawnGroupId: entry.SpawnGroupId,
		Flags: entry.Flags,
	}
}

func (data InstanceSpawnGroupsSlice) ToDB() DBInstanceSpawnGroupsSlice {
	result := make(DBInstanceSpawnGroupsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBInstanceSpawnGroupsSlice) ToWeb() InstanceSpawnGroupsSlice {
	result := make(InstanceSpawnGroupsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
