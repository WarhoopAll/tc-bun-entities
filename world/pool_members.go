package model

import "github.com/uptrace/bun"

type PoolMembers struct {
	Type int16 `json:"type,omitempty"`
	SpawnId int `json:"spawnid,omitempty"`
	PoolSpawnId int `json:"poolspawnid,omitempty"`
	Chance float64 `json:"chance,omitempty"`
	Description string `json:"description,omitempty"`
}

type PoolMembersSlice []PoolMembers

type DBPoolMembers struct {
	bun.BaseModel `bun:"table:pool_members,alias:pool_members"`
	Type int16 `bun:"type"`
	SpawnId int `bun:"spawnId"`
	PoolSpawnId int `bun:"poolSpawnId"`
	Chance float64 `bun:"chance"`
	Description string `bun:"description"`
}

type DBPoolMembersSlice []DBPoolMembers

func (entry *PoolMembers) ToDB() *DBPoolMembers {
	if entry == nil {
		return nil
	}
	return &DBPoolMembers{
		Type: entry.Type,
		SpawnId: entry.SpawnId,
		PoolSpawnId: entry.PoolSpawnId,
		Chance: entry.Chance,
		Description: entry.Description,
	}
}

func (entry *DBPoolMembers) ToWeb() *PoolMembers {
	if entry == nil {
		return nil
	}
	return &PoolMembers{
		Type: entry.Type,
		SpawnId: entry.SpawnId,
		PoolSpawnId: entry.PoolSpawnId,
		Chance: entry.Chance,
		Description: entry.Description,
	}
}

func (data PoolMembersSlice) ToDB() DBPoolMembersSlice {
	result := make(DBPoolMembersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPoolMembersSlice) ToWeb() PoolMembersSlice {
	result := make(PoolMembersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
