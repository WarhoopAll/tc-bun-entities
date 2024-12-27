package model

import "github.com/uptrace/bun"

type AreatriggerTeleport struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	TargetMap int16 `json:"target_map,omitempty"`
	TargetPositionX float64 `json:"target_position_x,omitempty"`
	TargetPositionY float64 `json:"target_position_y,omitempty"`
	TargetPositionZ float64 `json:"target_position_z,omitempty"`
	TargetOrientation float64 `json:"target_orientation,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type AreatriggerTeleportSlice []AreatriggerTeleport

type DBAreatriggerTeleport struct {
	bun.BaseModel `bun:"table:areatrigger_teleport,alias:areatrigger_teleport"`
	ID int `bun:"ID"`
	Name string `bun:"Name"`
	TargetMap int16 `bun:"target_map"`
	TargetPositionX float64 `bun:"target_position_x"`
	TargetPositionY float64 `bun:"target_position_y"`
	TargetPositionZ float64 `bun:"target_position_z"`
	TargetOrientation float64 `bun:"target_orientation"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBAreatriggerTeleportSlice []DBAreatriggerTeleport

func (entry *AreatriggerTeleport) ToDB() *DBAreatriggerTeleport {
	if entry == nil {
		return nil
	}
	return &DBAreatriggerTeleport{
		ID: entry.ID,
		Name: entry.Name,
		TargetMap: entry.TargetMap,
		TargetPositionX: entry.TargetPositionX,
		TargetPositionY: entry.TargetPositionY,
		TargetPositionZ: entry.TargetPositionZ,
		TargetOrientation: entry.TargetOrientation,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBAreatriggerTeleport) ToWeb() *AreatriggerTeleport {
	if entry == nil {
		return nil
	}
	return &AreatriggerTeleport{
		ID: entry.ID,
		Name: entry.Name,
		TargetMap: entry.TargetMap,
		TargetPositionX: entry.TargetPositionX,
		TargetPositionY: entry.TargetPositionY,
		TargetPositionZ: entry.TargetPositionZ,
		TargetOrientation: entry.TargetOrientation,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data AreatriggerTeleportSlice) ToDB() DBAreatriggerTeleportSlice {
	result := make(DBAreatriggerTeleportSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBAreatriggerTeleportSlice) ToWeb() AreatriggerTeleportSlice {
	result := make(AreatriggerTeleportSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
