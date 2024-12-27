package model

import "github.com/uptrace/bun"

type Creature struct {
	Guid int `json:"guid,omitempty"`
	Id int `json:"id,omitempty"`
	Map int16 `json:"map,omitempty"`
	ZoneId int16 `json:"zoneid,omitempty"`
	AreaId int16 `json:"areaid,omitempty"`
	SpawnMask int8 `json:"spawnmask,omitempty"`
	PhaseMask int `json:"phasemask,omitempty"`
	Modelid int `json:"modelid,omitempty"`
	EquipmentId int8 `json:"equipment_id,omitempty"`
	PositionX float64 `json:"position_x,omitempty"`
	PositionY float64 `json:"position_y,omitempty"`
	PositionZ float64 `json:"position_z,omitempty"`
	Orientation float64 `json:"orientation,omitempty"`
	Spawntimesecs int `json:"spawntimesecs,omitempty"`
	WanderDistance float64 `json:"wander_distance,omitempty"`
	Currentwaypoint int `json:"currentwaypoint,omitempty"`
	Curhealth int `json:"curhealth,omitempty"`
	Curmana int `json:"curmana,omitempty"`
	MovementType int8 `json:"movementtype,omitempty"`
	Npcflag int `json:"npcflag,omitempty"`
	UnitFlags int `json:"unit_flags,omitempty"`
	Dynamicflags int `json:"dynamicflags,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
	StringId string `json:"stringid,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type CreatureSlice []Creature

type DBCreature struct {
	bun.BaseModel `bun:"table:creature,alias:creature"`
	Guid int `bun:"guid"`
	Id int `bun:"id"`
	Map int16 `bun:"map"`
	ZoneId int16 `bun:"zoneId"`
	AreaId int16 `bun:"areaId"`
	SpawnMask int8 `bun:"spawnMask"`
	PhaseMask int `bun:"phaseMask"`
	Modelid int `bun:"modelid"`
	EquipmentId int8 `bun:"equipment_id"`
	PositionX float64 `bun:"position_x"`
	PositionY float64 `bun:"position_y"`
	PositionZ float64 `bun:"position_z"`
	Orientation float64 `bun:"orientation"`
	Spawntimesecs int `bun:"spawntimesecs"`
	WanderDistance float64 `bun:"wander_distance"`
	Currentwaypoint int `bun:"currentwaypoint"`
	Curhealth int `bun:"curhealth"`
	Curmana int `bun:"curmana"`
	MovementType int8 `bun:"MovementType"`
	Npcflag int `bun:"npcflag"`
	UnitFlags int `bun:"unit_flags"`
	Dynamicflags int `bun:"dynamicflags"`
	ScriptName string `bun:"ScriptName"`
	StringId string `bun:"StringId"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBCreatureSlice []DBCreature

func (entry *Creature) ToDB() *DBCreature {
	if entry == nil {
		return nil
	}
	return &DBCreature{
		Guid: entry.Guid,
		Id: entry.Id,
		Map: entry.Map,
		ZoneId: entry.ZoneId,
		AreaId: entry.AreaId,
		SpawnMask: entry.SpawnMask,
		PhaseMask: entry.PhaseMask,
		Modelid: entry.Modelid,
		EquipmentId: entry.EquipmentId,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		Spawntimesecs: entry.Spawntimesecs,
		WanderDistance: entry.WanderDistance,
		Currentwaypoint: entry.Currentwaypoint,
		Curhealth: entry.Curhealth,
		Curmana: entry.Curmana,
		MovementType: entry.MovementType,
		Npcflag: entry.Npcflag,
		UnitFlags: entry.UnitFlags,
		Dynamicflags: entry.Dynamicflags,
		ScriptName: entry.ScriptName,
		StringId: entry.StringId,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBCreature) ToWeb() *Creature {
	if entry == nil {
		return nil
	}
	return &Creature{
		Guid: entry.Guid,
		Id: entry.Id,
		Map: entry.Map,
		ZoneId: entry.ZoneId,
		AreaId: entry.AreaId,
		SpawnMask: entry.SpawnMask,
		PhaseMask: entry.PhaseMask,
		Modelid: entry.Modelid,
		EquipmentId: entry.EquipmentId,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		Spawntimesecs: entry.Spawntimesecs,
		WanderDistance: entry.WanderDistance,
		Currentwaypoint: entry.Currentwaypoint,
		Curhealth: entry.Curhealth,
		Curmana: entry.Curmana,
		MovementType: entry.MovementType,
		Npcflag: entry.Npcflag,
		UnitFlags: entry.UnitFlags,
		Dynamicflags: entry.Dynamicflags,
		ScriptName: entry.ScriptName,
		StringId: entry.StringId,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data CreatureSlice) ToDB() DBCreatureSlice {
	result := make(DBCreatureSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureSlice) ToWeb() CreatureSlice {
	result := make(CreatureSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
