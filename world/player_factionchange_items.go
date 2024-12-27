package model

import "github.com/uptrace/bun"

type PlayerFactionchangeItems struct {
	RaceA int `json:"race_a,omitempty"`
	AllianceId int `json:"alliance_id,omitempty"`
	CommentA string `json:"commenta,omitempty"`
	RaceH int `json:"race_h,omitempty"`
	HordeId int `json:"horde_id,omitempty"`
	CommentH string `json:"commenth,omitempty"`
}

type PlayerFactionchangeItemsSlice []PlayerFactionchangeItems

type DBPlayerFactionchangeItems struct {
	bun.BaseModel `bun:"table:player_factionchange_items,alias:player_factionchange_items"`
	RaceA int `bun:"race_A"`
	AllianceId int `bun:"alliance_id"`
	CommentA string `bun:"commentA"`
	RaceH int `bun:"race_H"`
	HordeId int `bun:"horde_id"`
	CommentH string `bun:"commentH"`
}

type DBPlayerFactionchangeItemsSlice []DBPlayerFactionchangeItems

func (entry *PlayerFactionchangeItems) ToDB() *DBPlayerFactionchangeItems {
	if entry == nil {
		return nil
	}
	return &DBPlayerFactionchangeItems{
		RaceA: entry.RaceA,
		AllianceId: entry.AllianceId,
		CommentA: entry.CommentA,
		RaceH: entry.RaceH,
		HordeId: entry.HordeId,
		CommentH: entry.CommentH,
	}
}

func (entry *DBPlayerFactionchangeItems) ToWeb() *PlayerFactionchangeItems {
	if entry == nil {
		return nil
	}
	return &PlayerFactionchangeItems{
		RaceA: entry.RaceA,
		AllianceId: entry.AllianceId,
		CommentA: entry.CommentA,
		RaceH: entry.RaceH,
		HordeId: entry.HordeId,
		CommentH: entry.CommentH,
	}
}

func (data PlayerFactionchangeItemsSlice) ToDB() DBPlayerFactionchangeItemsSlice {
	result := make(DBPlayerFactionchangeItemsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayerFactionchangeItemsSlice) ToWeb() PlayerFactionchangeItemsSlice {
	result := make(PlayerFactionchangeItemsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
