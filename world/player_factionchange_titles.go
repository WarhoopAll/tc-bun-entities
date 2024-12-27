package model

import "github.com/uptrace/bun"

type PlayerFactionchangeTitles struct {
	AllianceId int `json:"alliance_id,omitempty"`
	HordeId int `json:"horde_id,omitempty"`
}

type PlayerFactionchangeTitlesSlice []PlayerFactionchangeTitles

type DBPlayerFactionchangeTitles struct {
	bun.BaseModel `bun:"table:player_factionchange_titles,alias:player_factionchange_titles"`
	AllianceId int `bun:"alliance_id"`
	HordeId int `bun:"horde_id"`
}

type DBPlayerFactionchangeTitlesSlice []DBPlayerFactionchangeTitles

func (entry *PlayerFactionchangeTitles) ToDB() *DBPlayerFactionchangeTitles {
	if entry == nil {
		return nil
	}
	return &DBPlayerFactionchangeTitles{
		AllianceId: entry.AllianceId,
		HordeId: entry.HordeId,
	}
}

func (entry *DBPlayerFactionchangeTitles) ToWeb() *PlayerFactionchangeTitles {
	if entry == nil {
		return nil
	}
	return &PlayerFactionchangeTitles{
		AllianceId: entry.AllianceId,
		HordeId: entry.HordeId,
	}
}

func (data PlayerFactionchangeTitlesSlice) ToDB() DBPlayerFactionchangeTitlesSlice {
	result := make(DBPlayerFactionchangeTitlesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayerFactionchangeTitlesSlice) ToWeb() PlayerFactionchangeTitlesSlice {
	result := make(PlayerFactionchangeTitlesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
