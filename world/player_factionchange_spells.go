package model

import "github.com/uptrace/bun"

type PlayerFactionchangeSpells struct {
	AllianceId int `json:"alliance_id,omitempty"`
	HordeId int `json:"horde_id,omitempty"`
}

type PlayerFactionchangeSpellsSlice []PlayerFactionchangeSpells

type DBPlayerFactionchangeSpells struct {
	bun.BaseModel `bun:"table:player_factionchange_spells,alias:player_factionchange_spells"`
	AllianceId int `bun:"alliance_id"`
	HordeId int `bun:"horde_id"`
}

type DBPlayerFactionchangeSpellsSlice []DBPlayerFactionchangeSpells

func (entry *PlayerFactionchangeSpells) ToDB() *DBPlayerFactionchangeSpells {
	if entry == nil {
		return nil
	}
	return &DBPlayerFactionchangeSpells{
		AllianceId: entry.AllianceId,
		HordeId: entry.HordeId,
	}
}

func (entry *DBPlayerFactionchangeSpells) ToWeb() *PlayerFactionchangeSpells {
	if entry == nil {
		return nil
	}
	return &PlayerFactionchangeSpells{
		AllianceId: entry.AllianceId,
		HordeId: entry.HordeId,
	}
}

func (data PlayerFactionchangeSpellsSlice) ToDB() DBPlayerFactionchangeSpellsSlice {
	result := make(DBPlayerFactionchangeSpellsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayerFactionchangeSpellsSlice) ToWeb() PlayerFactionchangeSpellsSlice {
	result := make(PlayerFactionchangeSpellsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
