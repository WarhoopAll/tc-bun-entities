package model

import "github.com/uptrace/bun"

type PlayerFactionchangeQuests struct {
	AllianceId int `json:"alliance_id,omitempty"`
	HordeId int `json:"horde_id,omitempty"`
}

type PlayerFactionchangeQuestsSlice []PlayerFactionchangeQuests

type DBPlayerFactionchangeQuests struct {
	bun.BaseModel `bun:"table:player_factionchange_quests,alias:player_factionchange_quests"`
	AllianceId int `bun:"alliance_id"`
	HordeId int `bun:"horde_id"`
}

type DBPlayerFactionchangeQuestsSlice []DBPlayerFactionchangeQuests

func (entry *PlayerFactionchangeQuests) ToDB() *DBPlayerFactionchangeQuests {
	if entry == nil {
		return nil
	}
	return &DBPlayerFactionchangeQuests{
		AllianceId: entry.AllianceId,
		HordeId: entry.HordeId,
	}
}

func (entry *DBPlayerFactionchangeQuests) ToWeb() *PlayerFactionchangeQuests {
	if entry == nil {
		return nil
	}
	return &PlayerFactionchangeQuests{
		AllianceId: entry.AllianceId,
		HordeId: entry.HordeId,
	}
}

func (data PlayerFactionchangeQuestsSlice) ToDB() DBPlayerFactionchangeQuestsSlice {
	result := make(DBPlayerFactionchangeQuestsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayerFactionchangeQuestsSlice) ToWeb() PlayerFactionchangeQuestsSlice {
	result := make(PlayerFactionchangeQuestsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
