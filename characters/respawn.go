package model

import "github.com/uptrace/bun"

type Respawn struct {
	Type int16 `json:"type,omitempty"`
	SpawnId int `json:"spawnid,omitempty"`
	RespawnTime int `json:"respawntime,omitempty"`
	MapId int16 `json:"mapid,omitempty"`
	InstanceId int `json:"instanceid,omitempty"`
}

type RespawnSlice []Respawn

type DBRespawn struct {
	bun.BaseModel `bun:"table:respawn,alias:respawn"`
	Type int16 `bun:"type"`
	SpawnId int `bun:"spawnId"`
	RespawnTime int `bun:"respawnTime"`
	MapId int16 `bun:"mapId"`
	InstanceId int `bun:"instanceId"`
}

type DBRespawnSlice []DBRespawn

func (entry *Respawn) ToDB() *DBRespawn {
	if entry == nil {
		return nil
	}
	return &DBRespawn{
		Type: entry.Type,
		SpawnId: entry.SpawnId,
		RespawnTime: entry.RespawnTime,
		MapId: entry.MapId,
		InstanceId: entry.InstanceId,
	}
}

func (entry *DBRespawn) ToWeb() *Respawn {
	if entry == nil {
		return nil
	}
	return &Respawn{
		Type: entry.Type,
		SpawnId: entry.SpawnId,
		RespawnTime: entry.RespawnTime,
		MapId: entry.MapId,
		InstanceId: entry.InstanceId,
	}
}

func (data RespawnSlice) ToDB() DBRespawnSlice {
	result := make(DBRespawnSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBRespawnSlice) ToWeb() RespawnSlice {
	result := make(RespawnSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
