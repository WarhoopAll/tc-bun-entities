package model

import "github.com/uptrace/bun"

type PlayerFactionchangeReputations struct {
	AllianceId int `json:"alliance_id,omitempty"`
	HordeId int `json:"horde_id,omitempty"`
}

type PlayerFactionchangeReputationsSlice []PlayerFactionchangeReputations

type DBPlayerFactionchangeReputations struct {
	bun.BaseModel `bun:"table:player_factionchange_reputations,alias:player_factionchange_reputations"`
	AllianceId int `bun:"alliance_id"`
	HordeId int `bun:"horde_id"`
}

type DBPlayerFactionchangeReputationsSlice []DBPlayerFactionchangeReputations

func (entry *PlayerFactionchangeReputations) ToDB() *DBPlayerFactionchangeReputations {
	if entry == nil {
		return nil
	}
	return &DBPlayerFactionchangeReputations{
		AllianceId: entry.AllianceId,
		HordeId: entry.HordeId,
	}
}

func (entry *DBPlayerFactionchangeReputations) ToWeb() *PlayerFactionchangeReputations {
	if entry == nil {
		return nil
	}
	return &PlayerFactionchangeReputations{
		AllianceId: entry.AllianceId,
		HordeId: entry.HordeId,
	}
}

func (data PlayerFactionchangeReputationsSlice) ToDB() DBPlayerFactionchangeReputationsSlice {
	result := make(DBPlayerFactionchangeReputationsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayerFactionchangeReputationsSlice) ToWeb() PlayerFactionchangeReputationsSlice {
	result := make(PlayerFactionchangeReputationsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
