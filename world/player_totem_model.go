package model

import "github.com/uptrace/bun"

type PlayerTotemModel struct {
	TotemSlot int8 `json:"totemslot,omitempty"`
	RaceId int8 `json:"raceid,omitempty"`
	DisplayId int `json:"displayid,omitempty"`
}

type PlayerTotemModelSlice []PlayerTotemModel

type DBPlayerTotemModel struct {
	bun.BaseModel `bun:"table:player_totem_model,alias:player_totem_model"`
	TotemSlot int8 `bun:"TotemSlot"`
	RaceId int8 `bun:"RaceId"`
	DisplayId int `bun:"DisplayId"`
}

type DBPlayerTotemModelSlice []DBPlayerTotemModel

func (entry *PlayerTotemModel) ToDB() *DBPlayerTotemModel {
	if entry == nil {
		return nil
	}
	return &DBPlayerTotemModel{
		TotemSlot: entry.TotemSlot,
		RaceId: entry.RaceId,
		DisplayId: entry.DisplayId,
	}
}

func (entry *DBPlayerTotemModel) ToWeb() *PlayerTotemModel {
	if entry == nil {
		return nil
	}
	return &PlayerTotemModel{
		TotemSlot: entry.TotemSlot,
		RaceId: entry.RaceId,
		DisplayId: entry.DisplayId,
	}
}

func (data PlayerTotemModelSlice) ToDB() DBPlayerTotemModelSlice {
	result := make(DBPlayerTotemModelSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPlayerTotemModelSlice) ToWeb() PlayerTotemModelSlice {
	result := make(PlayerTotemModelSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
