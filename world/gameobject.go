package model

import "github.com/uptrace/bun"

type Gameobject struct {
	Guid int `json:"guid,omitempty"`
	Id int `json:"id,omitempty"`
	Map int16 `json:"map,omitempty"`
	ZoneId int16 `json:"zoneid,omitempty"`
	AreaId int16 `json:"areaid,omitempty"`
	SpawnMask int8 `json:"spawnmask,omitempty"`
	PhaseMask int `json:"phasemask,omitempty"`
	PositionX float64 `json:"position_x,omitempty"`
	PositionY float64 `json:"position_y,omitempty"`
	PositionZ float64 `json:"position_z,omitempty"`
	Orientation float64 `json:"orientation,omitempty"`
	Rotation0 float64 `json:"rotation0,omitempty"`
	Rotation1 float64 `json:"rotation1,omitempty"`
	Rotation2 float64 `json:"rotation2,omitempty"`
	Rotation3 float64 `json:"rotation3,omitempty"`
	Spawntimesecs int `json:"spawntimesecs,omitempty"`
	Animprogress int8 `json:"animprogress,omitempty"`
	State int8 `json:"state,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
	StringId string `json:"stringid,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type GameobjectSlice []Gameobject

type DBGameobject struct {
	bun.BaseModel `bun:"table:gameobject,alias:gameobject"`
	Guid int `bun:"guid"`
	Id int `bun:"id"`
	Map int16 `bun:"map"`
	ZoneId int16 `bun:"zoneId"`
	AreaId int16 `bun:"areaId"`
	SpawnMask int8 `bun:"spawnMask"`
	PhaseMask int `bun:"phaseMask"`
	PositionX float64 `bun:"position_x"`
	PositionY float64 `bun:"position_y"`
	PositionZ float64 `bun:"position_z"`
	Orientation float64 `bun:"orientation"`
	Rotation0 float64 `bun:"rotation0"`
	Rotation1 float64 `bun:"rotation1"`
	Rotation2 float64 `bun:"rotation2"`
	Rotation3 float64 `bun:"rotation3"`
	Spawntimesecs int `bun:"spawntimesecs"`
	Animprogress int8 `bun:"animprogress"`
	State int8 `bun:"state"`
	ScriptName string `bun:"ScriptName"`
	StringId string `bun:"StringId"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBGameobjectSlice []DBGameobject

func (entry *Gameobject) ToDB() *DBGameobject {
	if entry == nil {
		return nil
	}
	return &DBGameobject{
		Guid: entry.Guid,
		Id: entry.Id,
		Map: entry.Map,
		ZoneId: entry.ZoneId,
		AreaId: entry.AreaId,
		SpawnMask: entry.SpawnMask,
		PhaseMask: entry.PhaseMask,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		Rotation0: entry.Rotation0,
		Rotation1: entry.Rotation1,
		Rotation2: entry.Rotation2,
		Rotation3: entry.Rotation3,
		Spawntimesecs: entry.Spawntimesecs,
		Animprogress: entry.Animprogress,
		State: entry.State,
		ScriptName: entry.ScriptName,
		StringId: entry.StringId,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBGameobject) ToWeb() *Gameobject {
	if entry == nil {
		return nil
	}
	return &Gameobject{
		Guid: entry.Guid,
		Id: entry.Id,
		Map: entry.Map,
		ZoneId: entry.ZoneId,
		AreaId: entry.AreaId,
		SpawnMask: entry.SpawnMask,
		PhaseMask: entry.PhaseMask,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		Rotation0: entry.Rotation0,
		Rotation1: entry.Rotation1,
		Rotation2: entry.Rotation2,
		Rotation3: entry.Rotation3,
		Spawntimesecs: entry.Spawntimesecs,
		Animprogress: entry.Animprogress,
		State: entry.State,
		ScriptName: entry.ScriptName,
		StringId: entry.StringId,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data GameobjectSlice) ToDB() DBGameobjectSlice {
	result := make(DBGameobjectSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameobjectSlice) ToWeb() GameobjectSlice {
	result := make(GameobjectSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
