package model

import "github.com/uptrace/bun"

type CreatureMovementInfo struct {
	MovementID int `json:"movementid,omitempty"`
	WalkSpeed float64 `json:"walkspeed,omitempty"`
	RunSpeed float64 `json:"runspeed,omitempty"`
}

type CreatureMovementInfoSlice []CreatureMovementInfo

type DBCreatureMovementInfo struct {
	bun.BaseModel `bun:"table:creature_movement_info,alias:creature_movement_info"`
	MovementID int `bun:"MovementID"`
	WalkSpeed float64 `bun:"WalkSpeed"`
	RunSpeed float64 `bun:"RunSpeed"`
}

type DBCreatureMovementInfoSlice []DBCreatureMovementInfo

func (entry *CreatureMovementInfo) ToDB() *DBCreatureMovementInfo {
	if entry == nil {
		return nil
	}
	return &DBCreatureMovementInfo{
		MovementID: entry.MovementID,
		WalkSpeed: entry.WalkSpeed,
		RunSpeed: entry.RunSpeed,
	}
}

func (entry *DBCreatureMovementInfo) ToWeb() *CreatureMovementInfo {
	if entry == nil {
		return nil
	}
	return &CreatureMovementInfo{
		MovementID: entry.MovementID,
		WalkSpeed: entry.WalkSpeed,
		RunSpeed: entry.RunSpeed,
	}
}

func (data CreatureMovementInfoSlice) ToDB() DBCreatureMovementInfoSlice {
	result := make(DBCreatureMovementInfoSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureMovementInfoSlice) ToWeb() CreatureMovementInfoSlice {
	result := make(CreatureMovementInfoSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
